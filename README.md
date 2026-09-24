# GCP Calculator MCP Server (Golang)

[![Go Reference](https://pkg.go.dev/badge/github.com/Mireu-Lab/GCP-Calculator-MCP.svg)](https://pkg.go.dev/github.com/Mireu-Lab/GCP-Calculator-MCP)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go)](https://go.dev/)

> **GCP Calculator MCP Server**는 Google Cloud Platform(GCP)의 115개 전체 서비스에 대해 리전별 단가 차율, CUD(약정 할인), 이중 통화(USD & KRW)를 지원하는 고성능 **Model Context Protocol (MCP)** 서버 및 CLI 계산기 패키지입니다.

---

## ✨ 핵심 특징 (Key Features)

- 🌍 **리전별 요금 가중치 반영 (Region Multipliers)**
  - 리전별 단가 차이 자동 적용 (`us-central1`: 1.00, `asia-northeast3` (서울): 1.20, `asia-northeast1` (도쿄): 1.18, `europe-west1`: 1.10 등).
- 📉 **약정 할인 및 Spot 할인 (CUD & Spot Discounts)**
  - Spot/Preemptible 할인(~70%) 및 **1년 약정(37% 할인)**, **3년 약정(55% 할인)** 옵션 제공.
- 💵 **이중 통화 표기 (USD $ & KRW ₩)**
  - 계산 결과에 USD와 환율(1 USD = 1,350 KRW 기준)을 적용한 KRW 금액을 동시에 정밀 제공.
- ⚡ **115개 전 서비스 지원 (Universal Dynamic Evaluator)**
  - GCP Catalog DB 기반으로 115개 주요 GCP 서비스 및 커스텀 파라미터 기반 견적 자동 산출.
- 🔌 **Model Context Protocol (MCP) 완벽 지원**
  - Antigravity, Claude Desktop, Cursor 등 LLM 에이전트와 Stdio 기반 MCP 호환 tool 8종 제공:
    - `calculate_compute_engine`
    - `calculate_cloud_run`
    - `calculate_storage`
    - `calculate_gke`
    - `calculate_cloud_sql`
    - `calculate_bigquery`
    - `calculate_gcp_service`
    - `estimate_infrastructure`

---

## 🚀 빠른 시작 (Quick Start)

### 1. Installation

```bash
git clone https://github.com/Mireu-Lab/GCP-Calculator-MCP.git
cd GCP-Calculator-MCP
go build -o gcp-calculator-mcp main.go
```

### 2. MCP Server Configuration (Antigravity / Claude Desktop / Cursor)

`claude_desktop_config.json` 또는 `mcp_config.json` 설정 파일에 추가하여 에이전트 도구로 사용합니다.

```json
{
  "mcpServers": {
    "gcp-calculator": {
      "command": "/absolute/path/to/gcp-calculator-mcp",
      "args": ["server"]
    }
  }
}
```

---

## 💻 CLI 사용법 (CLI Examples)

```bash
# 1. Compute Engine VM (서울 리전 + 3년 CUD 약정 할인)
./gcp-calculator-mcp compute -machine e2-standard-4 -region asia-northeast3 -commitment 3-year

# 2. Cloud Storage (서울 리전 + Standard 500GB)
./gcp-calculator-mcp storage -class standard -gb 500 -region asia-northeast3

# 3. 115개 유니버설 GCP 서비스 동적 계산 (Sole-tenant Nodes)
./gcp-calculator-mcp service "Sole-tenant Nodes" '{"nodes": 2, "hours_per_month": 730, "region": "asia-northeast3", "commitment": "3-year"}'

# 4. MCP 서버 실행 (Stdio Mode)
./gcp-calculator-mcp server
```

---

## 🛠 MCP Tools 사양

| Tool 이름 | 설명 |
| :--- | :--- |
| `calculate_compute_engine` | Compute Engine (VM, vCPU, RAM, Boot Disk, CUD, Spot 등) 계산 |
| `calculate_cloud_run` | Cloud Run (vCPU, Memory, Requests 백만 건 단위) 계산 |
| `calculate_storage` | Cloud Storage (Standard, Nearline, Coldline, Archive GB/월) 계산 |
| `calculate_gke` | GKE Cluster (Autopilot, Standard Cluster Fee + Node Pool) 계산 |
| `calculate_cloud_sql` | Cloud SQL (MySQL, PostgreSQL, SQL Server, High Availability) 계산 |
| `calculate_bigquery` | BigQuery (Storage GB + On-Demand Query TB) 계산 |
| `calculate_gcp_service` | 115개 GCP 전체 catalog 서비스 유니버설 견적 계산 |
| `estimate_infrastructure` | 복합 인프라 아키텍처 토탈 월간 비용 한 번에 견적 |

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
