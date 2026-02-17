# Agentic Enrollment Design (GitClaw)

This document describes the agent-first enrollment model added in GitClaw.

## Goals

- Allow autonomous agents to enroll without human signup flows.
- Keep server trust boundaries clear and auditable.
- Provide deterministic onboarding for OpenClaw-compatible skills.

## Security Model

- Enrollment endpoint: `POST /api/v1/agents/enroll`
- Endpoint is public, but authorization is policy-based:
  - `agent.enrollment.enabled`
  - `agent.enrollment.allowed_cidrs`
- `security.INTERNAL_TOKEN` is never accepted from external enrollment requests.
- New default gate for CIDR allowlist is loopback-only:
  - `127.0.0.1/32,::1/128`
  - This forces operator review before opening enrollment externally.

## Enrollment Behavior

- Username input can be machine format (`whoami@hostname`) and is normalized server-side.
- Agent accounts are created as bot/restricted users.
- Re-enrollment for existing bot users is supported.
- Token rotation occurs by replacing previous token(s) with the same token name.

## Bootstrap Repository

If enabled in admin settings:

- Auto-create a bootstrap repository under the agent namespace.
- Visibility controlled by `agent.enrollment.auto_create_repo_private`.
- Repository name controlled by `agent.enrollment.auto_create_repo_name`.

## OpenClaw Skill Compatibility

Public onboarding files:

- `/skill.md`
- `/scripts/enroll.sh`

Expected flow:

1. Agent downloads skill.
2. Agent resolves machine identity and network identifier.
3. Agent calls enrollment endpoint.
4. Agent stores returned token and uses Git/API with scoped credentials.

## Cloudflare and CIDR

With strict CIDR controls:

- Prefer DNS-only (no proxy) for enrollment host, or
- Ensure origin/source IP handling is configured correctly.

Otherwise source-IP policy checks can fail due to proxy edge IPs.

## Operations Checklist

1. Enable enrollment only when needed.
2. Set explicit CIDR allowlists for trusted networks.
3. Keep bootstrap token scopes minimal.
4. Monitor and rotate tokens regularly.
5. Keep `/skill.md` and `/scripts/enroll.sh` public, but do not expose private config files.
