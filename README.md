# Distributed Log Monitoring System

A centralized system to collect, process, search, and visualize logs from distributed services in real time.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Configuration](#configuration)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)

## 🎯 Overview

The Distributed Log Monitoring System is an enterprise-grade solution for centralized log management. It aggregates logs from multiple sources, enables real-time searching, and provides actionable alerts and analytics.

### Key Capabilities

- **Real-time Log Ingestion**: Handle high-volume log streams from multiple sources
- **Full-Text Search**: Search across millions of logs with advanced filters
- **Dynamic Dashboards**: Visualize log patterns, trends, and anomalies
- **Intelligent Alerting**: Automatic detection and notification of issues
- **Horizontal Scalability**: Scale components independently based on demand

## 🏗️ Architecture

### System Components

```
┌─────────────────────────────────────────────────────────────┐
│                    Log Sources                              │
│  (App Servers, Microservices, Containers, Syslog)          │
└────────────────────┬────────────────────────────────────────┘
                     │
        ┌────────────┴────────────┐
        │                         │
   ┌────▼───────┐        ┌───────▼────┐
   │  Fluentd   │        │  Logstash  │
   │  Agents    │        │  Agents    │
   └────┬───────┘        └───────┬────┘
        │                         │
        └────────────┬────────────┘
                     │
              ┌──────▼──────┐
              │   Kafka     │
              │   Streaming │
              └──────┬──────┘
                     │
        ┌────────────┼────────────┐
        │            │            │
   ┌────▼────┐  ┌────▼────┐  ┌──▼─────┐
   │   ES    │  │   API   │  │ Alert  │
   │ Indexing│  │ Processing│ │Engine  │
   └────┬────┘  └────┬────┘  └──┬─────┘
        │            │           │
   ┌────▼────────────▼───────────▼────┐
   │       API Backend (Go)             │
   └────────────┬──────────────────────┘
                │
        ┌───────▼────────┐
        │ React Dashboard│
        │   (Frontend)   │
        └────────────────┘
```

### Data Flow

```
1. Log Collection: Fluentd/Logstash collect logs from sources
2. Streaming: Kafka distributes logs for processing
3. Processing: Parse, enrich, and transform logs
4. Storage: Index logs in Elasticsearch
5. API: Serve logs via Go REST endpoints with Gorilla Mux
6. Visualization: React dashboard displays real-time data
7. Alerting: Detect anomalies and trigger notifications
8. Monitoring: Prometheus metrics collection and visualization
```

## ✨ Features

### Core Features

- ✅ Real-time log ingestion and streaming
- ✅ Full-text search with advanced filters
- ✅ Log filtering by service, level, host
- ✅ Centralized dashboard
- ✅ JSON and plaintext log support
- ✅ Request/trace ID correlation

### Advanced Features

- ✅ Error pattern detection
- ✅ Alerting via email, webhooks, Slack
- ✅ Log anomaly detection
- ✅ Saved queries and dashboards
- ✅ Service discovery and registry
- ✅ Analytics and trend visualization
- ✅ User authentication and authorization
- ✅ Multi-environment support

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gorilla Mux
- **HTTP Server**: Native Go http
- **Performance**: 2-3x faster than Python, ~50% lower memory

### Log Ingestion
- **Fluentd**: Log collection and forwarding
- **Logstash**: Advanced log processing

### Streaming & Queuing
- **Apache Kafka**: Real-time log streaming

### Search & Analytics
- **Elasticsearch**: Log indexing and full-text search

### Frontend
- **Framework**: React 18+
- **Styling**: Tailwind CSS
- **Charts**: Recharts, Chart.js
- **State Management**: Zustand
- **API Client**: Axios, React Query

### Deployment
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Container Registry**: Docker Hub / Private Registry

## 📦 Installation

### Prerequisites

- Docker & Docker Compose (for containerized setup)
- OR:
  - Go 1.21+
  - Node.js 20+
  - Elasticsearch 8.x
  - Kafka 7.x

  