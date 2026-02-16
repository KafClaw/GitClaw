# Gitea

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

The goal of this project is to make the easiest, fastest, and most
painless way of setting up a self-hosted Git service.

As Gitea is written in Go, it works across **all** the platforms and
architectures that are supported by Go, including Linux, macOS, and
Windows on x86, amd64, ARM and PowerPC architectures.
This project has been
[forked](https://blog.gitea.com/welcome-to-gitea/) from
[Gogs](https://gogs.io) since November of 2016, but a lot has changed.

For online demonstrations, you can visit [demo.gitea.com](https://demo.gitea.com).

For accessing free Gitea service (with a limited number of repositories), you can visit [gitea.com](https://gitea.com/user/login).

To quickly deploy your own dedicated Gitea instance on Gitea Cloud, you can start a free trial at [cloud.gitea.com](https://cloud.gitea.com).

## Documentation

You can find comprehensive documentation on Gitea's official [documentation website](https://docs.gitea.com/).

It includes installation, administration, usage, development, contributing guides, and more to help you get started and explore all features effectively.

If you have any suggestions or would like to contribute to it, you can visit the [documentation repository](https://gitea.com/gitea/docs)

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

We provide an official [go-sdk](https://gitea.com/gitea/go-sdk), a CLI tool called [tea](https://gitea.com/gitea/tea) and an [action runner](https://gitea.com/gitea/act_runner) for Gitea Action.

We maintain a list of Gitea-related projects at [gitea/awesome-gitea](https://gitea.com/gitea/awesome-gitea), where you can discover more third-party projects, including SDKs, plugins, themes, and more.

## Authors

- [Maintainers](https://github.com/KafClaw/GitClaw/graphs/contributors)
- [Contributors](https://github.com/KafClaw/GitClaw/graphs/contributors)

## License

This project is licensed under the MIT License.
See the [LICENSE](https://github.com/go-gitea/gitea/blob/main/LICENSE) file
for the full license text.

## Further information

<details>
<summary>Looking for an overview of the interface? Check it out!</summary>

### Login/Register Page

![Login](https://dl.gitea.com/screenshots/login.png)
![Register](https://dl.gitea.com/screenshots/register.png)

### User Dashboard

![Home](https://dl.gitea.com/screenshots/home.png)
![Issues](https://dl.gitea.com/screenshots/issues.png)
![Pull Requests](https://dl.gitea.com/screenshots/pull_requests.png)
![Milestones](https://dl.gitea.com/screenshots/milestones.png)

### User Profile

![Profile](https://dl.gitea.com/screenshots/user_profile.png)

### Explore

![Repos](https://dl.gitea.com/screenshots/explore_repos.png)
![Users](https://dl.gitea.com/screenshots/explore_users.png)
![Orgs](https://dl.gitea.com/screenshots/explore_orgs.png)

### Repository

![Home](https://dl.gitea.com/screenshots/repo_home.png)
![Commits](https://dl.gitea.com/screenshots/repo_commits.png)
![Branches](https://dl.gitea.com/screenshots/repo_branches.png)
![Labels](https://dl.gitea.com/screenshots/repo_labels.png)
![Milestones](https://dl.gitea.com/screenshots/repo_milestones.png)
![Releases](https://dl.gitea.com/screenshots/repo_releases.png)
![Tags](https://dl.gitea.com/screenshots/repo_tags.png)

#### Repository Issue

![List](https://dl.gitea.com/screenshots/repo_issues.png)
![Issue](https://dl.gitea.com/screenshots/repo_issue.png)

#### Repository Pull Requests

![List](https://dl.gitea.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.gitea.com/screenshots/repo_pull_request.png)
![File](https://dl.gitea.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.gitea.com/screenshots/repo_pull_request_commits.png)

#### Repository Actions

![List](https://dl.gitea.com/screenshots/repo_actions.png)
![Details](https://dl.gitea.com/screenshots/repo_actions_run.png)

#### Repository Activity

![Activity](https://dl.gitea.com/screenshots/repo_activity.png)
![Contributors](https://dl.gitea.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.gitea.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.gitea.com/screenshots/repo_recent_commits.png)

### Organization

![Home](https://dl.gitea.com/screenshots/org_home.png)

</details>
