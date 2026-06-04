# System Architecture

## Overview

The Distributed Log Monitoring System is designed for scalability, reliability, and real-time processing of logs from distributed systems.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                           Sources                               │
│  (Microservices, Containers, VMs, Load Balancers, Databases)   │
└────────────────────┬────────────────────────────────────────────┘
                     │
        ┌────────────┴────────────┐
        │                         │
   ┌────▼───────┐        ┌───────▼────┐
   │  Fluentd   │        │  Logstash  │
   │  Agents    │        │  Agents    │
   │            │        │            │
   │ ┌────────┐ │        │ ┌────────┐ │
   │ │Parser  │ │        │ │Parser  │ │
   │ ├────────┤ │        │ ├────────┤ │
   │ │Filter  │ │        │ │Filter  │ │
   │ ├────────┤ │        │ ├────────┤ │
   │ │Buffer  │ │        │ │Buffer  │ │
   │ └────────┘ │        │ └────────┘ │
   └────┬───────┘        └───────┬────┘
        │                         │
        └────────────┬────────────┘
                     │
           ┌─────────▼──────────┐
           │    Kafka Topics    │
           │  (Partitioned)     │
           │                    │
           │ - logs (3 part)    │
           │ - alerts (1 part)  │
           │ - metrics (2 part) │
           └─────────┬──────────┘
                     │
        ┌────────────┼────────────────┐
        │            │                │
   ┌────▼─────┐ ┌───▼─────┐ ┌───────▼──┐
   │    Log   │ │  Alert  │ │  Metric  │
   │ Processor│ │ Detector│ │Aggregator│
   │          │ │         │ │          │
   │ ┌──────┐ │ │┌──────┐ │ │┌──────┐  │
   │ │Index │ │ ││Check │ │ ││  Sum │  │
   │ │      │ │ ││Error │ │ ││Aggr  │  │
   │ │Enrich│ │ ││  %   │ │ ││      │  │
   │ └──────┘ │ │└──────┘ │ │└──────┘  │
   └────┬─────┘ └───┬─────┘ └───────┬──┘
        │           │               │
   ┌────▼────────┐  │        ┌──────▼──────┐
   │Elasticsearch│  │        │  Database   │
   │             │  │        │  (Stats)    │
   │ Inverted    │  │        │             │
   │ Index       │  │        │ Metrics     │
   │             │  │        │ Baselines   │
   │ - logs-*    │  │        │ Thresholds  │
   │   - shards  │  │        └──────┬──────┘
   │   - docs    │  │               │
   └────┬────────┘  │               │
        │           │               │
        │      ┌────▼───────┐      │
        │      │  Database  │      │
        │      │ (PostgreSQL)       │
        │      │             │      │
        │      │ - Alerts    │      │
        │      │ - Users     │      │
        │      │ - Dashboards       │
        │      │ - Queries   │      │
        │      └────┬────────┘      │
        │           │               │
        └───────────┼───────────────┘
                    │
         ┌──────────▼──────────┐
         │    FastAPI Backend  │
         │  (REST API Layer)   │
         │                     │
         │ /api/v1/logs        │
         │ /api/v1/alerts      │
         │ /api/v1/system      │
         └──────────┬──────────┘
                    │
         ┌──────────▼──────────┐
         │  React Frontend     │
         │  (Dashboard UI)     │
         │                     │
         │ - Log Viewer        │
         │ - Alerts Panel      │
         │ - Analytics         │
         │ - Dashboards        │
         └─────────────────────┘
```

## Component Details

### 1. Log Collection Layer

**Fluentd & Logstash**:
- Collect logs from multiple sources
- Parse structured and unstructured logs
- Apply filters and transformations
- Buffer logs for reliability
- Forward to Kafka

**Key Features**:
- Multi-source support
- Format normalization
- Deduplication
- Sampling (optional)

### 2. Streaming Layer

**Apache Kafka**:
- Highly available message broker
- Partitioned topics for parallelism
- Consumer groups for processing
- Log retention policies

**Topics**:
- `logs`: Raw log events (3 partitions, 7-day retention)
- `alerts`: Alert events (1 partition, 30-day retention)
- `metrics`: System metrics (2 partitions, 14-day retention)

### 3. Processing Layer

**Log Processor**:
- Consumes from `logs` topic
- Enriches logs (GeoIP, service metadata)
- Validates and transforms
- Indexes into Elasticsearch

**Alert Detector**:
- Monitors `logs` topic
- Detects anomalies:
  - Error rate spikes
  - Keyword matches
  - Service downtime
  - Performance degradation
- Publishes alerts to `alerts` topic

**Metrics Aggregator**:
- Consumes from `metrics` topic
- Aggregates statistics
- Stores baseline data
- Updates dashboards

### 4. Storage Layer

**Elasticsearch**:
- Inverted index for full-text search
- Sharded for horizontal scalability
- Real-time analytics (aggregations)
- Time-series data (logs-YYYY.MM.DD)

**PostgreSQL**:
- Structured metadata
- Alert rules and events
- User preferences
- Dashboard configurations
- Audit logs

**Redis**:
- Query result caching
- Session storage
- Rate limiting counters
- Real-time metrics

### 5. API Layer

**FastAPI Backend**:
- RESTful API endpoints
- Request validation (Pydantic)
- Rate limiting
- Error handling
- CORS support

**Key Endpoints**:
- Logs: Search, ingest, batch operations, statistics
- Alerts: CRUD, triggering, event management
- System: Health, stats, configuration

### 6. Frontend Layer

**React Dashboard**:
- Real-time log viewer
- Advanced search filters
- Alert management interface
- Analytics and charts
- Custom dashboards

## Data Flow

### Log Ingestion Flow

```
Log Source → Fluentd/Logstash → Kafka Topic (logs)
                                    ↓
                            Log Processor
                                    ↓
                        Enrichment + Validation
                                    ↓
                        Elasticsearch + Database
                                    ↓
                            FastAPI Backend
                                    ↓
                            React Dashboard
```

### Alert Detection Flow

```
Elasticsearch ← Log Processor ← Kafka (logs)
     ↑
     │ Query for patterns
     │
Alert Detector
     │
     ├→ Error Rate Check
     ├→ Keyword Check
     ├→ Downtime Check
     ├→ Anomaly Check
     │
     └→ Kafka (alerts)
          ↓
      Notification Service
          ├→ Email
          ├→ Slack
          └→ Webhooks
          
     Database
          ↓
      Dashboard
```

## Scalability Features

### Horizontal Scaling

- **Fluentd/Logstash**: Multiple agent instances
- **Kafka**: Increased partitions for parallelism
- **Log Processors**: Multiple consumer instances
- **Elasticsearch**: Add nodes for sharding
- **FastAPI Backend**: Multiple instances behind load balancer
- **Frontend**: Served via CDN

### Performance Optimization

- **Caching**: Redis for frequently accessed data
- **Batching**: Batch log ingestion (1000s at a time)
- **Compression**: Snappy compression for Kafka
- **Indexing**: Time-series indices for faster queries
- **Aggregation**: Pre-computed statistics

### Resource Management

- **Memory**: Limits on JVM (ES, Kafka), Go goroutines
- **CPU**: Multiple concurrent Go routines
- **Disk**: Log rotation, retention policies
- **Network**: Compression, optimization

## High Availability

### Redundancy

- **Kafka**: Replicated partitions
- **Elasticsearch**: Replica shards
- **Database**: Read replicas, backups
- **Backend**: Load balanced instances
- **Frontend**: CDN distribution

### Failover

- Automatic replica failover
- Health checks and monitoring
- Graceful degradation
- Retry logic with backoff

### Disaster Recovery

- Regular backups (Elasticsearch snapshots)
- Database replication and backups
- Configuration version control
- Documentation and runbooks

## Security Architecture

### Network Security

- TLS/SSL for all communication
- Network policies in Kubernetes
- Firewall rules
- VPC isolation

### Data Security

- Encryption at rest (optional)
- Encryption in transit
- PII redaction in logs
- Access control policies

### Authentication & Authorization

- JWT tokens for API access
- Role-based access control (RBAC)
- User authentication
- Audit logging

## Monitoring & Observability

### Metrics

- Log ingestion rate
- Processing latency
- Search response time
- Error rates
- Resource utilization

### Logging

- Application logs
- System logs
- Access logs
- Error traces

### Tracing

- Distributed tracing (optional)
- Request correlation IDs
- Service call paths
- Performance profiling

### Alerting

- Threshold-based alerts
- Anomaly detection
- Multi-channel notifications
- Incident escalation

---

For deployment instructions, see DEPLOYMENT.md
