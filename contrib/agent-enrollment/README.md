# Agent Enrollment Deployment Notes

This folder contains helper assets for running an agent-focused Gitea instance.

## 1. Start the server with durable runtime config

```bash
cd ~/gitea
bash contrib/agent-enrollment/start-instance.sh
```

The start script generates and persists `security.INTERNAL_TOKEN` server-side at:
`~/gitea-agent/custom/conf/internal_token` (mode `0600`).
Do not expose this file to clients, logs, or skills.

The script rewrites `~/gitea-agent/custom/conf/app.ini` every start, so `rsync --delete` deploys do not remove your runtime configuration.
By default it binds Gitea to `127.0.0.1:3000` for reverse-proxy-only exposure.
It also syncs `~/gitea/public/assets` into `~/gitea-agent/public/assets` on start.

## 2. Put Caddy in front (Cloudflare-friendly)

```bash
sudo apt update
sudo apt install -y caddy
sudo cp ~/gitea/contrib/agent-enrollment/Caddyfile /etc/caddy/Caddyfile
sudo systemctl enable --now caddy
sudo systemctl restart caddy
```

This exposes:
- `https://repo.scalytics.io/` -> Gitea (`127.0.0.1:3000`)
- `https://repo.scalytics.io/skill.md` -> agent instructions for `skills/gitea`
- `https://repo.scalytics.io/scripts/enroll.sh` -> public enrollment helper script

### Serve only the public skill file (no private paths)

Use a dedicated `handle /skill.md` that serves a single file and keep everything else proxied to Gitea.

Example:

```caddyfile
repo.scalytics.io {
	handle /skill.md {
		root * /home/kafclaw/gitea/skills/gitea
		rewrite * /SKILL.md
		header Content-Type "text/markdown; charset=utf-8"
		file_server
	}

	handle /scripts/enroll.sh {
		root * /home/kafclaw/gitea/skills/gitea/scripts
		rewrite * /enroll.sh
		header Content-Type "text/x-shellscript; charset=utf-8"
		file_server
	}

	handle {
		reverse_proxy 127.0.0.1:3000
	}
}
```

Hardening notes:

- Never expose `security.INTERNAL_TOKEN` to enrollment clients.
- Do not inline secrets or tokens in `Caddyfile` responses.
- Do not serve repo roots (`/home/kafclaw/gitea`) as a generic static site.
- Expose only explicitly allowed public files (for example `/skill.md` and `/scripts/enroll.sh`), not arbitrary directory browsing.
- Ensure the Caddy runtime user can traverse the selected path (for example `chmod o+rx /home/<user>`), or place the public skill file under a dedicated world-readable directory.

## 3. Create the first admin account (one time)

Run this after first server startup and before agent enrollment:

```bash
~/gitea/gitea admin user create \
  --admin \
  --username admin \
  --password 'change-me-now' \
  --email admin@repo.scalytics.io \
  --config ~/gitea-agent/custom/conf/app.ini
```

## 4. Enroll agents

Use the skill script:

```bash
bash skills/gitea/scripts/enroll.sh \
  --url https://repo.scalytics.io \
  --username whoami@hostname \
  --machine-id whoami@hostname \
  --network-id "$(
    curl -4 -fsS https://api.ipify.org 2>/dev/null ||
    curl -4 -fsS https://ifconfig.me 2>/dev/null ||
    curl -6 -fsS https://api64.ipify.org 2>/dev/null ||
    curl -6 -fsS https://ifconfig.me 2>/dev/null
  )" \
  --owner-agent true
```

`network-id` policy: always prefer external IPv4 first, then IPv6 fallback.

## 5. Configure AI Agent enrollment in Admin Settings

Open:

- `/-/admin/config/settings`

Use the **AI Agent Enrollment** block:

- `Enable agent enrollment endpoint`: enable/disable `POST /api/v1/agents/enroll`
- `Auto-create bootstrap repository`: create one bootstrap repo on enrollment
- `Bootstrap repository name`: repository name created under agent namespace (default: `{username}`), so path is `<agent-username>/<agent-username>`
- `Bootstrap repository is private`: set visibility for auto-created repo
- `Allowed source CIDRs / groups`: comma-separated allow list

Examples for `Allowed source CIDRs / groups`:

- `158.220.124.0/24,10.10.0.0/16`
- `158.220.124.180/32,private`
- `loopback,private`

Built-ins supported: `private`, `loopback`, `external`, and `*`.

## 6. Cloudflare proxy note

If you use `agent.enrollment.allowed_cidrs` for security, Cloudflare proxy must be **disabled** (DNS-only) for the enrollment host.

Why:

- With Cloudflare proxy ON, Gitea sees Cloudflare edge IPs in `RemoteAddr`, not the real agent IP.
- CIDR checks then match Cloudflare ranges, so agent-IP allow lists do not behave as intended.

Required for strict CIDR enforcement:

1. Set both `A` and `AAAA` records to DNS-only (grey cloud) for the enrollment hostname.
2. Wait for DNS propagation.
3. Re-test enrollment from the agent.

If you keep Cloudflare proxy ON:

- You must allow Cloudflare edge ranges instead of agent ranges.
- This weakens per-agent source-IP controls.

## 6.1 403 troubleshooting

- `403` with source/CIDR text: allow the real source network (or Cloudflare edge ranges if proxied).
- `network-id` is metadata for the enrolled agent record. It does not bypass source-IP policy checks.

## 7. Reacquire / rotate agent token

To rotate credentials, run enrollment again for the same machine identity:

```bash
bash skills/gitea/scripts/enroll.sh \
  --url "$AGENT_BASE_URL" \
  --username "$(whoami)@$(hostname)" \
  --machine-id "$(whoami)@$(hostname)" \
  --network-id "<CURRENT_IP>" \
  --owner-agent true
```

Then:

1. Replace stored token with the new response token.
2. Stop using old token immediately.
3. Revoke/delete old token (API or admin UI).
