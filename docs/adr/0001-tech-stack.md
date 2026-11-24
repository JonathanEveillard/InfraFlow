# ADR-0001: Tech stack

- **Status:** Accepted
- **Date:** 2026-11-23

## Context

InfraFlow is a visual control plane for a Proxmox VE homelab cluster. It must
provision and configure VMs and LXCs from a canvas, and show deployment
status, live inventory, templates, platform changes (audit) and errors.

Constraints:
- One developer, learning project; time should go into Proxmox logic, not
  generic plumbing.
- One admin user, one cluster (pve1, pve2).
- Must run as a self-hosted web app in a container.
- Background work (polling, applies, template builds) must keep running
  when no browser is open.

## Decision

| Layer        | Choice                                   |
|--------------|------------------------------------------|
| Language     | Go                                       |
| Backend core | PocketBase, embedded as a Go library     |
| Database     | SQLite (via PocketBase), `pb_data` volume|
| Frontend     | React + TypeScript, built with Vite      |
| Canvas       | React Flow                               |
| Realtime     | PocketBase realtime subscriptions        |
| Proxmox      | REST API with a scoped API token         |
| Logging      | Go `log/slog`                            |
| Packaging    | One multi-stage Dockerfile, one image    |
| Runtime      | Docker Compose, single instance          |

The frontend build is embedded into the Go binary with `go:embed`, so the
whole app ships as one executable in one container.

## Alternatives considered

- **Cross-platform desktop app (Tauri/Electron/Wails)** background work
  stops when the app closes, state is per-machine, and three platforms to
  package. See ADR-0002.
- **Postgres + hand-built auth/realtime** about one extra phase of
  non-Proxmox work and a second container. See ADR-0003.
- **Python (FastAPI) or Node backend** workable, but Go gives a single
  static binary, strong concurrency for pollers and executors, and matches
  the Proxmox/Terraform/Packer ecosystem whose code is worth reading.
- **Kubernetes deployment** no benefit for a single-instance app, and
  running inside VMs it manages risks it taking itself down. See ADR-0004.

## Consequences

**Positive**
- One binary, one container, one process: simple to build, run and back up.
- Auth, migrations, realtime and an admin UI come with PocketBase.
- Go concurrency fits the poller, executor and task tracking well.
- Domain packages stay plain Go and testable without a cluster.

**Negative**
- Single instance only (SQLite writer, in-process poller and apply lock).
- PocketBase is pre-1.0; upgrades need care.
- Secrets encryption must be built (AES-GCM).

**Follow-ups**
- Only `store`, `api`, `migrations` and `main.go` import PocketBase.
- Pin all versions (`go.mod`, `package-lock.json`, base images).
- Decide direct Proxmox API vs generated Terraform.