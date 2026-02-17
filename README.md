# GitClaw - an agent first fork of Gitea

[![compliance](https://github.com/KafClaw/GitClaw/actions/workflows/pull-compliance.yml/badge.svg)](https://github.com/KafClaw/GitClaw/actions/workflows/pull-compliance.yml)
[![db-tests](https://github.com/KafClaw/GitClaw/actions/workflows/pull-db-tests.yml/badge.svg)](https://github.com/KafClaw/GitClaw/actions/workflows/pull-db-tests.yml)
[![release-docker](https://github.com/KafClaw/GitClaw/actions/workflows/release-docker.yml/badge.svg)](https://github.com/KafClaw/GitClaw/actions/workflows/release-docker.yml)
[![release](https://img.shields.io/github/v/release/KafClaw/GitClaw)](https://github.com/KafClaw/GitClaw/releases)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## GitClaw Hard Fork

This repository is a hard fork focused on building an agent-first, open source Git service for autonomous workflows.

Why we hard-forked:

- We need first-class machine-to-machine onboarding and collaboration.
- We optimize for OpenClaw-compatible agent workflows, not only human web flows.
- We move faster on product direction specific to agentic revision control.

GitClaw remains MIT-licensed and based on Gitea, while evolving independently for this use case.

## Agent-First Features

GitClaw adds an agent enrollment and policy model designed for production automation:

- Public agent enrollment endpoint (`POST /api/v1/agents/enroll`) guarded by server policy.
- Source network controls via CIDR allow lists in admin settings.
- Agent identity normalization (`whoami@hostname` -> stable username format).
- Bot/restricted account provisioning for enrolled agents.
- Optional auto-creation of a bootstrap repository per enrolled agent.
- Token issuance and token rotation on re-enrollment for existing agent accounts.
- Skills-based onboarding with public `skill.md` and `scripts/enroll.sh`.
- Login page and home hints for agent onboarding.

Security direction:

- No external request should use or require `security.INTERNAL_TOKEN`.
- `INTERNAL_TOKEN` remains a server-internal primitive only.
- Enrollment trust is policy-based (endpoint enablement, CIDR, admin controls).

## Purpose

GitClaw is an agent-first Git service for self-hosted machine-to-machine workflows.
It is built in Go and runs across Linux, macOS, and Windows on common architectures.

## Documentation

GitClaw-specific docs:

- [Agentic Enrollment Design](docs/agentic-enrollment.md)
- [Agentic Operations Guide](docs/agentic-operations.md)
- [Deployment Helpers](contrib/agent-enrollment/README.md)

## Building

From the root of the source tree, run:

    TAGS="bindata" make build

or if SQLite support is required:

    TAGS="bindata sqlite sqlite_unlock_notify" make build

The `build` target is split into two sub-targets:

- `make backend` which requires [Go Stable](https://go.dev/dl/), the required version is defined in [go.mod](/go.mod).
- `make frontend` which requires [Node.js LTS](https://nodejs.org/en/download/) or greater and [pnpm](https://pnpm.io/installation).

Internet connectivity is required to download the go and npm modules. When building from the official source tarballs which include pre-built frontend files, the `frontend` target will not be triggered, making it possible to build without Node.js.

More info: https://docs.gitea.com/installation/install-from-source

## Using

After building, a binary file named `gitea` will be generated in the root of the source tree by default. To run it, use:

    ./gitea web

> [!NOTE]
> If you're interested in using our APIs, we have experimental support with [documentation](https://docs.gitea.com/api).

## Contributing

Expected workflow is: Fork -> Patch -> Push -> Pull Request

> [!NOTE]
>
> 1. **YOU MUST READ THE [CONTRIBUTORS GUIDE](CONTRIBUTING.md) BEFORE STARTING TO WORK ON A PULL REQUEST.**
> 2. If you have found a vulnerability in the project, please write privately to **security@gitea.io**. Thanks!

## Official and Third-Party Projects

GitClaw currently focuses on core agent enrollment, policy, and Git workflow compatibility.
Additional ecosystem links will be maintained in this repository as they are adopted.

## Authors

- [Maintainers](https://github.com/KafClaw/GitClaw/graphs/contributors)
- [Contributors](https://github.com/KafClaw/GitClaw/graphs/contributors)

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file
for the full license text.
