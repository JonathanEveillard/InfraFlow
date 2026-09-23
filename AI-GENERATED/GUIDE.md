# InfraFlow — Full Project Notes

*Open-source visual infrastructure control plane: Electron · Spring Boot · C++ · PostgreSQL · Proxmox · Terraform · Ansible · Prometheus · Loki · Kubernetes*

## Contents

**Part 1 — Original ChatGPT conversation**
- Q1. Project for Electron.js, PostgreSQL, Java & C++
- Q2. Visual IaC (Ansible + Terraform, n8n-style) via Proxmox
- Q3. Adding metrics like Prometheus / Grafana / Loki
- Q4. Title of the project
- Q5. How to show it on GitHub
- Q6. Should it use Java Spring Boot?
- Q7. Purpose of the app and what it does
- Q8. Kubernetes as an add-on
- Q9. Open-source title *(unanswered by ChatGPT)*

**Part 2 — Continued with Claude**
- Q9. A title that relates to open source (answered)
- Q10. Explain the entire project
- Q11. UI mockups — all 17 pages

**Part 3 — Building InfraFlow from scratch**
- Phases 0–19 with Google search keywords

---

# Part 1 — Original ChatGPT conversation

---

## Q1: What would be a project for Electron.js, PostgreSQL, Java & C++?

If you want one serious portfolio project that genuinely uses **Electron.js + PostgreSQL + Java + C++**, a good option is a desktop system-monitoring and automation platform.

### Project: Cross-Platform System Monitor & Automation Suite

A desktop app that monitors a computer, stores historical data, and can run/manage system tasks.

### How each technology fits

**Electron.js — Desktop GUI**
- Dashboard with CPU/RAM/disk/network graphs
- Process manager
- Logs and alerts
- Settings and configuration

**C++ — Low-level system agent**
- Collect CPU, RAM, disk, GPU, and network statistics
- Monitor processes
- Perform OS-level operations
- Communicate with the Electron app through a local API/socket

**Java — Backend/business layer**
- Authentication and authorization
- Task scheduling
- Alert/notification engine
- Data processing
- REST API between the desktop application and database

**PostgreSQL — Persistent storage**
- User accounts
- System metrics
- Process history
- Scheduled tasks
- Audit logs
- Alerts

### Architecture

```
                 ┌─────────────────────┐
                 │     Electron.js     │
                 │   Desktop Frontend  │
                 └──────────┬──────────┘
                            │ REST/WebSocket
                            ▼
                 ┌─────────────────────┐
                 │        Java         │
                 │   Backend Service   │
                 └───────┬─────┬───────┘
                         │     │
                    SQL  │     │ IPC/API
                         │     │
                         ▼     ▼
              ┌─────────────┐ ┌──────────────┐
              │ PostgreSQL  │ │     C++      │
              │   Database  │ │ System Agent │
              └─────────────┘ └──────────────┘
```

### Example features

**Dashboard**
- CPU usage: 37%
- RAM usage: 62%
- Disk usage: 71%
- Network throughput
- Running processes
- Temperature/GPU information where supported

**Process manager**

The C++ component could expose:

```
GET  /processes
POST /process/{pid}/terminate
```

Electron displays the results, while C++ actually interacts with the operating system.

**Historical analytics**

PostgreSQL could store:

```
system_metrics
----------------
id
timestamp
cpu_usage
ram_usage
disk_usage
network_in
network_out
gpu_usage
```

Then the Electron frontend can display:

```
CPU
100% ┤                 ╭─╮
 75% ┤       ╭──╮     ╭╯ ╰╮
 50% ┤  ╭────╯  ╰─────╯   ╰─
 25% ┤──╯
  0% └────────────────────────
       10:00   11:00   12:00
```

**Automation**

Users could create rules such as:

> If CPU > 90% for 5 minutes → notify me.

Or:

> Every day at 2 AM → run cleanup task.

Java handles the scheduling/rules, C++ performs the system-level operation, and PostgreSQL records what happened.

### Why this is a strong project

It gives you a reason to use all four technologies for different responsibilities, rather than artificially putting four languages into the same application.

It also demonstrates:
- Desktop development
- C++ systems programming
- Java backend development
- SQL/database design
- REST APIs
- IPC
- Multithreading
- Authentication
- Real-time data
- Scheduling
- Logging
- Software architecture

---

## Q2: How about a cross-platform system monitor & automation (like a visual IaC setup using Ansible + Terraform, just a visual way of doing it like n8n does it) via Proxmox?

That is a much stronger project. You could essentially build a **visual infrastructure automation platform for Proxmox**, combining the workflow experience of n8n with IaC concepts from Terraform/Ansible.

### Project concept: Visual Proxmox Infrastructure Orchestrator

> n8n + Terraform + Ansible + Proxmox, packaged as a cross-platform desktop application.

The user visually builds infrastructure and automation workflows instead of writing everything manually.

```
┌─────────────┐
│ Proxmox     │
│ Cluster     │
└──────┬──────┘
       │
       ▼
┌─────────────────┐
│ Create VM       │
│ Ubuntu 24.04    │
│ 4 CPU / 8GB RAM │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Cloud-init      │
│ Configure SSH   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Ansible         │
│ Install Docker  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Terraform       │
│ Create services │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Health Check    │
└────────┬────────┘
         │
       ┌─┴─┐
       ▼   ▼
   SUCCESS FAIL
       │     │
       ▼     ▼
    Notify  Rollback
```

The visual workflow is the product.

### How to divide the technologies

#### Electron.js — the visual IDE

Electron is responsible for the entire user-facing application, with an interface similar to n8n:

```
┌──────────────────────────────────────────────────────────┐
│ Proxmox Flow                                  ▶ Deploy   │
├──────────────┬───────────────────────────────────────────┤
│ Nodes        │                                           │
│              │       ┌──────────────┐                    │
│ Proxmox      │       │ Create VM    │                    │
│ ├─ VM        │       │ Ubuntu       │                    │
│ ├─ Container │       └──────┬───────┘                    │
│              │              │                            │
│ Terraform    │              ▼                            │
│ ├─ Resource  │       ┌──────────────┐                    │
│              │       │ Cloud Init   │                    │
│ Ansible      │       └──────┬───────┘                    │
│ ├─ Playbook  │              │                            │
│              │              ▼                            │
│ Logic        │       ┌──────────────┐                    │
│ ├─ Condition │       │ Ansible      │                    │
│ ├─ Wait      │       │ Docker       │                    │
│              │       └──────────────┘                    │
└──────────────┴───────────────────────────────────────────┘
```

You'd probably want a React-based node editor inside Electron.

#### Java — orchestration engine

Java is the core orchestration service. It could manage:
- Workflow execution
- Job queues
- Dependencies
- Credentials
- State management
- Scheduling
- Execution history
- Retry policies
- Rollbacks
- WebSocket events
- API communication

```
Workflow
   │
   ├── Create VM
   │      │
   │      └── Proxmox
   │
   ├── Wait for SSH
   │
   ├── Run Ansible
   │
   └── Health Check
```

Java turns that graph into an executable DAG. The frontend doesn't actually do the deployment — it sends:

```json
{
  "workflow": "deploy-web-server",
  "nodes": [
    {
      "id": "vm",
      "type": "proxmox.vm"
    },
    {
      "id": "ansible",
      "type": "ansible.playbook"
    }
  ]
}
```

The Java engine determines what needs to happen and in what order.

#### C++ — infrastructure/runtime layer

Instead of having C++ simply be a random system-monitoring component, make it your **native execution agent**. It could handle:
- SSH connections
- File transfers
- Process management
- Network discovery
- System metrics
- Native OS operations
- VM/container monitoring
- Secure credential handling
- Local execution
- Agent communication

A small native daemon (`proxmox-flow-agent`) runs on a Linux server:

```
                Electron
                   │
                   │ WebSocket/REST
                   ▼
             Java Engine
                   │
          ┌────────┼────────┐
          │        │        │
          ▼        ▼        ▼
      Terraform  Ansible  C++ Agent
          │        │        │
          └────────┼────────┘
                   ▼
                Proxmox
```

The C++ agent could also provide live system telemetry back to Electron.

#### PostgreSQL — the state of your infrastructure

Don't just use PostgreSQL for user accounts. Store the **desired and observed state** of your infrastructure:

```
projects
workflows
workflow_nodes
workflow_edges

proxmox_clusters
proxmox_nodes
virtual_machines
containers

deployments
deployment_steps
deployment_logs

terraform_states
ansible_runs

credentials
secrets

metrics
alerts
audit_logs
```

This enables a deployment history view:

```
Deployment #184
────────────────────────────────

✓ Create VM                 12.4s
✓ Configure networking       2.1s
✓ Cloud-init                 8.7s
✓ Wait for SSH               4.2s
✓ Ansible provisioning      31.8s
✓ Docker installation       18.3s
✓ Health check               1.4s

Total: 78.9 seconds

Status: SUCCESS
```

### Visual Terraform

Don't try to replace Terraform. Your application becomes a **visual authoring layer**.

```
        ┌─────────────────┐
        │ Proxmox Cluster │
        └────────┬────────┘
                 │
        ┌────────┴────────┐
        ▼                 ▼
  ┌───────────┐     ┌───────────┐
  │ Web VM    │     │ DB VM     │
  │ 4 CPU     │     │ 8 CPU     │
  │ 8GB RAM   │     │ 16GB RAM  │
  └─────┬─────┘     └─────┬─────┘
        │                 │
        ▼                 ▼
    Ansible            Ansible
    Docker             PostgreSQL
```

The application generates Terraform:

```hcl
resource "proxmox_virtual_environment_vm" "web" {
  name      = "web"
  node_name = "pve01"

  cpu {
    cores = 4
  }

  memory {
    dedicated = 8192
  }
}
```

The user doesn't necessarily need to see the HCL, but the flow could be:

> **Visual → Generated IaC → Plan → Apply**

### Visual Ansible

Same concept:

```
[Ubuntu VM]
     │
     ▼
[Install Docker]
     │
     ▼
[Copy docker-compose.yml]
     │
     ▼
[Start containers]
```

Generates something equivalent to:

```yaml
- hosts: web
  become: true

  tasks:
    - name: Install Docker
      ...

    - name: Copy compose file
      ...

    - name: Start application
      ...
```

This gives users both **Visual mode** and **Code mode**:

```
┌───────────────────────────────────────────┐
│ Visual | Terraform | Ansible | Logs       │
├───────────────────────────────────────────┤
│                                           │
│       visual workflow here                │
│                                           │
└───────────────────────────────────────────┘
```

Developers get the convenience of visual automation without being locked into a proprietary representation.

### Proxmox integration

Make Proxmox the initial target rather than supporting every hypervisor. The first version could support:
- Proxmox clusters
- Nodes
- QEMU VMs
- LXC containers
- Storage
- Networks
- VM templates
- Snapshots
- Cloning
- Start/stop/restart
- VM provisioning
- Cloud-init
- Resource monitoring

Eventually:

```
Infrastructure
├── Proxmox
├── Docker
├── Kubernetes
├── AWS
├── Azure
└── GCP
```

But don't start there — Proxmox alone gives you plenty to build.

### Killer feature: "Plan"

Include a Terraform-style plan. If the user changes:

```
Web VM
RAM: 8 GB → 16 GB
CPU: 4 → 6
```

The UI shows:

```
Infrastructure Plan

~ web-server
    cpu:
      4 → 6

    memory:
      8192 MB → 16384 MB


1 resource to modify

        [Cancel] [Apply]
```

For destructive operations:

```
⚠ Destructive change

The following resources will be destroyed:

    database-01

This operation cannot be automatically reversed.

                    [Cancel] [Confirm]
```

That makes the application feel like a legitimate infrastructure tool rather than an Electron GUI around API calls.

### Workflows (the n8n influence)

```
┌───────────────────┐
│ Schedule          │
└─────────┬─────────┘
          ▼
┌───────────────────┐
│ Check Proxmox     │
│ Node Health       │
└─────────┬─────────┘
          ▼
       ┌───────┐
       │ CPU > │
       │ 80%?  │
       └───┬───┘
          / \
        YES  NO
        /     \
       ▼       ▼
 [Scale VM]  [Finish]
```

**Triggers:**
- Cron
- Webhook
- Proxmox event
- VM status change
- CPU threshold
- Memory threshold
- Disk threshold
- Deployment completion
- Manual trigger

**Actions:**
- Create VM
- Destroy VM
- Snapshot VM
- Clone VM
- Run Terraform
- Run Ansible
- SSH command
- HTTP request
- Send notification
- Wait
- Conditional branch
- Approval gate

### Architecture

```
┌──────────────────────────────────────────────────────┐
│                    Electron App                      │
│                                                      │
│  Workflow Editor │ Infrastructure │ Deployments      │
│  Logs            │ Metrics        │ Settings         │
└───────────────────────┬──────────────────────────────┘
                        │
                   REST / WS
                        │
                        ▼
┌──────────────────────────────────────────────────────┐
│                  Java Control Plane                  │
│                                                      │
│ Workflow Engine │ Scheduler │ API    │ Auth          │
│ Job Queue       │ State     │ Events │ Audit         │
└───────┬──────────────┬───────────────┬───────────────┘
        │              │               │
        ▼              ▼               ▼
 PostgreSQL       Terraform        Ansible
        │              │               │
        └──────────────┼───────────────┘
                       │
                       ▼
              ┌─────────────────┐
              │     Proxmox     │
              │     Cluster     │
              └────────┬────────┘
                       │
                       ▼
                ┌─────────────┐
                │ C++ Agent   │
                │ Monitoring  │
                │ Native Ops  │
                └─────────────┘
```

### Important design decision

Avoid making Java directly execute arbitrary shell commands everywhere. Define a controlled execution abstraction:

```
Node
 ├── input
 ├── configuration
 ├── dependencies
 ├── execution
 ├── output
 └── error
```

Every visual node implements that interface:

```
ProxmoxCreateVM
TerraformApply
AnsiblePlaybook
SSHCommand
Wait
Condition
Webhook
Notification
```

This gives you an extensible node/plugin architecture. Eventually users could install:

```
@proxmox-flow/node-docker
@proxmox-flow/node-kubernetes
@proxmox-flow/node-github
@proxmox-flow/node-cloudflare
```

…and extend the workflow system without modifying the core application.

### The project to build: "Proxmox Flow"

**MVP:**
- Electron workflow editor
- PostgreSQL backend
- Java orchestration engine
- Proxmox API integration
- Visual VM creation node
- Terraform generation/execution
- Ansible playbook execution
- Live deployment logs
- Workflow persistence
- Basic C++ monitoring agent
- CPU/RAM/network graphs
- Deployment history
- Import/export workflows as JSON
- Terraform/Ansible code preview

**"Wow" version adds:**
- Terraform plan visualization
- Rollbacks
- Approval nodes
- Secrets management
- Multi-Proxmox-cluster support
- HA-aware deployment
- Live infrastructure graph
- Plugin SDK
- Remote C++ agents
- Webhooks
- Scheduled workflows
- Role-based access control

Each technology has a legitimate purpose rather than being included just to satisfy a tech-stack checklist.

---

## Q3: It would also be used as a way to show metrics of infra like Prometheus/Grafana/Loki

That makes the project much more compelling: you're not just building a visual IaC tool, you're building a **visual infrastructure control plane + observability platform**.

### Visual Infrastructure Orchestrator & Observability Platform

> n8n's workflow editor + Terraform's desired-state model + Ansible's configuration management + Proxmox + Prometheus/Grafana/Loki in one cross-platform application.

Your application becomes the **single pane of glass** for the infrastructure.

### The main UI

```
┌──────────────────────────────────────────────────────────────┐
│  INFRA FLOW                           Cluster: homelab       │
├────────────┬─────────────────────────────────────────────────┤
│            │                                                 │
│ Dashboard  │  Infrastructure                                 │
│            │                                                 │
│ Workflows  │   ┌────────────┐       ┌────────────┐           │
│            │   │ pve01      │       │ pve02      │           │
│ Resources  │   │ ● Healthy  │       │ ● Healthy  │           │
│            │   │ CPU 34%    │       │ CPU 61%    │           │
│ Metrics    │   │ RAM 52%    │       │ RAM 71%    │           │
│            │   └─────┬──────┘       └─────┬──────┘           │
│ Logs       │         │                    │                  │
│            │     ┌───┴──────┐        ┌────┴──────┐           │
│ Alerts     │     │ VM 101   │        │ VM 201    │           │
│            │     │ web      │        │ postgres  │           │
│            │     └──────────┘        └───────────┘           │
│ Deployments│                                                 │
│            │                                                 │
└────────────┴─────────────────────────────────────────────────┘
```

Clicking a VM:

```
VM 101 — web
────────────────────────────────────────────

CPU       ███████░░░  67%
Memory    █████░░░░░  48%
Disk      ████████░░  81%
Network   ↓ 142 MB/s ↑ 31 MB/s

[Metrics] [Logs] [Processes] [Network] [Configuration]

CPU usage
 100% ┤                 ╭──╮
  75% ┤       ╭────╮   ╭╯  ╰╮
  50% ┤───────╯    ╰───╯    ╰──
  25% ┤
   0% └──────────────────────────
       10:00     11:00     12:00
```

### Prometheus/Grafana/Loki become integrations

Don't reinvent Prometheus, Grafana, and Loki. Your application acts as an **observability frontend/orchestrator** around them.

#### Prometheus

The system could automatically discover infrastructure and configure monitoring:

```
Proxmox cluster
      │
      ▼
Discovery
      │
      ├── pve01
      ├── pve02
      ├── VM 101
      ├── VM 102
      └── CT 201
             │
             ▼
        Prometheus
```

The UI queries Prometheus for:
- CPU
- Memory
- Disk
- Network
- VM availability
- Node health
- Container metrics
- Application metrics

Users don't need to understand PromQL:

```
Metric:   CPU Usage
Resource: VM 101
Period:   Last 6 hours
```

The application translates that into the appropriate Prometheus query, while still providing **Advanced → PromQL** for power users.

#### Loki for logs

Instead of opening a separate Loki/Grafana interface, users click **VM 101 → Logs**:

```
┌─────────────────────────────────────────────────────────────┐
│ Logs: VM 101                                  [Live ●]      │
├─────────────────────────────────────────────────────────────┤
│ 12:42:31  nginx     200 GET /                               │
│ 12:42:31  nginx     200 GET /api/health                     │
│ 12:42:34  ssh       Accepted publickey                      │
│ 12:42:38  docker    Started container web                   │
│ 12:42:41  systemd   Started backup.service                  │
└─────────────────────────────────────────────────────────────┘
```

With filters:

```
[VM: web] [Service: nginx] [Level: ERROR] [Last: 1h]
```

And you could **link metrics → logs**. For example: CPU spiked to 94% at 12:42 → click the spike → *View logs around 12:42*. That's a genuinely useful observability workflow.

#### Grafana can still exist underneath

Make Grafana optional. The application provides its own basic dashboards, while letting users open Grafana dashboards when they need more sophisticated visualization:

```
Your Application
       │
       ├── Prometheus ──► Metrics
       │
       ├── Loki ────────► Logs
       │
       └── Grafana ─────► Advanced dashboards
```

That prevents you from spending years rebuilding Grafana.

### Automation based on observability

```
        ┌───────────────────┐
        │ Prometheus Alert  │
        │ CPU > 90%         │
        │ for 5 minutes     │
        └─────────┬─────────┘
                  │
                  ▼
        ┌───────────────────┐
        │ Get VM metrics    │
        └─────────┬─────────┘
                  │
                  ▼
             ┌────────┐
             │ CPU >  │
             │ 95% ?  │
             └───┬────┘
                / \
              YES  NO
              /     \
             ▼       ▼
      ┌───────────┐  ┌────────┐
      │ Snapshot  │  │ Finish │
      │ VM        │  └────────┘
      └─────┬─────┘
            │
            ▼
      ┌─────────────┐
      │ Run Ansible │
      │ diagnostics │
      └──────┬──────┘
             │
             ▼
      ┌─────────────┐
      │ Send Alert  │
      └─────────────┘
```

Observability isn't just "look at pretty graphs" — it becomes an **input into the automation system**.

### Three kinds of state

**Desired state** — what the user wants:

```
Web VM
4 CPU
8 GB RAM
Ubuntu
Docker
nginx
```

**Actual infrastructure state** — what Proxmox reports:

```
Web VM
4 CPU
8 GB RAM
Ubuntu
Docker
nginx
```

**Observed runtime state** — what Prometheus/Loki report:

```
CPU: 82%
RAM: 74%
Disk: 91%

nginx:
  healthy

docker:
  healthy

errors:
  3 in last 5 minutes
```

The application shows all three:

```
             Desired
                │
                ▼
          ┌───────────┐
          │ Terraform │
          └─────┬─────┘
                │
                ▼
          Actual State
                │
                ▼
          ┌────────────┐
          │ Prometheus │
          │   Loki     │
          └─────┬──────┘
                │
                ▼
         Observed State
```

### Where C++ becomes even more useful

With the observability layer, the C++ component becomes an agent (`infra-agent`) running on Linux VMs/servers, collecting information not conveniently available through Proxmox:
- CPU
- Memory
- Disk
- Processes
- Network interfaces
- Network connections
- Filesystem usage
- Service status
- Docker status
- Hardware information
- Process-level statistics

It exposes metrics in a Prometheus-compatible format:

```
infra_cpu_usage 0.67
infra_memory_used_bytes 8589934592
infra_disk_used_bytes 214748364800
infra_process_count 142
```

```
C++ Agent
    │
    ▼
Prometheus
    │
    ▼
Your Electron application
```

This gives the C++ code a real reason to exist.

### Java becomes the control plane

```
Java Control Plane
│
├── Workflow Engine
├── Infrastructure State
├── Proxmox Integration
├── Terraform Runner
├── Ansible Runner
├── Prometheus Integration
├── Loki Integration
├── Alert Engine
├── Scheduler
├── Secrets
├── Authentication
└── WebSocket Event Bus
```

### Separation of responsibilities

| Component | Role |
|---|---|
| **Electron** | IDE / control interface (not responsible for deploying) |
| **Java** | Control plane |
| **C++** | Native / agent layer |
| **PostgreSQL** | Persistent state store |
| **Prometheus / Loki** | Observability backends |
| **Terraform / Ansible** | Execution engines |
| **Proxmox** | Infrastructure platform |

### Final architecture

```
                         ┌───────────────────────┐
                         │       Electron        │
                         │                       │
                         │ Workflow Editor       │
                         │ Infrastructure        │
                         │ Metrics               │
                         │ Logs                  │
                         │ Deployments           │
                         └───────────┬───────────┘
                                     │
                              REST / WebSocket
                                     │
                                     ▼
                         ┌───────────────────────┐
                         │     Java Control      │
                         │        Plane          │
                         │                       │
                         │ Workflow Engine       │
                         │ Scheduler             │
                         │ State Manager         │
                         │ API                   │
                         │ Event Bus             │
                         └──────┬──────┬─────────┘
                                │      │
                    ┌───────────┘      └─────────────┐
                    ▼                                ▼
             ┌────────────┐                    ┌────────────┐
             │ PostgreSQL │                    │ Terraform  │
             └────────────┘                    └─────┬──────┘
                                                     │
                                               ┌─────▼─────┐
                                               │  Proxmox  │
                                               └─────┬─────┘
                                                     │
                         ┌───────────────────────────┼──────────────┐
                         │                           │              │
                         ▼                           ▼              ▼
                    ┌─────────┐                ┌─────────┐    ┌─────────┐
                    │ C++     │                │ C++     │    │ C++     │
                    │ Agent   │                │ Agent   │    │ Agent   │
                    └────┬────┘                └────┬────┘    └────┬────┘
                         │                           │              │
                         └───────────────────────────┼──────────────┘
                                                     ▼
                                               ┌────────────┐
                                               │ Prometheus │
                                               └─────┬──────┘
                                                     │
                                               ┌─────▼─────┐
                                               │   Loki    │
                                               └───────────┘
```

### Summary

This is ambitious enough for a serious portfolio/capstone project. The strongest part isn't any individual technology — it's that you're designing a small infrastructure control plane where **provisioning, configuration, observability, and automation all operate on the same visual model of the infrastructure**.

---

## Q4: Title of the project

**My pick:**

### InfraFlow
**Visual Infrastructure Orchestration & Observability Platform**

Other strong options:

- **InfraForge** — Visual Infrastructure Automation Platform
- **ProxFlow** — Visual Proxmox Infrastructure Orchestrator
- **StackFlow** — Infrastructure Automation & Observability
- **InfraPilot** — Visual Infrastructure Control Plane
- **NodeForge** — Infrastructure Provisioning & Monitoring Platform
- **OpsFlow** — Visual DevOps & Infrastructure Automation
- **InfraCanvas** — Visual Infrastructure-as-Code Platform
- **FlowOps** — Visual Infrastructure Orchestration & Observability

For a university/portfolio project, **InfraFlow** works best because it doesn't lock the project to Proxmox while still communicating the visual workflow aspect.

---

## Q5: How would you show it on GitHub?

Present InfraFlow as a serious open-source infrastructure platform rather than just "an Electron project."

> **InfraFlow** — Visual Infrastructure Orchestration & Observability Platform
>
> A cross-platform desktop application for visually designing, deploying, monitoring, and automating infrastructure across Proxmox using Terraform, Ansible, Prometheus, Loki, and a native C++ agent.

### Repository structure

```
infraflow/
│
├── README.md
├── LICENSE
├── CONTRIBUTING.md
├── SECURITY.md
├── docker-compose.yml
├── .github/
│   └── workflows/
│       ├── build.yml
│       ├── test.yml
│       └── release.yml
│
├── apps/
│   └── desktop/
│       ├── electron/
│       ├── src/
│       └── package.json
│
├── services/
│   └── control-plane/
│       ├── src/
│       ├── pom.xml
│       └── Dockerfile
│
├── agents/
│   └── infra-agent/
│       ├── src/
│       ├── include/
│       ├── CMakeLists.txt
│       └── Dockerfile
│
├── infrastructure/
│   ├── terraform/
│   ├── ansible/
│   ├── prometheus/
│   └── loki/
│
├── packages/
│   ├── workflow-schema/
│   └── node-sdk/
│
├── docs/
│   ├── architecture/
│   ├── workflows/
│   ├── deployment/
│   └── development/
│
└── examples/
    ├── homelab/
    ├── web-server/
    └── monitoring/
```

The separation makes the architecture immediately obvious:

```
Electron     → UI
Java         → Control plane
C++          → Infrastructure agent
PostgreSQL   → State
Terraform    → Provisioning
Ansible      → Configuration
Prometheus   → Metrics
Loki         → Logs
Proxmox      → Infrastructure
```

### The README

The README should lead with screenshots/GIFs, not a giant wall of text. Example README content:

> #### InfraFlow — Visual Infrastructure Orchestration & Observability Platform
>
> InfraFlow is a cross-platform desktop platform for visually designing, deploying, monitoring, and automating infrastructure. It combines Infrastructure-as-Code, configuration management, workflow automation, and observability into a single visual interface.
>
> *Build infrastructure visually. Deploy it reproducibly. Observe it in real time.*

**Features**

- 🧩 Visual infrastructure and automation workflows
- 🖥️ Proxmox VM and container management
- 🏗️ Terraform-based infrastructure provisioning
- ⚙️ Ansible-based configuration management
- 📊 Prometheus metrics integration
- 📜 Loki log aggregation
- 🔎 Real-time infrastructure monitoring
- 🔄 Workflow-based automation
- 📈 Infrastructure and application dashboards
- 📝 Deployment history and audit logs
- 🔐 Credential and secret management
- 🤖 Native C++ monitoring agent
- 🖥️ Cross-platform Electron desktop application

**Architecture**

```
                         ┌─────────────────────┐
                         │      Electron       │
                         │     Desktop App     │
                         └──────────┬──────────┘
                                    │
                              REST / WebSocket
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │   Java Control      │
                         │       Plane         │
                         └──────┬──────┬───────┘
                                │      │
                   ┌────────────┘      └────────────┐
                   ▼                                ▼
             ┌────────────┐                   ┌───────────┐
             │ PostgreSQL │                   │ Terraform │
             └────────────┘                   └─────┬─────┘
                                                    │
                                                    ▼
                                             ┌────────────┐
                                             │  Proxmox   │
                                             └─────┬──────┘
                                                   │
                              ┌────────────────────┼──────────────────┐
                              ▼                    ▼                  ▼
                         ┌─────────┐          ┌─────────┐        ┌─────────┐
                         │ C++     │          │ C++     │        │ C++     │
                         │ Agent   │          │ Agent   │        │ Agent   │
                         └────┬────┘          └────┬────┘        └────┬────┘
                              └────────────────────┼──────────────────┘
                                                   ▼
                                            ┌────────────┐
                                            │ Prometheus │
                                            └─────┬──────┘
                                                  ▼
                                             ┌─────────┐
                                             │  Loki   │
                                             └─────────┘
```

**Example workflow**

```
┌──────────────┐
│ Create VM    │
│ Ubuntu       │
│ 4 CPU / 8GB  │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Cloud-init   │
│ Configure    │
│ SSH          │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Ansible      │
│ Install      │
│ Docker       │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Health Check │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Start        │
│ Monitoring   │
└──────────────┘
```

The visual workflow can be translated into Terraform and Ansible operations while maintaining deployment state in PostgreSQL.

**Observability**

InfraFlow integrates with Prometheus and Loki to provide infrastructure-wide visibility from the same interface.

*Metrics* — monitor CPU utilization, memory utilization, disk usage, network throughput, VM availability, host health, process statistics, and application metrics.

*Logs* — query and stream logs from infrastructure directly within the application. Metrics and logs can be correlated with deployments and automation workflows to help identify infrastructure problems.

**Technology stack**

| Component | Technology | Responsibility |
|---|---|---|
| Desktop UI | Electron | Cross-platform interface |
| Control Plane | Java | Orchestration and workflow execution |
| Native Agent | C++ | System monitoring and native operations |
| Database | PostgreSQL | Persistent infrastructure state |
| Provisioning | Terraform | Infrastructure-as-Code |
| Configuration | Ansible | Server configuration |
| Metrics | Prometheus | Time-series monitoring |
| Logs | Loki | Log aggregation |
| Virtualization | Proxmox | Infrastructure platform |

**Project status:** 🚧 Early development — InfraFlow is under active development.

**Roadmap**

- [ ] Electron application shell
- [ ] Visual workflow editor
- [ ] PostgreSQL persistence
- [ ] Proxmox integration
- [ ] VM/container management
- [ ] Terraform integration
- [ ] Ansible integration
- [ ] Java workflow engine
- [ ] C++ monitoring agent
- [ ] Prometheus integration
- [ ] Loki integration
- [ ] Deployment history
- [ ] Infrastructure health dashboard
- [ ] Workflow scheduling
- [ ] Alerts and automation
- [ ] Secrets management
- [ ] Plugin/node SDK

**Development prerequisites:** Node.js, Java, CMake, a C++ compiler, PostgreSQL, Terraform, Ansible, Proxmox VE.

**Running locally**

```bash
git clone https://github.com/<your-username>/infraflow.git
cd infraflow

# Start development dependencies
docker compose up -d

# Start the control plane
cd services/control-plane
./mvnw spring-boot:run

# Start the desktop application
cd ../../apps/desktop
npm install
npm run dev
```

**License:** MIT

### GitHub's visual presentation

Make the repository's social preview image look like a product:

```
┌─────────────────────────────────────────────────────────┐
│                                                         │
│                     INFRAFLOW                           │
│                                                         │
│       Visual Infrastructure Orchestration               │
│                  & Observability                        │
│                                                         │
│    Proxmox • Terraform • Ansible • Prometheus • Loki    │
│                                                         │
│                [ dashboard screenshot ]                 │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

The first few screenshots should show:

1. **Dashboard** — infrastructure overview + health
2. **Workflow editor** — the n8n-like visual interface
3. **Infrastructure topology** — Proxmox cluster → VMs → containers
4. **Metrics view** — CPU/RAM/network time-series
5. **Logs view** — Loki integration
6. **Deployment view** — Terraform/Ansible execution and logs

That makes someone understand the project in 30 seconds.

### One thing to emphasize

Don't market it as:

> "A project using Java, C++, Electron, and PostgreSQL."

That's a technology demo. Market it as:

> "A visual infrastructure control plane for provisioning, configuring, monitoring, and automating Proxmox environments."

Then the technology stack explains *how* you built it. That framing makes the project look like real infrastructure software rather than a school assignment.

---

## Q6: Would it be better to use Java Spring Boot?

Yes. For this project, **Spring Boot** is a very good choice for the Java control plane, rather than building the backend from scratch.

### Why Spring Boot fits InfraFlow

The Java service has many responsibilities Spring Boot is well suited for:

- REST API for Electron
- WebSocket/SSE events for live deployment output
- PostgreSQL access
- Authentication/authorization
- Workflow execution
- Job scheduling
- Transaction management
- Proxmox API integration
- Terraform/Ansible process orchestration
- Prometheus/Loki API integration
- Configuration management
- Structured logging
- Testing

```
                    Electron
                       │
                 REST / WebSocket
                       │
                       ▼
             ┌───────────────────┐
             │   Spring Boot     │
             │                   │
             │ REST API          │
             │ Workflow Engine   │
             │ Job Scheduler     │
             │ Auth              │
             │ State Management  │
             │ Event System      │
             └──────┬────────────┘
                    │
       ┌────────────┼──────────────┐
       │            │              │
       ▼            ▼              ▼
 PostgreSQL     Terraform       Ansible
       │            │              │
       └────────────┼──────────────┘
                    ▼
                 Proxmox
                    │
                    ▼
              C++ Agents
                    │
             ┌──────┴──────┐
             ▼             ▼
        Prometheus        Loki
```

### Spring Boot project structure

```
control-plane/
├── src/main/java/com/infraflow/
│
├── api/
│   ├── workflow/
│   ├── infrastructure/
│   ├── deployment/
│   ├── metrics/
│   └── logs/
│
├── workflow/
│   ├── engine/
│   ├── execution/
│   ├── nodes/
│   └── scheduler/
│
├── infrastructure/
│   ├── proxmox/
│   ├── terraform/
│   └── ansible/
│
├── observability/
│   ├── prometheus/
│   └── loki/
│
├── persistence/
│   ├── entity/
│   ├── repository/
│   └── migration/
│
├── security/
│
└── config/
```

Libraries to use:

- **Spring Web** — REST API
- **Spring WebSocket** — live workflow/deployment events
- **Spring Data JPA** — PostgreSQL persistence
- **Flyway** — database migrations
- **Spring Security** — authentication/authorization
- **Quartz or Spring scheduling** — scheduled workflows
- **Testcontainers** — integration testing with PostgreSQL
- **Micrometer** — metrics for the control plane itself

Spring WebFlux is an option if you later need reactive/non-blocking workloads, but don't start there. Regular Spring MVC is simpler and perfectly reasonable for the initial architecture.

### How a deployment flows through Spring Boot

Think of Spring Boot as the control plane. When the user clicks **Deploy**, Electron sends:

```
POST /api/deployments
```

Spring Boot creates a deployment:

```
Deployment #184
       │
       ├── Terraform Plan
       │
       ├── Terraform Apply
       │
       ├── Wait for VM
       │
       ├── Ansible
       │
       └── Health Check
```

The workflow engine executes those jobs and publishes events:

```
DEPLOYMENT_STARTED
      ↓
TERRAFORM_PLAN_STARTED
      ↓
TERRAFORM_PLAN_COMPLETE
      ↓
VM_CREATED
      ↓
ANSIBLE_STARTED
      ↓
ANSIBLE_COMPLETE
      ↓
HEALTH_CHECK_PASSED
      ↓
DEPLOYMENT_COMPLETE
```

Electron receives those events over WebSocket and updates the UI in real time.

### A backend that's usable on its own

The backend can expose:

```
/api/workflows
/api/infrastructure
/api/deployments
/api/metrics
/api/logs
/api/agents
```

So later you could add other clients:

```
Electron ──┐
           │
Web UI ────┼──► Spring Boot ──► Infrastructure
           │
CLI ───────┘
```

That makes InfraFlow feel like an actual platform rather than an Electron app with a backend attached.

### Recommended stack

| Layer | Choice |
|---|---|
| Frontend | Electron + React + TypeScript |
| Backend | Java 21/25 + Spring Boot |
| Database | PostgreSQL + Flyway |
| Infrastructure | Proxmox |
| IaC | Terraform |
| Configuration | Ansible |
| Metrics | Prometheus |
| Logs | Loki |
| Native agent | C++ + CMake |
| Communication | REST + WebSocket |
| Testing | JUnit + Testcontainers |
| Packaging | Docker + native Electron builds |

Spring Boot gives you enough backend functionality to make the control plane substantial without reinventing HTTP, persistence, security, scheduling, and dependency management.

---

## Q7: So what is the purpose of this app and what does it do?

**InfraFlow is a visual infrastructure management platform that lets you design, deploy, monitor, and automate Proxmox infrastructure from one application.** Think of it as combining the ideas behind n8n, Terraform, Ansible, Prometheus, and Grafana into one visual control plane.

### What problem does it solve?

Managing a homelab or private cloud normally means jumping between tools:

```
Terraform       → Create infrastructure
Ansible         → Configure servers
Proxmox         → Manage VMs/containers
Prometheus      → Collect metrics
Grafana         → Visualize metrics
Loki            → Search logs
SSH             → Troubleshoot machines
Terminal        → Run commands
```

InfraFlow puts a visual layer over all of that:

```
                 ┌───────────────────┐
                 │     InfraFlow     │
                 │                   │
                 │ Deploy            │
                 │ Configure         │
                 │ Monitor           │
                 │ Automate          │
                 │ Troubleshoot      │
                 └─────────┬─────────┘
                           │
        ┌──────────────────┼──────────────────┐
        ▼                  ▼                  ▼
    Proxmox            Prometheus            Loki
  Infrastructure        Metrics              Logs
```

### What would you actually do with it?

To deploy a web app on Proxmox, you create a workflow:

```
┌──────────────┐
│ Create VM    │
│ Ubuntu       │
│ 4 CPU / 8GB  │
└──────┬───────┘
       ▼
┌──────────────┐
│ Cloud-init   │
│ Networking   │
└──────┬───────┘
       ▼
┌──────────────┐
│ Ansible      │
│ Install      │
│ Docker       │
└──────┬───────┘
       ▼
┌──────────────┐
│ Deploy App   │
└──────┬───────┘
       ▼
┌──────────────┐
│ Health Check │
└──────────────┘
```

You click **Deploy**, and behind the scenes Spring Boot might:

1. Generate Terraform configuration
2. Run `terraform plan`
3. Show the plan to the user
4. Run `terraform apply`
5. Wait for the VM to become available
6. Run Ansible
7. Install/configure software
8. Register monitoring
9. Run health checks
10. Report deployment status

### Then it becomes your monitoring dashboard

```
Infrastructure
────────────────────────────────────

Proxmox Cluster                 ● Healthy

├── pve01                        ● Healthy
│   ├── web-01                   ● Running
│   ├── database-01              ● Running
│   └── monitoring-01            ● Running
│
└── pve02                        ● Healthy
    ├── web-02                   ● Running
    └── backup-01                ● Running
```

Click `web-01`:

```
web-01
────────────────────────────────

CPU             42%
Memory          67%
Disk            71%
Network         ↓ 18 MB/s ↑ 4 MB/s

Containers      7
Processes       184

              [Metrics] [Logs] [Processes]
```

Metrics come from Prometheus, logs from Loki, infrastructure information from Proxmox, and the C++ agent provides additional machine-level information.

### Automation

If a web server starts running out of disk space:

```
┌────────────────────┐
│ Disk > 85%         │
│ for 10 minutes     │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│ Query Loki         │
│ Find errors        │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│ Run Ansible        │
│ Cleanup logs       │
└─────────┬──────────┘
          ▼
┌────────────────────┐
│ Check disk again   │
└─────────┬──────────┘
          ▼
    ┌───────────┐
    │ < 80% ?   │
    └─────┬─────┘
         YES
          ▼
   ┌──────────────┐
   │ Send success │
   └──────────────┘
```

Monitoring becomes an input to automation. That's what makes the project more interesting than "a GUI for Proxmox."

### The four pillars

1. **Provision** — "Create my infrastructure." Terraform + Proxmox: VMs, containers, networks, storage, snapshots, clones.
2. **Configure** — "Set up my machines." Ansible: install Docker, configure Nginx, install packages, create users, deploy applications, configure services.
3. **Observe** — "What's happening right now?" Prometheus + Loki + C++ agent: CPU, RAM, disk, network, processes, services, logs, alerts.
4. **Automate** — "Do something when something happens." The visual workflow engine: Trigger → Condition → Action → Check → Action.

### Where each technology fits

| Technology | Purpose |
|---|---|
| Electron + React | Visual desktop interface |
| Java + Spring Boot | Control plane and workflow engine |
| C++ | Native infrastructure monitoring agent |
| PostgreSQL | Infrastructure/workflow/deployment state |
| Terraform | Infrastructure provisioning |
| Ansible | Configuration and application deployment |
| Proxmox | Virtualization infrastructure |
| Prometheus | Metrics |
| Loki | Logs |

Interview answers:

- **"Why C++?"** — "The C++ component is a lightweight native agent responsible for collecting low-level system information and communicating with the control plane."
- **"Why Java?"** — "Spring Boot provides the control plane that manages workflows, infrastructure state, deployments, scheduling, authentication, and integrations."
- **"Why Electron?"** — "The application needs a rich cross-platform visual editor for infrastructure and automation workflows."

### The ultimate goal

```
                    INFRAFLOW

┌──────────────────────────────────────────────────────┐
│                                                      │
│  Proxmox Cluster                                     │
│                                                      │
│  ┌──────────┐      ┌──────────┐                      │
│  │  pve01   │──────│  pve02   │                      │
│  │  Healthy │      │  Healthy │                      │
│  └────┬─────┘      └────┬─────┘                      │
│       │                 │                            │
│    ┌──┴───┐          ┌──┴───┐                        │
│    │ VMs  │          │ VMs  │                        │
│    └──────┘          └──────┘                        │
│                                                      │
├──────────────────────────────────────────────────────┤
│ CPU 42% │ RAM 61% │ Disk 73% │ 2 Alerts │ 0 Errors   │
├──────────────────────────────────────────────────────┤
│                                                      │
│ Recent Deployment                                    │
│ ✓ web-01 deployed                                    │
│ ✓ Docker configured                                  │
│ ✓ Monitoring enabled                                 │
│                                                      │
└──────────────────────────────────────────────────────┘
```

> **InfraFlow turns infrastructure management from a collection of CLI tools into a visual, state-aware, observable automation system.**

---

## Q8: Would Kubernetes be a good add-on to this project?

Yes, but as a **second-stage feature**, not part of the MVP. Kubernetes needs a specific role, not just a spot on the tech list.

### Where Kubernetes fits

Current model:

```
                 InfraFlow
                    │
          ┌─────────┴─────────┐
          │                   │
     Provision             Observe
          │                   │
     Terraform          Prometheus/Loki
          │
       Proxmox
          │
      ┌───┴────┐
      │        │
     VMs      LXCs
      │
      ▼
   Ansible
```

With Kubernetes on top of the infrastructure:

```
                       InfraFlow
                           │
              ┌────────────┼────────────┐
              │            │            │
          Provision     Configure     Observe
              │            │            │
         Terraform      Ansible    Prometheus/Loki
              │            │            │
              └──────┬─────┘            │
                     ▼                  │
                  Proxmox               │
                     │                  │
                  ┌──┴──┐               │
                  │ VMs │               │
                  └──┬──┘               │
                     │                  │
                     ▼                  │
               Kubernetes ◄─────────────┘
                Cluster
              ┌────┼────┐
              ▼    ▼    ▼
             Pod  Pod   Pod
```

### Visual Kubernetes deployment

```
┌─────────────────────┐
│ Kubernetes Cluster  │
│ 3 Control Planes    │
│ 3 Workers           │
└──────────┬──────────┘
           ▼
┌─────────────────────┐
│ PostgreSQL          │
│ 3 replicas          │
└──────────┬──────────┘
           ▼
┌─────────────────────┐
│ Backend             │
│ 3 replicas          │
└──────────┬──────────┘
           ▼
┌─────────────────────┐
│ Frontend            │
│ 2 replicas          │
└─────────────────────┘
```

InfraFlow translates this into Kubernetes manifests/Helm configuration. The full pipeline:

```
Visual workflow → Terraform → Proxmox VMs → Ansible → Kubernetes cluster → Kubernetes workloads
```

### Kubernetes as workflow nodes

```
Infrastructure
├── Proxmox VM
├── Proxmox Container
├── Proxmox Snapshot
│
├── Kubernetes Cluster
├── Kubernetes Namespace
├── Kubernetes Deployment
├── Kubernetes Service
├── Kubernetes Ingress
├── Kubernetes Scale
│
Automation
├── Terraform
├── Ansible
├── SSH
├── HTTP Request
├── Condition
└── Wait
```

Example workflow:

```
┌───────────────┐
│ Create 3 VMs  │
│ on Proxmox    │
└───────┬───────┘
        ▼
┌───────────────┐
│ Ansible       │
│ Install K8s   │
└───────┬───────┘
        ▼
┌───────────────┐
│ K8s Cluster   │
│ Ready?        │
└───────┬───────┘
        ▼
┌───────────────┐
│ Deploy App    │
│ 3 replicas    │
└───────┬───────┘
        ▼
┌───────────────┐
│ Health Check  │
└───────────────┘
```

### Kubernetes + observability

```
Kubernetes
────────────────────────────────────

Cluster: homelab

Nodes                  3 / 3 Healthy

Pods                   27 Running
                        2 Pending
                        0 Failed

CPU                    47%
Memory                 62%

Namespaces

default                 8 pods
monitoring              7 pods
production             12 pods
```

Clicking a workload:

```
production / web
────────────────────────────────

Replicas
  Desired: 3
  Ready:   3

CPU       41%
Memory    58%

Pods

● web-7d8f9     Running
● web-7d8f1     Running
● web-7d8f2     Running

[Metrics] [Logs] [Events] [Deployment]
```

Prometheus supplies metrics, Loki supplies logs, Kubernetes supplies workload state, and InfraFlow ties them together.

### Infrastructure topology

```
                         Internet
                            │
                            ▼
                       ┌─────────┐
                       │ Ingress │
                       └────┬────┘
                            │
                       ┌────▼────┐
                       │ Service │
                       └────┬────┘
                            │
                  ┌─────────┼─────────┐
                  ▼         ▼         ▼
                Pod       Pod       Pod
                  │         │         │
                  └─────────┼─────────┘
                            │
                       ┌────▼────┐
                       │   DB    │
                       └─────────┘
```

Zoomed out: `Internet → Kubernetes → Proxmox → Physical Nodes`. InfraFlow becomes a visual map of the entire stack.

### Automation examples

Scaling the application:

```
Prometheus Alert → CPU > 80% → Check replicas → Scale Deployment (3 → 6)
    → Wait for Ready → Verify CPU → Create deployment record
```

Scaling the infrastructure underneath Kubernetes:

```
Kubernetes needs more capacity
          │
          ▼
Are worker nodes available?
          │ NO
          ▼
Create Proxmox VM → Terraform → Ansible → Join Kubernetes cluster
          │
          ▼
Kubernetes schedules workloads
```

`Application → Kubernetes → VM → Proxmox → Physical infrastructure`, all in one automation system.

### Build in stages

| Stage | Adds | Builds |
|---|---|---|
| **V1 — Proxmox** | Electron, Spring Boot, PostgreSQL, C++, Proxmox | VM/LXC management, visual workflows, Terraform, Ansible, deployment history, C++ monitoring agent |
| **V2 — Observability** | Prometheus, Loki | Metrics, logs, alerts, dashboards, infrastructure health, monitoring-driven workflows |
| **V3 — Kubernetes** | Kubernetes, Helm | Cluster discovery, node/pod visualization, workload & deployment management, K8s workflow nodes, K8s metrics/logs, application topology |
| **V4 — Infrastructure automation** | Everything combined | End-to-end workflows across all layers |

```
             ┌────────────────────┐
             │     InfraFlow      │
             └─────────┬──────────┘
                       │
              Visual Workflow
                       │
          ┌────────────┼────────────┐
          ▼            ▼            ▼
      Terraform     Ansible     Kubernetes
          │            │            │
          └────────────┼────────────┘
                       ▼
                    Proxmox
                       │
                       ▼
                  VMs / LXCs
                       │
                       ▼
                 Applications
                       │
             ┌─────────┴─────────┐
             ▼                   ▼
        Prometheus              Loki
             │                   │
             └─────────┬─────────┘
                       ▼
                  InfraFlow
                   Dashboard
```

At that point it isn't a Proxmox GUI anymore — it's a visual infrastructure control plane. Start narrow with Proxmox, then make Kubernetes one of the major layers InfraFlow can orchestrate and observe.

---

## Q9: What is a title that relates to open source?

*(ChatGPT didn't answer this one — the conversation continued with Claude in Part 2.)*

---

# Part 2 — Continued with Claude

## Q9 (answered): A title that relates to open source

The usual approach is to keep a short, memorable name and let the tagline say it's open source. Most well-known open-source tools (Grafana, n8n, Prometheus, Terraform) don't have "open" in the name.

**Pick:**

> **InfraFlow — Open-Source Visual Infrastructure Control Plane**

You keep the name you already have, and "control plane" matches how the architecture is described.

**If you want open source in the name itself:**

- **OpenPlane** — *Open-source infrastructure control plane for Proxmox & Kubernetes.* Short, and it plays on "control plane."
- **InfraCommons** — *Community-driven visual infrastructure automation.* "Commons" signals shared, community-owned software.
- **LibreFlow** — *Libre visual infrastructure orchestration & observability.* "Libre" is strongly tied to free/open-source culture.
- **OpenInfraFlow** — *Open visual IaC, automation & observability.* Very explicit, but overlaps with the OpenInfra Foundation (formerly the OpenStack Foundation).
- **HomelabOS Flow** / **OpenHomelab** — only if targeting the homelab/self-hosting community specifically.

**Avoid:** FlowForge (an existing Node-RED platform, now FlowFuse), OpenFlow (a well-known SDN protocol), anything with "OpenStack" in it. Before committing, search GitHub, npm, and Docker Hub for the name, and check the .dev/.io domain.

**Making the repo read as open source:**

GitHub "About" line:

> Open-source visual control plane to provision, configure, observe, and automate Proxmox & Kubernetes infrastructure.

Repository topics: `open-source`, `self-hosted`, `homelab`, `proxmox`, `terraform`, `ansible`, `kubernetes`, `prometheus`, `loki`, `infrastructure-as-code`, `devops`, `spring-boot`, `electron`.

Include a clear license (MIT or Apache-2.0; Apache-2.0 is common for infrastructure tools because of its patent grant), a `CONTRIBUTING.md`, a `CODE_OF_CONDUCT.md`, and "good first issue" labels. These signal an active open-source project more than the name does.

---

## Q10: Explain the entire project

### In one sentence

InfraFlow is an open-source desktop app that lets you **build, configure, monitor, and automate** Proxmox (and later Kubernetes) infrastructure by drawing workflows on a canvas, instead of switching between Terraform, Ansible, Grafana, Loki, SSH, and the Proxmox web UI.

### The problem it solves

Running a homelab or small private cloud means juggling many separate tools. You write Terraform to create VMs, write Ansible playbooks to set them up, open Grafana to see how they're doing, search Loki when something breaks, and SSH in to fix it. Each tool is good at its job, but none of them knows about the others. When CPU spikes on a VM, nothing connects that spike to the deployment that caused it or the automation that could fix it.

InfraFlow is the layer on top that ties them together. It doesn't replace Terraform, Ansible, Prometheus, or Loki. It **drives** them and presents everything in one visual interface.

### What it does: the four pillars

| Pillar | Question it answers | Powered by |
|---|---|---|
| **Provision** | "Create my infrastructure" | Terraform + Proxmox API |
| **Configure** | "Set up my machines" | Ansible |
| **Observe** | "What's happening right now?" | Prometheus, Loki, C++ agent |
| **Automate** | "Do something when something happens" | InfraFlow's own workflow engine |

The fourth pillar is what makes the project more than a Proxmox GUI. Monitoring data feeds straight into automation, so a metric crossing a threshold can trigger a workflow that fixes the problem.

### How a user experiences it

1. **You draw it.** In the workflow editor (similar to n8n), you drag in nodes: *Create VM (Ubuntu, 4 CPU, 8 GB)* → *Cloud-init (SSH, networking)* → *Ansible: Install Docker* → *Deploy App* → *Health Check*. You can switch to a code tab at any time and see the Terraform and Ansible that InfraFlow generated from your drawing.
2. **You preview the plan.** Before anything changes, InfraFlow runs `terraform plan` and shows the result visually: "1 VM to create," or "web-01: CPU 4 → 6, RAM 8 → 16 GB." Destructive changes get a red warning and require confirmation.
3. **You deploy and watch.** You click Deploy, and each node on the canvas lights up as it runs, with live logs streaming in. When it finishes, you get a record like *Deployment #184 — 7 steps, 78.9 s, SUCCESS*, saved permanently.
4. **You monitor.** The new VM appears in the infrastructure view with live CPU, RAM, disk, and network data. Click into it for graphs from Prometheus and logs from Loki. Click a CPU spike on a graph, and you jump to the logs from that moment.
5. **You automate.** You build a second workflow: *If disk > 85% for 10 minutes → search Loki for errors → run an Ansible cleanup playbook → check disk again → notify me.* From then on, InfraFlow handles that problem without you.

### Architecture

```
            ┌──────────────────────────────┐
            │  Electron + React (desktop)  │   ← the client
            └──────────────┬───────────────┘
                           │ REST + WebSocket
 ═════════════════════════ │ ════════════════════ runs on your server
                           ▼
            ┌──────────────────────────────┐
            │  Spring Boot control plane   │ ──► PostgreSQL
            │  API · workflow engine ·     │
            │  scheduler · auth · events   │
            └───┬─────────┬─────────┬──────┘
                ▼         ▼         ▼
           Terraform   Ansible   Prometheus / Loki APIs
                │         │
                ▼         ▼
            ┌──────────────────────────────┐
            │  Proxmox  →  VMs / LXCs      │
            │   (each runs a C++ agent)    │
            └──────────────┬───────────────┘
                           │ metrics & logs
                           ▼
                  Prometheus  ·  Loki
```

**Electron is only the client.** The control plane, PostgreSQL, Prometheus, and Loki run on a server (for example, a VM on your Proxmox host, started with Docker Compose). Workflows and scheduled automations keep running when your laptop is closed, and a CLI or web UI could be added later on top of the same API.

### What each piece does and why

| Component | Role | Why this technology |
|---|---|---|
| **Electron + React + TypeScript** | Workflow canvas, dashboards, log viewer, topology map | A rich, cross-platform visual editor needs a full web UI stack |
| **Java + Spring Boot** | The "brain": turns workflows into executable steps, runs Terraform/Ansible, schedules jobs, streams events, handles auth and secrets | Mature support for REST, WebSockets, persistence, security, and scheduling |
| **C++ agent** | Lightweight daemon on each VM that collects process, service, and Docker stats and exposes them in Prometheus format | Needs to be small, fast, and close to the OS, and it runs on every machine |
| **PostgreSQL** | Stores workflows, infrastructure inventory, deployment history, Terraform state references, audit logs, and alerts | The single source of truth for what exists and what happened |
| **Terraform** | Creates and changes VMs, containers, networks, and storage on Proxmox | Industry-standard desired-state provisioning with plan/apply |
| **Ansible** | Installs and configures software on the machines | Agentless, idempotent configuration management |
| **Prometheus** | Collects and stores time-series metrics (scrapes the C++ agents and Proxmox) | Standard metrics backend with PromQL |
| **Loki** (with a log shipper such as Grafana Alloy) | Aggregates logs from all machines | Cheap log storage using the same labeling model as Prometheus |
| **Grafana** (optional) | Advanced dashboards when InfraFlow's built-in ones aren't enough | Avoids rebuilding Grafana |

### Key design ideas

**Workflows become DAGs.** The canvas is saved as JSON (nodes and edges). Spring Boot converts it into a directed acyclic graph, works out execution order and dependencies, and runs each step with retries, timeouts, and rollback hooks.

**Every node follows the same contract.** Each node type (`ProxmoxCreateVM`, `TerraformApply`, `AnsiblePlaybook`, `Condition`, `Wait`, `Webhook`, …) implements one interface: inputs, configuration, execute, outputs, errors. This keeps the engine safe (no arbitrary shell commands) and extensible (new nodes can ship as plugins).

**Visual first, code always available.** The canvas generates real HCL and YAML that you can view, export, and commit to Git. Users are never locked into a proprietary format.

**Three kinds of state.** InfraFlow tracks *desired state* (what you drew), *actual state* (what Proxmox reports), and *observed state* (how it's behaving, per Prometheus and Loki). Differences show up as drift warnings or alerts and can trigger workflows.

**Event-driven UI.** The engine publishes events such as `DEPLOYMENT_STARTED`, `TERRAFORM_PLAN_COMPLETE`, `VM_CREATED`, and `HEALTH_CHECK_PASSED` over WebSocket, and the canvas updates in real time.

### Roadmap

| Version | Adds | Delivers |
|---|---|---|
| **V1: Proxmox core** | Electron, Spring Boot, PostgreSQL, C++ agent, Terraform, Ansible | Workflow editor, VM/LXC management, plan/apply, live deployment logs, history, basic agent metrics |
| **V2: Observability** | Prometheus, Loki | Dashboards, log search, linking metrics to logs, alerts, workflows triggered by monitoring |
| **V3: Kubernetes** | Kubernetes, Helm | Build a K8s cluster on Proxmox VMs from a workflow, visualize nodes/pods/services, deploy and scale workloads |
| **V4: Full-stack automation** | All layers together | e.g. *high app load → scale pods → out of capacity → create a VM → join it to the cluster* |

Keep V1 small. A working "draw a VM → plan → deploy → see its metrics" loop is worth far more than a half-finished version of everything.

### Why it works as a portfolio and open-source project

Each technology has a job it's genuinely suited for, so you can answer "why did you use X?" in one sentence for every piece of the stack. The project demonstrates desktop UI work, backend and API design, systems programming, database modelling, infrastructure-as-code, observability, and distributed-systems thinking (state, events, DAG execution). It also targets a real, active community (homelab and self-hosting users) that adopts, stars, and contributes to tools like this.

> **InfraFlow turns infrastructure management from a pile of separate CLI tools into one visual, state-aware, observable automation system.**

---

## Q11: UI mockups — all pages

All screens share one layout (sidebar left, content right), except the login screen and the workflow editor, which use the full window.

**How the screens connect:**

```
Login ──► Dashboard
            │
Sidebar ────┼── Infra ──────► Resource detail (VM/LXC)
            ├── Workflows ──► Editor ⇄ Code view ──► Plan review ──► Live run
            ├── Deployments (history)
            ├── Metrics ⇄ Logs      (click a graph spike → jump to logs at that time)
            ├── Alerts ──────► triggers Workflows
            ├── Kubernetes  (V3)
            └── Agents · Secrets · Settings
```

For V1 you only need screens 1–10, 15, 16, and 17. Metrics, Logs, and Alerts (11–13) come in V2, and Kubernetes (14) in V3.

### 1. Login / connect to server

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│                            ◆  I N F R A F L O W                              │
│              Open-source visual infrastructure control plane                 │
│                                                                              │
│              ┌────────────────────────────────────────────────┐              │
│              │ Control plane URL                              │              │
│              │ [ https://infraflow.homelab.local:8443     ]   │              │
│              │                                                │              │
│              │ Username                                       │              │
│              │ [ admin                                    ]   │              │
│              │ Password                                       │              │
│              │ [ ••••••••••••                             ]   │              │
│              │                                                │              │
│              │ [x] Remember this server                       │              │
│              │                                                │              │
│              │                [  Connect  ]                   │              │
│              └────────────────────────────────────────────────┘              │
│                                                                              │
│           First run?  [ Setup wizard: add Proxmox, Prometheus, Loki ]        │
│                                                                              │
│ Server ● reachable                                                  v0.1.0   │
└──────────────────────────────────────────────────────────────────────────────┘
```

### 2. Dashboard

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Dashboard                        Cluster: homelab ▾   ⟳ 15s    │
│─────────────│────────────────────────────────────────────────────────────────│
│▸ Dashboard  │                                                                │
│  Infra      │ ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐ │
│  Workflows  │ │ Nodes      │  │ VMs / LXCs │  │ Alerts     │  │ Deploys 24h│ │
│  Deployments│ │ 2/2 healthy│  │ 14 running │  │ ▲ 2 firing │  │ 5 ✓   0 ✗  │ │
│  Metrics    │ └────────────┘  └────────────┘  └────────────┘  └────────────┘ │
│  Logs       │                                                                │
│  Alerts     │ CPU   ████████░░░░░░░░  42%     RAM   ██████████░░░░░░  61%    │
│  Kubernetes │ Disk  ███████████░░░░░  73%     Net   ↓ 84 MB/s  ↑ 12 MB/s     │
│  Agents     │                                                                │
│  Secrets    │ ┌─ Recent deployments ───────┐  ┌─ Active alerts ────────────┐ │
│  Settings   │ │ ✓ #184  web-01       78s   │  │ ▲▲ web-01   disk 87%       │ │
│             │ │ ✓ #183  db-01       112s   │  │ ▲  pve02    RAM 91%        │ │
│ ● homelab   │ │ ✗ #182  k8s-w-3      41s   │  │                            │ │
│ v0.1.0      │ │ ✓ #181  backup-01    55s   │  │            [View all →]    │ │
│             │ └────────────────────────────┘  └────────────────────────────┘ │
│             │                                                                │
│             │ Cluster CPU (last 6h)                                          │
│             │  100┤                                                          │
│             │   50┤        ╭──╮           ╭─╮                                │
│             │    0└──────────────────────────────────────────────            │
│             │      06:00      08:00      10:00      12:00                    │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 3. Infrastructure (inventory & drift)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Infrastructure            [Tree]  Graph      [+ New resource]  │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│▸ Infra      │ NAME                  TYPE   STATUS      CPU   RAM             │
│  Workflows  │ ▾ homelab             clust  ● healthy   38%   58%             │
│  Deployments│   ▾ pve01             node   ● online    34%   52%             │
│  Metrics    │     ├ 101 web-01      VM     ● running   42%   67%             │
│  Logs       │     ├ 102 db-01       VM     ● running   18%   71%             │
│  Alerts     │     └ 201 monitoring  LXC    ● running   11%   40%             │
│  Kubernetes │   ▾ pve02             node   ● online    61%   71%             │
│  Agents     │     ├ 103 web-02      VM     ● running   37%   55%  ▲ drift    │
│  Secrets    │     ├ 104 k8s-cp-1    VM     ● running   22%   48%             │
│  Settings   │     └ 105 backup-01   VM     ○ stopped    –     –              │
│             │                                                                │
│ ● homelab   │ ┌────────────────────────────────────────────────────────────┐ │
│ v0.1.0      │ │ ▲ Drift on web-02                                          │ │
│             │ │   Desired: 4 cores   Actual (Proxmox): 6 cores             │ │
│             │ │                                                            │ │
│             │ │   [Reconcile to desired]  [Accept actual]  [Ignore]        │ │
│             │ └────────────────────────────────────────────────────────────┘ │
│             │                                                                │
│             │ Storage: local-lvm 61% · nas-nfs 34%     Networks: vmbr0, vmbr1│
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 4. Resource detail (VM)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Infra › pve01 › web-01 (VM 101)   ● running   [▶][■][⟳][⋯]     │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │ [Overview]  Metrics  Logs  Processes  Config  History          │
│▸ Infra      │                                                                │
│  Workflows  │ ┌─ Summary ──────────────────┐  ┌─ Live ─────────────────────┐ │
│  Deployments│ │ OS       Ubuntu 24.04      │  │ CPU   ███████░░░  67%      │ │
│  Metrics    │ │ Size     4 vCPU · 8 GB     │  │ RAM   █████░░░░░  48%      │ │
│  Logs       │ │ Disk     64 GB local-lvm   │  │ Disk  ████████░░  81%      │ │
│  Alerts     │ │ IP       10.0.0.21         │  │ Net   ↓ 18 MB/s ↑ 4 MB/s   │ │
│  Kubernetes │ │ Agent    ● v0.1.0          │  │ Containers 7 · Procs 184   │ │
│  Agents     │ │ Uptime   12d 4h            │  │                            │ │
│  Secrets    │ └────────────────────────────┘  └────────────────────────────┘ │
│  Settings   │                                                                │
│             │ CPU usage                               [1h]  6h  24h  7d      │
│ ● homelab   │  100┤                    ╭──╮                                  │
│ v0.1.0      │   50┤─────────╮  ╭───────╯  ╰────────╮   ◀ click → logs @ time │
│             │    0└─────────────────────────────────────────                 │
│             │                                                                │
│             │ Recent events                                                  │
│             │ 12:42  ✓ Deployment #184 succeeded                             │
│             │ 12:40  ● Snapshot pre-deploy-184 created                       │
│             │ 09:15  ▲ Disk > 80% (resolved)                                 │
│             │ [Snapshot]  [Clone]  [Open console]  [Edit in workflow]        │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 5. Workflows (list)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Workflows          [+ New workflow]  [Import JSON]   ⌕ [     ] │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ NAME                 TRIGGER          LAST RUN    STATUS       │
│▸ Workflows  │ deploy-web-server    manual           2h ago      ✓ success    │
│  Deployments│ disk-cleanup         alert: disk>85%  yesterday   ✓ success    │
│  Metrics    │ nightly-snapshots    cron 0 2 * * *   10h ago     ✓ success    │
│  Logs       │ k8s-add-worker       webhook          3d ago      ✗ failed     │
│  Alerts     │ scale-on-load        alert: cpu>80%   never       ○ draft      │
│  Kubernetes │                                                                │
│  Agents     │ Start from a template                                          │
│  Secrets    │ ┌────────────┐  ┌────────────┐  ┌────────────┐  ┌────────────┐ │
│  Settings   │ │ Web server │  │ Postgres VM│  │ K3s cluster│  │ Monitoring │ │
│             │ │ VM+Docker  │  │ + backups  │  │ 3 VMs      │  │ Prom+Loki  │ │
│ ● homelab   │ └────────────┘  └────────────┘  └────────────┘  └────────────┘ │
│ v0.1.0      │                                                                │
│             │ Row actions:  [Open]  [Run ▶]  [Duplicate]  [Export]  [Delete] │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 6. Workflow editor (the n8n-style canvas)

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ ← deploy-web-server   ● saved                   [ Plan ]  [ Deploy ▶ ]       │
│ [Visual]   Terraform   Ansible   JSON   Runs                                 │
├───────────────┬──────────────────────────────────────┬───────────────────────┤
│ NODES         │                                      │ PROPERTIES            │
│ ⌕ filter      │          ┌──────────────┐            │ Create VM (Proxmox)   │
│▾ Triggers     │          │ ▶ Manual     │            │ ───────────────────── │
│  Manual       │          └──────┬───────┘            │ Name                  │
│  Cron         │          ╔══════▼═══════╗            │ [web-03            ]  │
│  Webhook      │          ║ Create VM    ║ ◀ selected │ Node     [pve01    ▾] │
│  Alert        │          ║ web-03       ║            │ Template              │
│▾ Proxmox      │          ╚══════╤═══════╝            │ [ubuntu-24.04     ▾]  │
│  Create VM    │          ┌──────▼───────┐            │ CPU [4 ]   RAM [8 GB] │
│  Create LXC   │          │ Cloud-init   │            │ Disk [32 GB]          │
│  Snapshot     │          └──────┬───────┘            │ Network               │
│  Clone        │          ┌──────▼───────┐            │ [vmbr0 · DHCP     ▾]  │
│▾ Config       │          │ Ansible      │            │ Tags [web, prod]      │
│  Ansible      │          │ docker.yml   │            │                       │
│  Terraform    │          └──────┬───────┘            │ ── Execution ──────── │
│  SSH command  │          ┌──────▼───────┐            │ Retries   [2]         │
│▾ Logic        │          │ Health check │            │ Timeout   [10m]       │
│  Condition    │          └──┬────────┬──┘            │ On fail   [rollback▾] │
│  Wait         │           ok│        │fail           │                       │
│  Approval     │        ┌────▼───┐ ┌──▼───────┐       │                       │
│▾ Notify       │        │ Notify │ │ Rollback │       │ [Delete node]         │
│  ntfy/Email   │        └────────┘ └──────────┘       │                       │
├───────────────┴──────────────────────────────────────┴───────────────────────┤
│ Zoom 100%  ▫ minimap     6 nodes · 7 edges · ✓ valid     Last run: #184 ✓    │
└──────────────────────────────────────────────────────────────────────────────┘
```

### 7. Code view (generated Terraform / Ansible)

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ ← deploy-web-server   ● saved                   [ Plan ]  [ Deploy ▶ ]       │
│  Visual   [Terraform]   Ansible   JSON   Runs                                │
│──────────────────────────────────────────────────────────────────────────────│
│  main.tf  │ variables.tf │ outputs.tf       Generated from canvas · read-only│
│                                                                              │
│   1  resource "proxmox_virtual_environment_vm" "web_03" {                    │
│   2    name      = "web-03"                                                  │
│   3    node_name = "pve01"                                                   │
│   4    tags      = ["web", "prod"]                                           │
│   5                                                                          │
│   6    cpu    { cores = 4 }                                                  │
│   7    memory { dedicated = 8192 }                                           │
│   8                                                                          │
│   9    clone { vm_id = 9000 }     # ubuntu-24.04 template                    │
│  10  }                                                                       │
│                                                                              │
│──────────────────────────────────────────────────────────────────────────────│
│  [Copy]  [Download .zip]  [Commit to Git…]  [Eject: edit as code]            │
└──────────────────────────────────────────────────────────────────────────────┘
```

### 8. Plan review (before apply)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Plan · deploy-web-server                       generated 12:39 │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ + create   vm.web_03             (pve01 · 4 vCPU · 8 GB)       │
│▸ Workflows  │ ~ update   vm.web_01                                           │
│  Deployments│              cpu.cores        4  →  6                          │
│  Metrics    │              memory        8192  →  16384                      │
│  Logs       │ - destroy  vm.old_test           (VM 199)                      │
│  Alerts     │                                                                │
│  Kubernetes │ Summary:  1 to add  ·  1 to change  ·  1 to destroy            │
│  Agents     │                                                                │
│  Secrets    │ ┌────────────────────────────────────────────────────────────┐ │
│  Settings   │ │ ▲ DESTRUCTIVE CHANGE                                       │ │
│             │ │   old-test (VM 199) and its disks will be deleted.         │ │
│ ● homelab   │ │   A snapshot will NOT be kept.                             │ │
│ v0.1.0      │ │                                                            │ │
│             │ │   Type the name to confirm:  [                  ]          │ │
│             │ └────────────────────────────────────────────────────────────┘ │
│             │                                                                │
│             │ Ansible:  2 playbooks · 14 tasks · 1 host (web-03)             │
│             │                                                                │
│             │ [Cancel]     [Download plan]     [Snapshot first ☑]    [Apply] │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 9. Live deployment run

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Deployment #185 · deploy-web-server · 12:40:02     ● running   │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ STEPS                                                          │
│  Workflows  │ ✓ Terraform plan        3.1s                                   │
│▸ Deployments│ ✓ Terraform apply      12.4s                                   │
│  Metrics    │ ✓ Wait for SSH          4.2s                                   │
│  Logs       │ ● Ansible: docker.yml   21s…                                   │
│  Alerts     │ ○ Health check                                                 │
│  Kubernetes │ ○ Register monitoring                                          │
│  Agents     │                                                                │
│  Secrets    │ ████████████████░░░░░░░░░░░░  4 / 6                            │
│  Settings   │                                                                │
│             │ ┌─ Live output  [Follow ●] [Copy] ───────────────────────────┐ │
│ ● homelab   │ │ 12:40:41 TASK [docker : install packages]                  │ │
│ v0.1.0      │ │ 12:40:52 changed: [web-03]                                 │ │
│             │ │ 12:40:53 TASK [docker : enable service]                    │ │
│             │ │ 12:40:55 ok: [web-03]                                      │ │
│             │ │ 12:40:56 TASK [app : copy compose file] ▌                  │ │
│             │ └────────────────────────────────────────────────────────────┘ │
│             │                                                                │
│             │ [Cancel run]    [Open canvas]    [Open web-03 →]               │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 10. Deployments (history & audit)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Deployments     Workflow [all ▾]  Status [all ▾]  Range [7d ▾] │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ #    WORKFLOW            BY      STARTED       TIME   RESULT   │
│  Workflows  │ 185  deploy-web-server   admin   today 12:40   1m44   ● running│
│▸ Deployments│ 184  deploy-web-server   admin   today 10:02   78s    ✓ success│
│  Metrics    │ 183  postgres-vm         admin   yest. 18:11   112s   ✓ success│
│  Logs       │ 182  k8s-add-worker      webhook yest. 03:00   41s    ✗ failed │
│  Alerts     │ 181  nightly-snapshots   cron    yest. 02:00   55s    ✓ success│
│  Kubernetes │                                                                │
│  Agents     │                                                                │
│  Secrets    │ ┌─ #182 details ─────────────────────────────────────────────┐ │
│  Settings   │ │ Failed at: Ansible › join cluster  (exit 2)                │ │
│             │ │ Error: kubelet unable to reach 10.0.0.30:6443              │ │
│ ● homelab   │ │ Rollback: ✓ VM 108 destroyed · ✓ state restored            │ │
│ v0.1.0      │ │                                                            │ │
│             │ │ [View full log]  [Diff vs #179]  [Re-run]                  │ │
│             │ └────────────────────────────────────────────────────────────┘ │
│             │                                                                │
│             │ Audit: every run stores who, what, plan, output & final state  │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 11. Metrics explorer (Prometheus)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Metrics                 [Simple]  PromQL      Range [6h ▾]     │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ Resource [web-01 ▾]  Metric [CPU usage ▾]  Compare [+ add]     │
│  Workflows  │                                                                │
│  Deployments│ %                                                              │
│▸ Metrics    │ 100┤                                                           │
│  Logs       │  75┤                 ╭─╮                                       │
│  Alerts     │  50┤      ╭──╮      ╭╯ ╰╮            ╭──╮                      │
│  Kubernetes │  25┤──────╯  ╰──────╯   ╰────────────╯  ╰────────              │
│  Agents     │   0└──────────────────────────────────────────────             │
│  Secrets    │     07:00     08:00     09:00     10:00     11:00              │
│  Settings   │                                                                │
│             │ ■ web-01   □ web-02   □ db-01                                  │
│ ● homelab   │                                                                │
│ v0.1.0      │ Generated query:                                               │
│             │   avg(infra_cpu_usage{vm="web-01"}) * 100                      │
│             │                                                                │
│             │ [Save to dashboard]  [Create alert]  [Open in Grafana]         │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 12. Logs explorer (Loki)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Logs            Source [web-01 ▾]  Level [≥ warn ▾]  [Live ●]  │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ Query  {vm="web-01", level=~"warn|error"} |= "timeout"         │
│  Workflows  │                                                                │
│  Deployments│ volume ▁▁▂▁▁▁▃▇█▅▂▁▁▁▁▂▁▁   (last 1h)                          │
│  Metrics    │                                                                │
│▸ Logs       │ 12:42:31  nginx    ERR   upstream timed out (110) /api/orders  │
│  Alerts     │ 12:42:31  nginx    ERR   upstream timed out (110) /api/cart    │
│  Kubernetes │ 12:42:34  app      WARN  db pool exhausted, waiting 5s         │
│  Agents     │ 12:42:38  docker   WARN  container api restarting (1/5)        │
│  Secrets    │ 12:42:41  app      ERR   timeout calling postgres:5432         │
│  Settings   │                                                                │
│             │                                                                │
│ ● homelab   │ ┌─ Selected line ────────────────────────────────────────────┐ │
│ v0.1.0      │ │ 12:42:34 app WARN db pool exhausted, waiting 5s            │ │
│             │ │ labels: vm=web-01 service=app container=api                │ │
│             │ │                                                            │ │
│             │ │ [±5 min context]  [Metrics @ 12:42]  [New alert]           │ │
│             │ └────────────────────────────────────────────────────────────┘ │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 13. Alerts & rules

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Alerts             [Firing 2]  Rules 6  Silenced 1  History    │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ SEV  ALERT        RESOURCE  SINCE  AUTOMATION                  │
│  Workflows  │ ▲▲   Disk > 85%   web-01    12m    → disk-cleanup  ● running   │
│  Deployments│ ▲    RAM > 90%    pve02     4m     → notify only               │
│  Metrics    │                                                                │
│  Logs       │ [Acknowledge]   [Silence 1h]   [Open resource]                 │
│▸ Alerts     │                                                                │
│  Kubernetes │                                                                │
│  Agents     │ ┌─ Rule editor ──────────────────────────────────────────────┐ │
│  Secrets    │ │ IF    [disk_used_pct ▾]  [ > ]  [ 85 ]  %                  │ │
│  Settings   │ │ FOR   [ 10m ]                                              │ │
│             │ │ ON    [ tag: web ▾ ]                                       │ │
│ ● homelab   │ │                                                            │ │
│ v0.1.0      │ │ THEN  run workflow   [ disk-cleanup ▾ ]                    │ │
│             │ │ AND   notify         [ ntfy: homelab ▾ ]                   │ │
│             │ │                                                            │ │
│             │ │                               [Cancel]  [Save rule]        │ │
│             │ └────────────────────────────────────────────────────────────┘ │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 14. Kubernetes (V3)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Kubernetes      Cluster [homelab-k3s ▾]   Namespace [prod ▾]   │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ Nodes 3/3 ●     Pods 27 running · 2 pending · 0 failed         │
│  Workflows  │                                                                │
│  Deployments│ WORKLOAD   KIND         READY  CPU   MEM                       │
│  Metrics    │ web        Deployment   3/3    41%   58%                       │
│  Logs       │ api        Deployment   2/3    63%   71%   ▲ 1 pending         │
│  Alerts     │ postgres   StatefulSet  3/3    22%   64%                       │
│▸ Kubernetes │                                                                │
│  Agents     │ Topology (web)                                                 │
│  Secrets    │   Internet → Ingress → svc/web ─┬─ pod web-7d8f9 ●             │
│  Settings   │                                 ├─ pod web-7d8f1 ●             │
│             │                                 └─ pod web-7d8f2 ●             │
│ ● homelab   │                                                                │
│ v0.1.0      │ Runs on Proxmox VMs:  104 k8s-cp-1 · 106 k8s-w-1 · 107 k8s-w-2 │
│             │                                                                │
│             │ [Scale]  [Restart]  [Logs]  [Events]  [Add worker VM →]        │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 15. Agents (C++ infra-agent)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Agents                                   [+ Install agent]     │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ HOST         VERSION  STATUS     LAST SEEN  OVERHEAD           │
│  Workflows  │ web-01       0.1.0    ● online   2s         0.3% CPU · 6 MB    │
│  Deployments│ web-02       0.1.0    ● online   1s         0.2% CPU · 6 MB    │
│  Metrics    │ db-01        0.1.0    ● online   3s         0.4% CPU · 7 MB    │
│  Logs       │ backup-01    0.0.9    ○ offline  2d         update available   │
│  Alerts     │                                                                │
│  Kubernetes │                                                                │
│▸ Agents     │ ┌─ Install ──────────────────────────────────────────────────┐ │
│  Secrets    │ │ 1. Pick target:  [web-03 ▾]   or   [x] via Ansible         │ │
│  Settings   │ │ 2. One-time token: ••••••••••••  [Copy]                    │ │
│             │ │ 3. Or run on the host:                                     │ │
│ ● homelab   │ │    sudo infra-agent enroll --server <url> --token <t>      │ │
│ v0.1.0      │ └────────────────────────────────────────────────────────────┘ │
│             │                                                                │
│             │ Agents expose /metrics for Prometheus and ship logs to Loki.   │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 16. Secrets

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Secrets                           [+ Add secret]   [Rotate]    │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ NAME             TYPE        USED BY            UPDATED        │
│  Workflows  │ proxmox-api      API token   4 workflows        3d ago         │
│  Deployments│ ssh-deploy-key   SSH key     6 workflows        12d ago        │
│  Metrics    │ k3s-kubeconfig   File        k8s integration    1mo ago        │
│  Logs       │ ntfy-token       Token       2 alert rules      1mo ago        │
│  Alerts     │                                                                │
│  Kubernetes │ Values are never shown · encrypted at rest · every read audited│
│  Agents     │                                                                │
│▸ Secrets    │ Nodes reference secrets by name:  {{ secret.ssh-deploy-key }}  │
│  Settings   │                                                                │
│             │                                                                │
│ ● homelab   │                                                                │
│ v0.1.0      │                                                                │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

### 17. Settings (integrations, users, plugins)

```
┌─────────────┬────────────────────────────────────────────────────────────────┐
│ ◆ INFRAFLOW │ Settings   [Integrations]  Users & roles  Notify  Plugins      │
│─────────────│────────────────────────────────────────────────────────────────│
│  Dashboard  │                                                                │
│  Infra      │ INTEGRATION  ENDPOINT                    STATUS                │
│  Workflows  │ Proxmox      https://pve01:8006          ● connected   [Edit]  │
│  Deployments│ Prometheus   http://10.0.0.5:9090        ● connected   [Edit]  │
│  Metrics    │ Loki         http://10.0.0.5:3100        ● connected   [Edit]  │
│  Logs       │ Grafana      –                           ○ optional [Connect]  │
│  Alerts     │ Kubernetes   kubeconfig: homelab-k3s     ● connected   [Edit]  │
│  Kubernetes │ Terraform    on control-plane host       ● detected            │
│  Agents     │ Ansible      on control-plane host       ● detected            │
│  Secrets    │                                                                │
│▸ Settings   │ PLUGINS (node packs)                                           │
│             │ @infraflow/node-docker       1.2.0   ✓ installed               │
│ ● homelab   │ @infraflow/node-cloudflare   0.4.1   [Install]                 │
│ v0.1.0      │ @infraflow/node-github       0.3.0   [Install]                 │
│             │                                                                │
│             │ Theme [System ▾]    Language [English ▾]    Telemetry [Off]    │
└─────────────┴────────────────────────────────────────────────────────────────┘
```

---

# Part 3 — Building InfraFlow from scratch (Google search roadmap)

Each step has a goal and the Google queries to research it. Build in order: every phase depends on the previous one.

**Search tips:** put exact phrases in quotes (`"proxmox" "cloud-init" template`); limit to official docs with `site:` (`site:docs.spring.io websocket`); remove noise with `-` (`electron react -angular`); use Google's **Tools → Any time → Past year** filter so you get current versions; add `github` or `example` to find working code.

## Phase 0 — Lab setup

**Goal:** a Proxmox host with an Ubuntu cloud-init template and an API token.

- `proxmox ve install guide`
- `proxmox nested virtualization vmware workstation` *(if you have no spare machine)*
- `proxmox create ubuntu cloud-init template qm`
- `proxmox api token permissions terraform`
- `proxmox vmbr0 bridge network setup`
- `proxmox lxc container template download`

## Phase 1 — Repository and tooling

**Goal:** a monorepo with `apps/`, `services/`, `agents/`, `packages/`, `docs/`.

- `monorepo structure java node c++ example`
- `github repository template open source README best practices`
- `choose a license MIT vs Apache 2.0`
- `CONTRIBUTING.md template open source`
- `docker compose postgres development environment`
- `conventional commits` and `semantic versioning`

## Phase 2 — Spring Boot control plane skeleton

**Goal:** a running API with a health endpoint and config.

- `spring initializr`
- `spring boot 3 rest controller tutorial`
- `spring boot application.yml profiles`
- `spring boot actuator health endpoint`
- `java 21 records virtual threads spring boot`
- `spring boot global exception handler ControllerAdvice`
- `springdoc openapi swagger ui spring boot 3`

## Phase 3 — Database and persistence

**Goal:** PostgreSQL schema for workflows, resources, deployments, and logs.

- `spring data jpa postgresql tutorial`
- `flyway migration spring boot`
- `postgresql jsonb column spring boot hibernate` *(storing workflow graphs)*
- `database design workflow engine tables nodes edges`
- `postgresql audit log table design`
- `testcontainers postgresql spring boot test`

## Phase 4 — Proxmox integration

**Goal:** list nodes/VMs and start/stop VMs from Java.

- `proxmox ve api documentation`
- `proxmox api /nodes/{node}/qemu status current`
- `proxmox api token authorization header PVEAPIToken`
- `spring boot RestClient example` or `spring webclient tutorial`
- `proxmox api self signed certificate java`
- `proxmox api clone vm from template`

## Phase 5 — Terraform generation and execution

**Goal:** turn a canvas VM node into Terraform, then run plan/apply from Java.

- `bpg proxmox terraform provider`
- `terraform proxmox_virtual_environment_vm clone example`
- `terraform json configuration syntax .tf.json` *(easier to generate than HCL)*
- `terraform plan -out show -json parse`
- `terraform state backend pg postgresql`
- `java ProcessBuilder stream output real time`
- `terraform workspace per environment`

## Phase 6 — Ansible execution

**Goal:** run playbooks against new VMs and capture structured results.

- `ansible playbook install docker ubuntu`
- `ansible dynamic inventory from json`
- `ansible json callback plugin stdout_callback`
- `ansible-runner` *(and compare with calling ansible-playbook directly)*
- `ansible wait_for_connection after vm creation`
- `ansible roles directory structure best practices`

## Phase 7 — Workflow engine (the core)

**Goal:** JSON graph → DAG → ordered execution with retries and rollback.

- `topological sort directed acyclic graph java`
- `detect cycle in directed graph java`
- `workflow engine design pattern node executor`
- `strategy pattern plugin architecture java`
- `spring boot async executor ThreadPoolTaskExecutor`
- `saga pattern compensation rollback` *(for rollback nodes)*
- `n8n workflow json format` *(for inspiration)*
- `quartz scheduler spring boot cron jobs`

## Phase 8 — Live events

**Goal:** stream deployment progress and logs to the UI.

- `spring boot websocket stomp tutorial`
- `spring boot server sent events SseEmitter`
- `websocket vs server sent events`
- `event driven architecture spring ApplicationEventPublisher`

## Phase 9 — Electron desktop app

**Goal:** an Electron + React + TypeScript shell with a sidebar and routing.

- `electron vite react typescript template` *(search `electron-vite`)*
- `electron contextBridge preload security best practices`
- `electron ipcMain ipcRenderer invoke`
- `react router sidebar layout`
- `tanstack query react api fetching`
- `tailwind css dark mode` or `shadcn ui`
- `electron store settings persist`

## Phase 10 — Visual workflow editor

**Goal:** the n8n-style canvas with a node palette and properties panel.

- `react flow tutorial` *(package: `@xyflow/react`)*
- `react flow custom nodes`
- `react flow drag and drop sidebar`
- `react flow save and restore json`
- `zustand state management react flow`
- `react hook form zod dynamic form schema` *(node properties panel)*
- `monaco editor react` *(read-only Terraform/Ansible code tab)*

## Phase 11 — Plan review and deployment screens

**Goal:** show Terraform plan diffs and live run output.

- `terraform plan json resource_changes actions`
- `react diff viewer`
- `xterm.js react` *(live terminal output)*
- `confirm destructive action ux pattern type name to confirm`

## Phase 12 — C++ agent

**Goal:** a small Linux daemon exposing `/metrics` in Prometheus format.

- `modern cmake project structure`
- `vcpkg` or `conan c++ package manager`
- `read cpu usage /proc/stat c++`
- `read memory /proc/meminfo c++`
- `linux /proc/[pid]/stat process list c++`
- `prometheus-cpp exposer example`
- `cpp-httplib` *(if you write the endpoint yourself)*
- `prometheus exposition format text`
- `systemd service unit file daemon`
- `c++ static binary linux musl` *(easy distribution)*

## Phase 13 — Prometheus integration (V2)

**Goal:** scrape agents and query metrics from Spring Boot.

- `prometheus docker compose setup`
- `prometheus file_sd_configs` *(the control plane writes targets)*
- `prometheus pve exporter` *(Proxmox host metrics)*
- `prometheus http api query_range`
- `promql rate avg by examples`
- `recharts line chart time series react` or `uplot react`

## Phase 14 — Loki integration (V2)

**Goal:** ship logs from VMs and search them in the app.

- `grafana loki docker compose`
- `grafana alloy loki config` *(log shipper; Promtail is deprecated)*
- `logql examples filter json`
- `loki http api query_range`
- `loki tail websocket api`

## Phase 15 — Alerts → workflows (V2)

**Goal:** metric thresholds trigger workflows.

- `prometheus alertmanager webhook receiver`
- `alertmanager webhook payload format`
- `spring boot webhook endpoint hmac signature verify`
- `ntfy self hosted notifications api`

## Phase 16 — Security and secrets

**Goal:** login, roles, and encrypted credentials.

- `spring security jwt authentication spring boot 3`
- `spring security role based access control`
- `encrypt secrets at rest java AES-GCM`
- `spring vault` or `hashicorp vault spring boot` *(optional)*
- `electron safeStorage` *(store tokens on the desktop)*

## Phase 17 — Kubernetes (V3)

**Goal:** build k3s on Proxmox VMs, then manage workloads.

- `k3s ansible playbook install cluster`
- `k3s high availability embedded etcd`
- `fabric8 kubernetes client java` or `kubernetes java client official`
- `helm install from java` or `helm cli automation`
- `kube-prometheus-stack helm`
- `kubernetes ingress nginx k3s traefik`
- `kubernetes horizontal pod autoscaler`

## Phase 18 — Testing, CI/CD, and packaging

**Goal:** automated builds and downloadable installers.

- `junit 5 mockito spring boot testing`
- `playwright electron testing`
- `googletest cmake tutorial`
- `github actions matrix build windows mac linux`
- `electron-builder github releases auto update`
- `spring boot docker image jib` or `spring boot buildpacks`
- `github actions build cmake c++`

## Phase 19 — Open-source launch

**Goal:** a repo people can understand in 30 seconds and contribute to.

- `good first issue labels github`
- `github social preview image size`
- `record gif screen for readme` *(e.g. search `peek` or `screentogif`)*
- `docusaurus` or `mkdocs material` *(docs site)*
- `r/homelab r/selfhosted posting rules` *(where to announce it)*
- `show hn guidelines`

### Build order at a glance

```
Lab ─► Repo ─► Spring Boot ─► DB ─► Proxmox API ─► Terraform ─► Ansible
  ─► Workflow engine ─► Events ─► Electron ─► Canvas ─► Plan/Run screens
  ─► C++ agent            ══════════ V1 done ══════════
  ─► Prometheus ─► Loki ─► Alerts→Workflows  ══ V2 ══
  ─► Security ─► Kubernetes                  ══ V3 ══
  ─► Tests/CI/Packaging ─► Launch
```
