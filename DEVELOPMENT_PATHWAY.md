# 🚀 Distributed Log Monitoring System - Development Pathway
## Docker-First Development Approach

This guide shows how to develop and build this project entirely within Docker containers. **No local installations needed**—just Docker and your code editor. Each commit represents meaningful progress in the system.

**Development Philosophy:** Everything runs in containers. Code is mounted as volumes so you edit locally, changes are instantly reflected in containers.

---

## 📊 Project Overview

**Name:** Distributed Log Monitoring System  
**Purpose:** Enterprise-grade centralized log aggregation, searching, and alerting  
**Tech Stack:**
- **Backend:** Go (REST API with Gorilla Mux, Zap logging, Prometheus metrics)
- **Frontend:** React + TypeScript (Vite, Tailwind CSS)
- **Data Pipeline:** Kafka (streaming), Elasticsearch (search/indexing)
- **Collectors:** Fluentd & Logstash (log collection)
- **Orchestration:** Docker & Docker Compose, Kubernetes manifests

**Key Metrics:**
- ~3000 LOC Backend (Go)
- ~800 LOC Frontend (React/TS)
- ~1500 LOC Config/Infrastructure
- ~15 Docker containers in full stack

**Development Requirements:**
- ✅ Docker & Docker Compose
- ✅ Text Editor (VS Code, Vim, etc.)
- ❌ NO Go, Node.js, or system-level package managers needed

---

## 🎯 Professional Commit Strategy

### Commit Philosophy
Each commit should be **atomic, meaningful, and self-contained**. Someone reading your commits should see:
1. Infrastructure setup (Docker-first)
2. Incremental feature development (inside containers)
3. Testing and quality practices
4. Kubernetes deployment specifications
5. Performance and scalability considerations

### Commit Message Format
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `chore:` Infrastructure, tooling, Docker setup
- `feat:` New feature
- `fix:` Bug fix
- `refactor:` Code restructuring (no behavior change)
- `test:` Adding/updating tests
- `docs:` Documentation
- `perf:` Performance improvements

**Example:**
```
feat(api): implement log search endpoint with full-text filtering

- Added LogHandler.SearchLogs() with Elasticsearch query builder
- Implemented pagination (offset/limit) for large result sets
- Added request validation and error handling
- Returns structured JSON with metadata

Developed and tested inside Docker container.
Addresses: Feature requirement #5
```

---

## 📋 Phase-by-Phase Development Plan (Docker-First)

### PHASE 1: Docker Infrastructure Setup (Commits 1-4)
*Goal:* Set up complete Docker development environment. Developers just run `docker-compose up` and code.

**Commit 1: Project Initialization & Git Setup**
```
chore: initialize distributed log monitoring project with Docker-first approach

- Create directory structure for backend, frontend, config, docs
- Add README with project vision and Docker setup instructions
- Add .gitignore for Go, Node, Docker
- Create LICENSE file (MIT)
- Document Docker-first development approach

Development Setup:
- Only requirement: Docker & Docker Compose
- All work happens inside containers with volume mounts
- No local Go, Node.js, or system-level tools needed
```
**Files:**
- `README.md` - Project overview with Docker setup
- `.gitignore` - Git ignore rules
- `LICENSE` - MIT license
- `docs/ARCHITECTURE.md` - High-level design
- `.dockerignore` - Docker ignore rules

**Commit 2: Docker Compose for Full Stack**
```
chore(docker): set up Docker Compose with all services for local development

Services in docker-compose.yml:
- backend (Go API) - port 8000 - with volume mount /app for live coding
- frontend (Node + Vite dev server) - port 3000 - with volume mount /app
- elasticsearch - port 9200 - persistent volume
- kafka - port 9092 - persistent volume
- zookeeper - port 2181 - for Kafka coordination
- postgres - port 5432 - for metadata (optional)

Features:
- Volume mounts for hot-reloading code
- Network isolation between services
- Health checks for all services
- Environment variables in .env.example
- Automatic dependency startup order
- Seed data initialization

Development Workflow:
1. docker-compose up
2. Edit code in your editor (mounted volumes)
3. Changes auto-reload in containers
4. View logs: docker-compose logs -f backend
```
**Files:**
- `docker/docker-compose.yml`
- `.env.example` - Environment variables for docker-compose
- `docker/.dockerignore`

**Commit 3: Backend Development Container**
```
chore(docker): create optimized Go development container

Dockerfile (docker/Dockerfile.backend.dev):
- Base: golang:1.21-alpine
- Install air (live reload tool)
- Volume mount: /app for source code
- Expose: 8000 for API, 6060 for pprof (debugging)
- CMD: air - watches for changes and rebuilds automatically

Features:
- Hot reloading: save file → instant restart
- Debugger support (Delve) on port 2345
- Go modules cached in Docker volume
- All Go tools available inside container
- No Go installation needed on host machine
```
**Files:**
- `docker/Dockerfile.backend.dev` - Development image
- `backend_go/.air.toml` - Air config for hot reload
- Update `docker-compose.yml` with backend service

**Commit 4: Frontend Development Container**
```
chore(docker): create Node development container with Vite

Dockerfile (docker/Dockerfile.frontend.dev):
- Base: node:20-alpine
- Install dependencies during build
- Volume mount: /app for source code
- Expose: 3000 for Vite dev server, 5173 for alternative
- CMD: npm run dev - starts Vite dev server with hot reload (HMR)

Features:
- Hot module replacement (HMR): instant updates on code change
- Full Node.js tooling inside container
- npm/yarn/pnpm available
- Source maps for debugging
- No Node.js installation needed on host machine

Development Experience:
- Edit React component → page updates instantly in browser
- Full TypeScript support with intellisense
```
**Files:**
- `docker/Dockerfile.frontend.dev` - Development image
- `frontend/package.json` - Node project with dev server script
- Update `docker-compose.yml` with frontend service

---

### PHASE 2: Backend API Development in Docker (Commits 5-14)
*Goal:* Build REST API inside the backend container with hot reload

**Commit 5: Backend Project Structure & Configuration**
```
feat(backend): initialize Go module and project structure inside Docker

Inside backend container (docker-compose exec backend bash):
- go mod init github.com/yourusername/distributed-log-monitoring
- Create cmd/backend/main.go entry point
- Create internal/ directory structure
- Add Makefile with targets (build, test, lint)
- Create internal/config/config.go with environment-based configuration

Configuration System:
- Read from environment variables
- Sensible defaults for development
- Override via .env file in docker-compose
- Supports: API_PORT, ELASTICSEARCH_HOST, KAFKA_BROKERS, LOG_LEVEL

Development Workflow:
1. docker-compose up (builds backend container, starts hot reload)
2. Edit go files in backend_go/
3. Changes auto-rebuild in container (via air)
4. View output: docker-compose logs -f backend
```
**Files:**
- `backend_go/go.mod` - Go module definition
- `backend_go/go.sum` - Dependency lock file
- `backend_go/cmd/backend/main.go` - Entry point
- `backend_go/Makefile` - Build targets
- `backend_go/internal/config/config.go` - Configuration loading

**Commit 6: Structured Logging with Zap**
```
feat(logging): implement structured logging with Zap inside container

Inside backend container:
- go get go.uber.org/zap
- Create internal/logger/logger.go
- Initialize logger with JSON format (production) or text (development)
- Add context-aware logging helpers

Features:
- Structured fields for better log queries
- Configurable levels: debug, info, warn, error
- Output to stdout (monitored by docker-compose)
- Performance-optimized for high-volume logging

Usage:
logger.Info("message", zap.String("key", "value"))
logger.Error("error occurred", zap.Error(err))
```
**Files:**
- `backend_go/internal/logger/logger.go`
- Update `cmd/backend/main.go` to initialize logger

**Commit 7: HTTP Router with Gorilla Mux**
```
feat(api): establish HTTP router with middleware foundation

Inside backend container:
- go get github.com/gorilla/mux
- Create internal/api/router.go
- Configure routes and middleware
- Add health check endpoint (/health)
- Implement JSON response formatting

Middleware:
- Request logging (log all API calls)
- CORS support (for frontend requests)
- Request/response timing for metrics
- Recovery from panics

API Endpoints started:
- GET /health - Health check (returns {status: "ok"})

Testing locally:
curl http://localhost:8000/health
```
**Files:**
- `backend_go/internal/api/router.go`
- `backend_go/internal/api/helpers.go` - Response formatting utilities

**Commit 8: Data Models & Type Definitions**
```
feat(models): define core data structures for logs and alerts

Created inside container:
- Create internal/models/types.go
- Define Log, LogEntry, Alert, AlertRule structs
- Add JSON tags for marshaling/unmarshaling
- Add validation methods

Data Structures:
type Log struct {
  Timestamp time.Time
  Service   string
  Level     string // ERROR, WARN, INFO, DEBUG
  Message   string
  TraceID   string
  Fields    map[string]interface{}
}

type Alert struct {
  ID        string
  Name      string
  Service   string
  Condition string
  Threshold float64
  Enabled   bool
}

Testing:
- JSON serialization/deserialization
- Validation logic
```
**Files:**
- `backend_go/internal/models/types.go`
- `backend_go/internal/models/validation.go` (optional)

**Commit 9: Elasticsearch Client Integration**
```
feat(services): implement Elasticsearch client inside backend container

Inside backend container:
- go get github.com/elastic/go-elasticsearch/v8
- Create internal/services/elasticsearch_service.go
- Implement connection pooling and health checking
- Create query builder for log searches

Features:
- Connection management with retry logic
- Index operations (create, delete, check existence)
- Query building for flexible searches
- Error handling and logging

Elasticsearch Integration:
- Connects to elasticsearch service in docker-compose
- Checks ES health on startup (with retries)
- Logs successful connections
- Ready for log indexing and searching

Testing:
curl http://localhost:9200/_health
docker-compose exec backend go test ./internal/services...
```
**Files:**
- `backend_go/internal/services/elasticsearch_service.go`
- `backend_go/internal/services/clients.go` - Client initialization

**Commit 10: Kafka Producer & Consumer Setup**
```
feat(worker): implement Kafka consumer for log streaming

Inside backend container:
- go get github.com/segmentio/kafka-go
- Create internal/worker/kafka_consumer.go
- Implement consumer group management
- Add graceful shutdown handling

Features:
- Consumer group for distributed processing
- Offset management (auto-commit)
- Error handling and reconnection
- Metrics for consumer lag
- Partition auto-discovery

Kafka Integration:
- Connects to kafka service in docker-compose (via zookeeper)
- Creates consumer group for log processing
- Logs consumer status and errors
- Ready for processing incoming logs

Testing:
docker-compose exec kafka kafka-topics --list --bootstrap-server localhost:9092
docker-compose logs -f backend | grep kafka
```
**Files:**
- `backend_go/internal/worker/kafka_consumer.go`
- `backend_go/internal/services/kafka_service.go`

**Commit 11: Log Ingestion Endpoint**
```
feat(api): implement log ingestion endpoint

Inside backend container:
- Add POST /api/logs/ingest endpoint
- Accept JSON log entries
- Validate and transform logs
- Store to Elasticsearch
- Queue to Kafka for processing

Endpoint Features:
- Accept single or batch log entries
- Auto-parse JSON logs
- Add timestamp, source IP
- Transform to standard Log structure
- Persist to Elasticsearch
- Publish to Kafka topic

Example Request:
POST /api/logs/ingest
{
  "service": "auth-service",
  "level": "ERROR",
  "message": "Failed login attempt"
}

Testing inside container:
curl -X POST http://localhost:8000/api/logs/ingest \
  -H "Content-Type: application/json" \
  -d '{"service":"api","level":"INFO","message":"request received"}'
```
**Files:**
- `backend_go/internal/api/log_handler.go` - Log endpoints
- `backend_go/internal/services/log_service.go` - Log business logic
- Update `internal/api/router.go` with new route

**Commit 12: Log Search Endpoint**
```
feat(api): implement log search endpoint with full-text filtering

Inside backend container:
- Add POST /api/logs/search endpoint
- Implement Elasticsearch DSL query builder
- Add filtering (service, level, time range)
- Add pagination (offset/limit)
- Add sorting (timestamp desc, relevance)

Search Features:
- Full-text search in messages
- Filter by service name
- Filter by log level
- Time range filtering
- Pagination for large result sets
- Result sorting

Example Request:
POST /api/logs/search
{
  "query": "error",
  "service": "payment",
  "level": "ERROR",
  "from": "2024-01-01T00:00:00Z",
  "to": "2024-01-02T00:00:00Z",
  "limit": 50,
  "offset": 0
}

Example Response:
{
  "logs": [...],
  "total_hits": 1234,
  "search_time_ms": 45,
  "query_id": "q-12345"
}

Testing:
curl -X POST http://localhost:8000/api/logs/search \
  -H "Content-Type: application/json" \
  -d '{"query":"error","limit":10}'
```
**Files:**
- Update `backend_go/internal/api/log_handler.go`
- Update `backend_go/internal/services/elasticsearch_service.go`

**Commit 13: Alert System Core**
```
feat(alert): implement alert rule engine and notifications

Inside backend container:
- Create internal/services/alert_service.go
- Implement alert rule evaluation logic
- Add alert state management (firing, resolved)
- Create alert trigger mechanism

Alert Features:
- Rule-based detection (error rate, frequency)
- Threshold evaluation
- State management (new, firing, resolved, acknowledged)
- Alert history tracking
- Notification channel abstraction

Rule Example:
{
  "name": "High Error Rate",
  "service": "payment-service",
  "condition": "error_rate > 5%",
  "duration": "5m",
  "severity": "critical"
}

Testing in container:
docker-compose exec backend go test ./internal/services -run TestAlert
```
**Files:**
- `backend_go/internal/services/alert_service.go`
- Update `backend_go/internal/models/types.go` with AlertRule

**Commit 14: Alert Management API Endpoints**
```
feat(api): implement alert management REST endpoints

Inside backend container:
- POST /api/alerts - Create new alert rule
- GET /api/alerts - List all alert rules
- GET /api/alerts/{id} - Get specific alert
- PUT /api/alerts/{id} - Update alert rule
- DELETE /api/alerts/{id} - Delete alert rule
- GET /api/alerts/history - Alert event history

Features:
- CRUD operations for alert rules
- Enable/disable alerts
- Trigger history with timestamps
- Audit logging of changes
- Error handling and validation

Example - Create Alert:
POST /api/alerts
{
  "name": "High CPU Usage",
  "service": "backend",
  "condition": "cpu > 80%",
  "duration": "5m",
  "channels": ["email"]
}

Testing:
curl http://localhost:8000/api/alerts
curl -X POST http://localhost:8000/api/alerts \
  -H "Content-Type: application/json" \
  -d '{...}'
```
**Files:**
- `backend_go/internal/api/alert_handler.go`
- Update `internal/api/router.go` with alert routes

---

### PHASE 3: Metrics & Observability (Commits 15-17)
*Goal:* Add monitoring to the backend

**Commit 15: Prometheus Metrics**
```
feat(metrics): add Prometheus metrics and health endpoints

Inside backend container:
- go get github.com/prometheus/client_golang
- Create internal/metrics/prometheus.go
- Expose /metrics endpoint
- Add custom metrics

Metrics Exposed:
- http_requests_total (by endpoint, method, status)
- http_request_duration_seconds (p50, p95, p99)
- elasticsearch_query_duration_seconds
- kafka_messages_processed_total
- logs_ingested_total
- alerts_triggered_total

Health Check Endpoint:
GET /health - Returns {status: "healthy", checks: {...}}

Testing:
curl http://localhost:8000/metrics | grep http_requests
curl http://localhost:8000/health
```
**Files:**
- `backend_go/internal/metrics/prometheus.go`
- `backend_go/internal/api/system_handler.go` - Health/metrics endpoints
- Update `internal/api/router.go`

**Commit 16: Request Validation & Error Handling**
```
refactor(api): standardize error handling and request validation

Inside backend container:
- Create centralized error types
- Add input validation middleware
- Standardize JSON error responses
- Add request rate limiting infrastructure

Error Response Format:
{
  "error": "INVALID_REQUEST",
  "message": "Required field 'service' missing",
  "request_id": "req-12345",
  "timestamp": "2024-01-01T00:00:00Z"
}

Validation:
- Required field checking
- Type validation
- Size/length limits
- Format validation (email, timestamp, etc.)

Testing:
curl -X POST http://localhost:8000/api/logs/search \
  -H "Content-Type: application/json" \
  -d '{}' # Invalid, should get error response
```
**Files:**
- Update `backend_go/internal/api/helpers.go`
- `backend_go/internal/middleware/validation.go` (new)

**Commit 17: Graceful Shutdown & Logging**
```
feat(backend): implement graceful shutdown and comprehensive logging

Inside backend container:
- Capture SIGTERM, SIGINT signals
- Close connections cleanly (Elasticsearch, Kafka)
- Drain message queues gracefully
- Flush metrics before exit
- Log shutdown process

Shutdown Sequence:
1. Receive signal
2. Stop accepting new requests
3. Wait for in-flight requests (30s timeout)
4. Close Elasticsearch client
5. Close Kafka consumer/producer
6. Flush metrics
7. Exit with code 0

Application Logging:
- Log all API requests/responses
- Log Elasticsearch queries
- Log Kafka consumer activity
- Structured fields (trace_id, request_id, service)

Testing:
docker-compose exec backend kill -TERM 1
# Should see graceful shutdown logs
```
**Files:**
- Update `backend_go/cmd/backend/main.go`
- Add logging throughout all packages

---

### PHASE 4: Frontend Development in Docker (Commits 18-25)
*Goal:* Build React dashboard with instant HMR inside container

**Commit 18: Frontend Project Setup**
```
chore(frontend): initialize React + TypeScript project in Docker

Inside frontend container:
- npm create vite@latest . -- --template react-ts
- Install Tailwind CSS and dependencies
- Configure TypeScript with strict mode
- Set up dev and build scripts

Container-based Development:
- Edit TSX/CSS files locally (mounted volume)
- Vite dev server auto-reloads in browser (port 3000)
- TypeScript compilation happens in container
- All Node.js tooling available inside

Directory Structure Created:
src/
  ├── App.tsx
  ├── main.tsx
  ├── globals.css
  └── components/

Testing:
Browser: http://localhost:3000 (auto-refreshes on save)
Container: npm run dev (already running)
```
**Files:**
- `frontend/package.json` - Dependencies and scripts
- `frontend/tsconfig.json` - TypeScript config (strict)
- `frontend/vite.config.ts` - Vite configuration
- `frontend/tailwind.config.js` - Tailwind setup
- `frontend/postcss.config.js` - PostCSS setup
- `frontend/index.html` - Entry HTML
- `frontend/src/main.tsx` - React entry point

**Commit 19: Dashboard Layout & Styling**
```
feat(dashboard): create responsive dashboard layout with Tailwind

Inside frontend container:
- Create App.tsx with main layout
- Header with navigation
- Sidebar with menu items
- Main content area
- Dark mode CSS variables
- Responsive grid system

Layout Structure:
┌─────────────────────────────┐
│  Header (Logo, Nav)         │
├──────────┬──────────────────┤
│          │                  │
│ Sidebar  │  Main Content    │
│ (Menu)   │  (Cards, Charts) │
│          │                  │
└──────────┴──────────────────┘

Features:
- Tailwind dark mode support
- Mobile responsive design
- Accessible navigation
- Clean component structure

Testing:
Browser: http://localhost:3000 (live updates on file save)
```
**Files:**
- `frontend/src/App.tsx` - Main app component
- `frontend/src/globals.css` - Global styles
- Update `frontend/src/main.tsx`

**Commit 20: Components - Stats Card**
```
feat(components): build reusable stats card component

Inside frontend container:
- Create components/StatsCard.tsx
- Display metric with value and label
- Optional trend indicator (up/down)
- Loading and error states
- Responsive sizing

Component Props:
interface StatsCardProps {
  title: string
  value: number
  unit?: string
  trend?: number // percentage change
  loading?: boolean
  error?: string
}

Example Usage:
<StatsCard
  title="Logs Ingested"
  value={150000}
  unit="/hour"
  trend={+5.2}
/>

Testing:
- Edit component, see changes instantly in browser
- Try different prop combinations
```
**Files:**
- `frontend/src/components/StatsCard.tsx`

**Commit 21: Components - Log Viewer**
```
feat(components): build log viewer with filtering and search

Inside frontend container:
- Create components/LogViewer.tsx
- Display logs in table format
- Columns: timestamp, service, level, message
- Color-coded log levels (ERROR=red, WARN=yellow)
- Expandable row details
- Filtering options

Features:
- Service filter dropdown
- Log level checkboxes
- Time range picker
- Search/filter input
- Copy to clipboard
- Pagination controls
- Responsive table

Example Usage in Dashboard:
<LogViewer
  logs={logs}
  onFilter={(filters) => searchLogs(filters)}
  loading={isLoading}
/>

Testing:
- Live filter updates as you type
- Click table rows to expand
```
**Files:**
- `frontend/src/components/LogViewer.tsx`

**Commit 22: Components - Alerts Panel**
```
feat(components): create alert rules management panel

Inside frontend container:
- Create components/AlertsPanel.tsx
- Display alert rules in table
- Create new alert form
- Edit/delete functionality
- Enable/disable toggle
- Alert history view

Features:
- Alert rule CRUD
- Form validation
- Channel selection (email, webhook)
- Threshold input
- Condition builder UI
- Rule enable/disable toggle

Form Fields:
- Alert name
- Service selector
- Condition (dropdown)
- Threshold value
- Duration
- Notification channels
- Enabled checkbox

Testing:
- Create, edit, delete alerts in UI
- Form validation on submit
```
**Files:**
- `frontend/src/components/AlertsPanel.tsx`

**Commit 23: Dashboard Page Integration**
```
feat(pages): create main dashboard page

Inside frontend container:
- Create pages/Dashboard.tsx
- Arrange components in grid layout
- Add Stats Cards section
- Add LogViewer section
- Add AlertsPanel section
- Add filters section

Dashboard Layout:
┌─ Header ─────────────────┐
│ Stats: Logs/h, Alerts/h  │
├──────────────────────────┤
│ Filters: Service, Level  │
├──────────────────────────┤
│ Recent Logs (LogViewer)  │
├──────────────────────────┤
│ Alerts (AlertsPanel)     │
└──────────────────────────┘

Features:
- Stats auto-refresh every 30s
- Filters persist in state
- Loading states for all sections
- Error handling with user feedback

Testing:
Browser: http://localhost:3000 (Dashboard loads)
```
**Files:**
- `frontend/src/pages/Dashboard.tsx`
- Update `frontend/src/App.tsx` to import Dashboard

**Commit 24: Custom Hook - useStore**
```
feat(hooks): implement global state management with useStore

Inside frontend container:
- Create hooks/useStore.ts
- Manage dashboard filters
- Manage alert rules state
- Manage UI state (theme, loading)
- Local storage persistence

State Structure:
{
  filters: {
    service: string
    level: string[]
    dateRange: { from, to }
  }
  alerts: AlertRule[]
  stats: {
    logsPerHour: number
    alertsPerHour: number
    errorCount: number
  }
  theme: 'light' | 'dark'
  loading: { [key]: boolean }
}

Usage Example:
const { filters, setFilters, stats } = useStore()

Testing:
- Filter changes update immediately
- State persists on page reload
```
**Files:**
- `frontend/src/hooks/useStore.ts`

**Commit 25: API Service Layer**
```
feat(services): create API client for backend communication

Inside frontend container:
- Create services/api.ts
- HTTP client setup (fetch wrapper)
- Error handling and retry logic
- Type-safe API calls
- Request/response interceptors

API Client Functions:
- fetchLogs(filters, pagination)
- searchLogs(query, filters)
- createAlert(rule)
- updateAlert(id, rule)
- deleteAlert(id)
- fetchAlerts()
- getSystemHealth()

Example Usage:
const logs = await api.searchLogs({
  query: "error",
  service: "payment",
  limit: 50
})

Error Handling:
- Retry on network errors
- Handle HTTP error codes
- Type-safe responses using TypeScript interfaces

Testing:
curl http://localhost:8000/api/logs/search (backend is running)
```
**Files:**
- `frontend/src/services/api.ts`
- `frontend/src/types/api.ts` (response types)

---

### PHASE 5: Backend-Frontend Integration (Commits 26-30)
*Goal:* Connect frontend to backend API, add real data flow

**Commit 26: Connect Dashboard to Backend**
```
feat(integration): integrate dashboard with backend API

Inside frontend container (npm runs dev, HMR enabled):
- Connect Dashboard component to API
- Fetch stats from backend (/health, /metrics)
- Fetch recent logs from backend
- Auto-refresh every 30 seconds
- Add loading and error states
- Toast notifications for errors

Integration Points:
- StatsCard fetches from /api/stats
- LogViewer fetches from /api/logs/search
- AlertsPanel fetches from /api/alerts

Example Flow:
1. Dashboard mounts
2. Call api.getMetrics()
3. Update stats state
4. Render StatsCards with real data
5. Set interval for auto-refresh (30s)
6. User changes filter
7. Call api.searchLogs(filters)
8. Update LogViewer with results

Testing:
Browser: http://localhost:3000
Backend: docker-compose logs -f backend
Watch real data flow from API to UI
```
**Files:**
- Update `frontend/src/pages/Dashboard.tsx`
- Add API calls and useEffect hooks

**Commit 27: Log Search Integration**
```
feat(search): implement full log search in frontend

Inside frontend container:
- LogViewer component sends searches to backend
- Search input triggers API call
- Filters apply to query (service, level, time range)
- Results update in real-time
- Pagination handles large result sets

Search Workflow:
1. User types in search box
2. Debounce 500ms to avoid too many requests
3. Call api.searchLogs() with query and filters
4. Display results in table
5. User clicks pagination
6. Load next page from API

Features:
- Debounced search input
- Loading skeleton while searching
- Result count display
- Empty state message
- Error handling

Testing:
- Type "error" in search → see results
- Change service filter → results update
- Pagination works
```
**Files:**
- Update `frontend/src/components/LogViewer.tsx` with API integration

**Commit 28: Alert Management Integration**
```
feat(alerts): implement alert CRUD operations in frontend

Inside frontend container:
- AlertsPanel fetches existing alerts from /api/alerts
- Create new alert via POST /api/alerts
- Edit alert via PUT /api/alerts/{id}
- Delete alert via DELETE /api/alerts/{id}
- Real-time list updates

Alert Workflow:
1. Load existing alerts on mount
2. Display in table
3. User clicks "Create Alert"
4. Show form modal
5. User submits form
6. POST to /api/alerts
7. Refresh alert list
8. Show success toast

Testing:
- Create alert → appears in list
- Edit alert → changes reflect
- Delete alert → removed from list
```
**Files:**
- Update `frontend/src/components/AlertsPanel.tsx` with API calls
- Update `frontend/src/hooks/useStore.ts` with alert state

**Commit 29: Real-time Updates & Auto-refresh**
```
feat(realtime): add auto-refresh and WebSocket prep

Inside frontend container:
- Stats auto-refresh every 30 seconds
- Logs auto-refresh every 60 seconds
- Graceful handling of stale data
- Loading indicators during refresh
- Connection status indicator

Auto-refresh Implementation:
- useEffect with setInterval
- Skip refresh if data is recent
- Visual indicator when refreshing
- User can manually refresh
- Configurable refresh intervals

Example:
useEffect(() => {
  const interval = setInterval(() => {
    fetchStats()
  }, 30000) // 30 seconds
  return () => clearInterval(interval)
}, [])

Testing:
- Watch data update automatically
- Leave dashboard open, see fresh logs
```
**Files:**
- Update `frontend/src/pages/Dashboard.tsx`
- Update `frontend/src/hooks/useStore.ts`

**Commit 30: Error Handling & User Feedback**
```
feat(ux): implement comprehensive error handling and notifications

Inside frontend container:
- Toast notifications for errors
- User-friendly error messages
- Retry buttons for failed operations
- Connection status indicator
- Loading skeletons for better UX

Error Scenarios Handled:
- Backend unreachable → show error, retry button
- Invalid input → show validation errors
- API errors → show specific error message
- Network timeout → show retry option
- Concurrent request conflict → show warning

Example Toast Notifications:
- "✓ Alert created successfully"
- "✗ Failed to fetch logs. Retry?"
- "⚠ Connection lost. Retrying..."

Testing:
- Stop backend (docker-compose down backend)
- See error messages in frontend
- Click retry → reconnects
```
**Files:**
- Create `frontend/src/components/Toast.tsx`
- Create `frontend/src/hooks/useToast.ts`
- Update components to use toast notifications

---

### PHASE 6: Testing in Docker (Commits 31-35)
*Goal:* Add automated tests running inside containers

**Commit 31: Backend Unit Tests - API Layer**
```
test(api): add unit tests for HTTP handlers

Inside backend container (docker-compose exec backend bash):
- go get github.com/stretchr/testify/assert
- Create internal/api/handlers_test.go
- Test log search endpoint
- Test alert CRUD endpoints
- Test error handling
- Test request validation

Test Examples:
- TestSearchLogsWithFilters()
- TestSearchLogsInvalidInput()
- TestCreateAlertSuccess()
- TestDeleteAlertNotFound()

Run Tests:
docker-compose exec backend go test -v ./internal/api
docker-compose exec backend go test -cover ./...

Target Coverage: 80% of api/ package
```
**Files:**
- `backend_go/internal/api/handlers_test.go`

**Commit 32: Backend Unit Tests - Services**
```
test(services): add unit tests for business logic

Inside backend container:
- go get github.com/golang/mock/gomock
- Test Elasticsearch query building
- Test alert evaluation logic
- Test log transformation
- Test error handling and retries

Test Examples:
- TestBuildSearchQuery()
- TestEvaluateAlertCondition()
- TestParseLogEntry()
- TestElasticsearchRetry()

Run Tests:
docker-compose exec backend go test -v ./internal/services

Target Coverage: 75% of services/ package
```
**Files:**
- `backend_go/internal/services/services_test.go`

**Commit 33: Backend Integration Tests**
```
test(integration): add integration tests with test Elasticsearch

Inside backend container:
- Use testcontainers for Elasticsearch
- Test full log search flow
- Test alert creation and triggering
- Test end-to-end log ingestion

Test Setup:
- Start temporary Elasticsearch container
- Create test index
- Run tests
- Clean up containers

Test Examples:
- TestLogIngestionAndSearch()
- TestAlertTriggerAndNotification()
- TestLogTransformationPipeline()

Run Tests:
docker-compose exec backend go test -v -tags=integration ./...

Target Coverage: Integration paths of core flows
```
**Files:**
- `backend_go/internal/api/integration_test.go`

**Commit 34: Frontend Component Tests**
```
test(components): add React component tests with Vitest

Inside frontend container:
- npm install -D vitest @testing-library/react
- Create tests for StatsCard, LogViewer, AlertsPanel
- Test rendering with different props
- Test user interactions (click, input)
- Test error states

Test Examples:
- Test StatsCard renders value correctly
- Test LogViewer filters by service
- Test AlertsPanel form validation
- Test error state display

Run Tests:
docker-compose exec frontend npm run test

Target Coverage: 70% of components/
```
**Files:**
- `frontend/src/components/__tests__/StatsCard.test.tsx`
- `frontend/src/components/__tests__/LogViewer.test.tsx`
- `frontend/src/components/__tests__/AlertsPanel.test.tsx`

**Commit 35: Test Automation in CI/CD**
```
chore(ci): add GitHub Actions for automated testing

.github/workflows/test.yml:
- Run on every push and PR
- Lint Go code (golangci-lint)
- Lint TypeScript (ESLint)
- Run backend unit tests
- Run frontend unit tests
- Run backend integration tests (with containers)
- Generate coverage reports

Benefits:
- Catch bugs before merge
- Ensure code quality
- Coverage tracking
- Automated deployment gate

Running Locally:
# Simulate CI/CD locally
docker-compose exec backend golangci-lint run
docker-compose exec backend go test ./...
docker-compose exec frontend npm run lint
docker-compose exec frontend npm run test
```
**Files:**
- `.github/workflows/test.yml`
- `.golangci.yml` (Go linting config)
- `.eslintrc.json` (TypeScript linting config)

---

### PHASE 7: Configuration & Data Setup (Commits 36-39)
*Goal:* Configure Elasticsearch, Kafka, and seed data

**Commit 36: Elasticsearch Mapping & Setup**
```
chore(config): create Elasticsearch index mapping

Inside backend container (or via curl):
- Create config/elasticsearch-mapping.json
- Define index mapping for log documents
- Specify field types (text, keyword, date)
- Configure analyzers for full-text search
- Set up index lifecycle management

Mapping Example:
{
  "mappings": {
    "properties": {
      "timestamp": { "type": "date" },
      "service": { "type": "keyword" },
      "level": { "type": "keyword" },
      "message": { "type": "text" },
      "trace_id": { "type": "keyword" }
    }
  }
}

Create Index:
curl -X PUT http://localhost:9200/logs-2024.01 \
  -H "Content-Type: application/json" \
  -d @config/elasticsearch-mapping.json

Testing:
curl http://localhost:9200/logs-2024.01/_mapping
```
**Files:**
- `config/elasticsearch-mapping.json`

**Commit 37: Kafka Topic Configuration**
```
chore(config): create Kafka topics for log streaming

Inside kafka container:
- Create config/kafka-topics.yaml
- Define topics: logs, alerts, metrics
- Set replication factor (3 for HA)
- Set partition count (3-5 for performance)

Topics Created:
- logs: 3 partitions, replication 3
  (High-volume log stream)
- alerts: 1 partition, replication 3
  (Alert events, order matters)
- metrics: 2 partitions, replication 3
  (System metrics)

Create Topics Script:
#!/bin/bash
docker-compose exec kafka kafka-topics \
  --create \
  --topic logs \
  --partitions 3 \
  --replication-factor 3 \
  --bootstrap-server localhost:9092

Testing:
docker-compose exec kafka kafka-topics --list --bootstrap-server localhost:9092
```
**Files:**
- `config/kafka-topics.yaml`
- `scripts/create-kafka-topics.sh`

**Commit 38: Sample Data Generation**
```
scripts(data): create sample log generation for testing

Inside backend container:
- Create scripts/generate-sample-logs.sh
- Generate realistic log data
- Different service names
- Varying log levels (ERROR, WARN, INFO)
- Realistic timestamps
- Post to /api/logs/ingest

Sample Data Includes:
- 1000 logs across 5 services
- Various log levels with realistic distribution
- Time spread over 24 hours
- Different message patterns

Generate Test Data:
docker-compose exec backend ./scripts/generate-sample-logs.sh

Benefits:
- Test search functionality
- Test alert triggering
- Visualize dashboard with real data
- Performance testing baseline

Testing:
After running script, check:
- Elasticsearch: curl http://localhost:9200/logs-*/_count
- Dashboard: http://localhost:3000 should show stats
```
**Files:**
- `scripts/generate-sample-logs.sh` - Bash script to generate logs
- `scripts/generate-sample-logs-go.go` (alternative Go version)

**Commit 39: Log Collectors Setup (Fluentd/Logstash)**
```
chore(config): configure Fluentd and Logstash for log collection

Inside docker-compose:
- Add Fluentd service (log collector agent)
- Add Logstash service (alternative pipeline)
- Configure collectors to:
  - Listen on network sockets
  - Parse JSON logs
  - Forward to Kafka topic (logs)
  - Forward to Elasticsearch (optional)

Fluentd Config (config/fluentd.conf):
- Input plugin: TCP on port 24224
- Filter: Parse JSON, add metadata
- Output plugin: Kafka (topic: logs)

Logstash Config (config/logstash.conf):
- Input: TCP/HTTP
- Filter: Grok patterns, field parsing
- Output: Kafka and Elasticsearch

Use Cases:
- Collect logs from Docker containers
- Collect logs from external services
- Parse and normalize different log formats

Testing:
echo '{"msg":"test"}' | nc localhost 24224
Check Kafka: docker-compose exec kafka kafka-console-consumer --topic logs
```
**Files:**
- `config/fluentd.conf`
- `config/logstash.conf`
- Update `docker/docker-compose.yml` with Fluentd/Logstash services

---

### PHASE 8: Kubernetes Deployment (Commits 40-43)
*Goal:* Prepare for production Kubernetes deployment

**Commit 40: Backend Kubernetes Deployment**
```
chore(k8s): create backend deployment manifest

kubernetes/backend-deployment.yaml:
- 3 replicas for high availability
- Resource requests: CPU 500m, Memory 256Mi
- Resource limits: CPU 1000m, Memory 512Mi
- Liveness probe: /health endpoint
- Readiness probe: /health endpoint
- Service for internal networking
- ConfigMap for environment variables
- Health check strategy

Deployment Features:
- Rolling updates
- Pod disruption budgets
- Resource quotas
- Service discovery
- Auto-scaling ready

Testing locally:
kubectl apply -f kubernetes/backend-deployment.yaml
kubectl get pods
kubectl logs pod/backend-xxxxx
```
**Files:**
- `kubernetes/backend-deployment.yaml`
- `kubernetes/backend-service.yaml`

**Commit 41: Frontend Kubernetes Deployment**
```
chore(k8s): create frontend deployment manifest

kubernetes/frontend-deployment.yaml:
- 2 replicas for high availability
- Static file serving via Nginx
- Resource requests: CPU 250m, Memory 128Mi
- Resource limits: CPU 500m, Memory 256Mi
- Service with LoadBalancer for external access
- ConfigMap for Nginx configuration
- Readiness probe: HTTP GET /

Deployment Features:
- CDN-ready headers
- Gzip compression
- Cache control
- Service discovery for backend API

Testing locally:
kubectl apply -f kubernetes/frontend-deployment.yaml
kubectl get services
curl http://localhost:80 (after port-forward)
```
**Files:**
- `kubernetes/frontend-deployment.yaml`
- `kubernetes/frontend-service.yaml`
- `kubernetes/nginx-configmap.yaml`

**Commit 42: Data Services (Elasticsearch, Kafka) Deployments**
```
chore(k8s): create stateful deployments for data services

StatefulSets:
- elasticsearch-statefulset.yaml (3 nodes)
  - Persistent volumes for data
  - Headless service for inter-node communication
  - Resource limits: 2 CPUs, 2Gi memory

- kafka-statefulset.yaml (3 brokers)
  - Persistent volumes for logs
  - Zookeeper StatefulSet (3 replicas)
  - Headless service for broker discovery
  - Resource limits: 1 CPU, 1Gi memory

Features:
- Data persistence
- Ordered pod naming
- Stable network identities
- Orderly scaling

Testing locally:
kubectl apply -f kubernetes/elasticsearch-statefulset.yaml
kubectl apply -f kubernetes/kafka-statefulset.yaml
kubectl get statefulsets
kubectl get pvc (check persistent volumes)
```
**Files:**
- `kubernetes/elasticsearch-statefulset.yaml`
- `kubernetes/elasticsearch-service.yaml`
- `kubernetes/kafka-statefulset.yaml`
- `kubernetes/zookeeper-statefulset.yaml`

**Commit 43: ConfigMaps and Secrets**
```
chore(k8s): create configuration and secrets management

ConfigMaps for:
- Backend configuration (log level, timeouts)
- Frontend configuration (API base URL, features)
- Elasticsearch mapping and settings
- Kafka topic configurations

Secrets for:
- Database credentials
- API keys
- TLS certificates
- OAuth tokens

Usage in Pods:
env:
  - name: API_LOG_LEVEL
    valueFrom:
      configMapKeyRef:
        name: backend-config
        key: log-level

Testing:
kubectl create configmap backend-config --from-literal=log-level=debug
kubectl get configmaps
kubectl describe configmap backend-config
```
**Files:**
- `kubernetes/configmap-backend.yaml`
- `kubernetes/configmap-frontend.yaml`
- `kubernetes/secret-credentials.yaml` (template)

---

### PHASE 9: Documentation & Deployment Guides (Commits 44-47)
*Goal:* Complete documentation for Docker and production deployment

**Commit 44: Quickstart Guide (Docker)**
```
docs: add 5-minute Docker quickstart guide

QUICKSTART.md:
1. Install Docker and Docker Compose
2. Clone repository
3. Copy .env.example to .env
4. docker-compose up
5. Access services:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8000
   - Elasticsearch: http://localhost:9200
   - Kafka: localhost:9092

Troubleshooting:
- Port conflicts
- Slow startup
- Connection errors

Benefits of Docker approach:
- No local dependencies to install
- Reproducible environment
- Easy team onboarding
```
**Files:**
- `QUICKSTART.md` (or update existing)
- `backend_go/QUICKSTART.md`

**Commit 45: Development Workflow Guide**
```
docs: document complete Docker-based development workflow

DEVELOPMENT_WORKFLOW.md:
1. Start Services
   docker-compose up

2. Edit Code
   - Backend code: edit backend_go/
   - Frontend code: edit frontend/
   - Changes auto-reload in containers

3. View Logs
   docker-compose logs -f backend
   docker-compose logs -f frontend

4. Run Tests
   docker-compose exec backend go test ./...
   docker-compose exec frontend npm test

5. Access Services
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8000
   - Database CLI: docker-compose exec elasticsearch curl http://localhost:9200

6. Stop Services
   docker-compose down

Advanced Topics:
- Attaching debuggers
- Database inspection
- Performance profiling
- Multi-service debugging

Example Scenarios:
- Testing new log search feature
- Creating new alert rule
- Fixing integration issue
```
**Files:**
- `docs/DEVELOPMENT_WORKFLOW.md` (or DEVELOPMENT.md)
- Update `README.md` with Docker setup section

**Commit 46: Production Deployment Guide**
```
docs: create production deployment guides

DEPLOYMENT.md:
1. Kubernetes Deployment
   - kubectl apply -f kubernetes/
   - Scale replicas
   - Monitor pods

2. Cloud Platforms
   - AWS EKS setup
   - GCP GKE setup
   - Azure AKS setup

3. Database Setup
   - Elasticsearch cluster setup
   - Kafka production config
   - Persistence configuration

4. Security Setup
   - TLS certificates
   - Network policies
   - Secret management
   - RBAC configuration

5. Monitoring & Logging
   - Prometheus setup
   - Grafana dashboards
   - ELK stack integration
   - Alert rules

6. Scaling Strategies
   - Horizontal pod autoscaling
   - Elasticsearch cluster scaling
   - Kafka broker scaling

7. High Availability
   - Multi-AZ deployment
   - Backup strategies
   - Disaster recovery
   - Failover procedures
```
**Files:**
- `docs/DEPLOYMENT.md`
- `docs/PRODUCTION_CHECKLIST.md` (optional)

**Commit 47: API Documentation**
```
docs: complete API endpoint documentation

API_GUIDE.md:
- Authentication & authorization
- Request/response format
- All endpoints with examples
- Error codes and meanings
- Rate limiting details
- Best practices

Endpoints Documented:
- GET /health
- GET /metrics
- POST /api/logs/ingest
- POST /api/logs/search
- GET|POST|PUT|DELETE /api/alerts
- GET /api/alerts/history

Example Requests:
curl -X POST http://localhost:8000/api/logs/search \
  -H "Content-Type: application/json" \
  -d '{"query":"error","service":"backend","limit":50}'

Response Schemas:
{
  "logs": [
    {
      "timestamp": "2024-01-01T12:00:00Z",
      "service": "backend",
      "level": "ERROR",
      "message": "...",
      "trace_id": "..."
    }
  ],
  "total_hits": 1234,
  "search_time_ms": 45
}

Error Responses:
{
  "error": "INVALID_REQUEST",
  "message": "...",
  "request_id": "..."
}
```
**Files:**
- `docs/API_GUIDE.md` (update or create)

---

### PHASE 10: Final Optimization & Polish (Commits 48-50)
*Goal:* Performance, security, and production-readiness

**Commit 48: Performance Optimization**
```
perf: optimize search queries and indexing

Backend Optimizations:
- Elasticsearch query caching (in-memory)
- Bulk indexing for log ingestion
- Index lifecycle management (daily rollover)
- Connection pooling tuning
- Goroutine limits to prevent resource exhaustion

Frontend Optimizations:
- Lazy load React components
- Code splitting in Vite
- Memoization of expensive components
- Virtual scrolling for log lists
- Image optimization

Measurable Results:
- Search latency: <500ms (p95)
- Log ingestion: >10k logs/sec
- Frontend load time: <2s
- Bundle size: <200KB gzipped

Testing:
docker-compose exec backend go test -bench ./...
docker-compose exec frontend npm run build
docker-compose exec frontend npm run preview
```
**Files:**
- Updates to backend services
- Updates to frontend components
- Vite config optimization

**Commit 49: Security Hardening**
```
feat(security): add authentication, authorization, and security headers

Security Features:
- JWT-based authentication
- Role-based access control (admin, analyst, viewer)
- API rate limiting (100 req/min per IP)
- Input validation and sanitization
- Security headers (CORS, CSP, HSTS)
- TLS enforcement
- Sensitive data redaction
- Log encryption at rest
- Dependency vulnerability scanning

Implementation:
- JWT middleware for protected endpoints
- RBAC middleware
- Rate limiting middleware
- Input validators
- Security headers middleware

Testing:
- Test unauthorized access (should fail)
- Test rate limiting (exceed limit)
- Test XSS prevention
- Test SQL injection prevention (if applicable)

Running:
docker-compose exec backend ./security-tests.sh
```
**Files:**
- New middleware files for auth, rate limiting
- Updates to router configuration
- Security policy documentation

**Commit 50: Release v1.0.0 & Final Polish**
```
chore: prepare project for v1.0.0 release

Version Updates:
- go.mod: module version
- package.json: version 1.0.0
- VERSION file: 1.0.0
- Dockerfile tags: v1.0.0

Release Materials:
- CHANGELOG.md: all features and fixes
- RELEASE_NOTES.md: highlights
- Migration guide (if applicable)
- Breaking changes list (none for v1)

Final Checklist:
✓ All tests passing
✓ Linting clean
✓ Documentation complete
✓ Docker images built and tagged
✓ Kubernetes manifests tested
✓ Security scan clean
✓ Performance benchmarks baseline
✓ API documentation final

Release Steps:
1. Update version numbers
2. Create CHANGELOG
3. Commit with message "chore(release): v1.0.0"
4. Create Git tag: git tag -a v1.0.0
5. Build Docker images: docker build -t <image>:v1.0.0
6. Push images to registry
7. Create GitHub release
8. Announce availability

Testing Final Release:
docker pull <image>:v1.0.0
docker-compose -f docker-compose-prod.yml up
Verify all features working
```
**Files:**
- Update `go.mod` with version
- Update `package.json` with version
- Create `CHANGELOG.md`
- Create `RELEASE_NOTES.md`
- Tag release in Git

---
*Goal:* Establish clean, professional project structure

**Commit 1: Project Initialization**
```
chore: initialize distributed log monitoring project

- Create directory structure for backend, frontend, config, docs
- Add README with project vision and architecture
- Add .gitignore for Go, Node, Docker
- Create LICENSE file (MIT)
```
**Files:**
- `README.md` - Project overview
- `.gitignore` - Git ignore rules
- `LICENSE` - MIT license
- `ARCHITECTURE.md` - High-level design

**Commit 2: Backend Project Setup**
```
chore(backend): initialize Go module and project structure

- Create Go module: github.com/yourusername/distributed-log-monitoring
- Set up go.mod and go.sum
- Create cmd/backend/main.go entry point
- Create internal/ directory structure
- Add Makefile with build targets
- Create Dockerfile for containerization
```
**Files:**
- `backend_go/go.mod`
- `backend_go/go.sum`
- `backend_go/Makefile`
- `backend_go/Dockerfile`
- `backend_go/cmd/backend/main.go` (minimal main)
- `backend_go/internal/` (directory structure)

**Commit 3: Configuration Management**
```
feat(config): implement environment-based configuration system

- Create internal/config/config.go with structured Config struct
- Support environment variables with sensible defaults
- Load defaults for Elasticsearch, Kafka, API endpoints
- Add configuration validation
- Document all configuration options

Config supports:
- API_HOST, API_PORT (HTTP server)
- ELASTICSEARCH_HOST, ELASTICSEARCH_PORT (search backend)
- KAFKA_BROKERS (comma-separated)
- LOG_LEVEL, LOG_FORMAT (debug/info/warn/error)
```
**Files:**
- `backend_go/internal/config/config.go`
- `backend_go/.env.example`
- `backend_go/README.md` (config section)

**Commit 4: Logging Infrastructure**
```
feat(logging): implement structured logging with Zap

- Initialize Uber Zap logger with JSON/text format support
- Support configurable log levels
- Add context-aware logging helpers
- Implement graceful logger shutdown

Features:
- Structured fields for better log queries
- Separate stdout/stderr handling
- Configurable output format for development/production
```
**Files:**
- `backend_go/internal/logger/logger.go`

**Commit 5: Frontend Project Setup**
```
chore(frontend): initialize React + TypeScript project with Vite

- Create Vite + React + TypeScript scaffold
- Configure Tailwind CSS for styling
- Set up PostCSS configuration
- Create tsconfig.json with strict mode
- Add build and dev scripts
```
**Files:**
- `frontend/package.json`
- `frontend/tsconfig.json`
- `frontend/vite.config.ts`
- `frontend/postcss.config.js`
- `frontend/tailwind.config.js`
- `frontend/index.html`
- `frontend/src/main.tsx`
- `frontend/src/App.tsx`

---

### PHASE 2: Backend API Core (Commits 6-12)
*Goal:* Build production-ready REST API layer

**Commit 6: HTTP Router Setup**
```
feat(api): establish HTTP router with Gorilla Mux

- Create internal/api/router.go
- Configure Gorilla Mux for request routing
- Add middleware for CORS, request logging, recovery
- Implement health check endpoint (/health)
- Add metrics middleware for Prometheus

Features:
- Request ID injection for tracing
- Request/response timing
- Structured error responses
```
**Files:**
- `backend_go/internal/api/router.go`
- `backend_go/internal/api/helpers.go` (response formatting)
- `backend_go/internal/middleware/` (directory)

**Commit 7: Data Models & Types**
```
feat(models): define core data structures

- Create Log, LogEntry, Alert, AlertRule types
- Implement JSON marshaling/unmarshaling
- Add validation methods
- Document field mappings to Elasticsearch

Structures:
- Log: timestamp, service, level, message, trace_id, context
- Alert: name, condition, threshold, channels, enabled
- SearchQuery: service, level, time_range, text_filter
```
**Files:**
- `backend_go/internal/models/types.go`
- `backend_go/internal/models/validation.go`

**Commit 8: Elasticsearch Client Integration**
```
feat(services): implement Elasticsearch client wrapper

- Create internal/services/elasticsearch_service.go
- Connection pool management and health checking
- Query builder for log searches (with filters, sorting)
- Index management (create, delete, mapping)
- Error handling and retry logic

Features:
- Connection retry with exponential backoff
- Structured query DSL builder
- Bulk operations support for performance
```
**Files:**
- `backend_go/internal/services/elasticsearch_service.go`
- `backend_go/internal/services/clients.go` (client initialization)

**Commit 9: Kafka Consumer Setup**
```
feat(worker): implement Kafka consumer for log streaming

- Create internal/worker/kafka_consumer.go
- Consumer group management
- Partition assignment strategies
- Offset management (auto-commit)
- Error handling and poison pill detection

Features:
- Graceful shutdown with in-flight message handling
- Metrics for consumer lag
- Configurable batch processing
```
**Files:**
- `backend_go/internal/worker/kafka_consumer.go`
- `backend_go/internal/services/kafka_service.go`

**Commit 10: Log Search & Retrieval**
```
feat(api): implement log search endpoint

- POST /api/logs/search - full-text search with filters
- Pagination support (offset, limit)
- Field filtering (service, level, time range)
- Result sorting and highlighting
- Performance optimization with indexes

Request:
{
  "query": "error",
  "service": "auth-service",
  "level": "ERROR",
  "from": "2024-01-01T00:00:00Z",
  "to": "2024-01-02T00:00:00Z",
  "limit": 50,
  "offset": 0
}

Response: List with metadata, total_hits, search_time_ms
```
**Files:**
- `backend_go/internal/api/log_handler.go`
- `backend_go/internal/services/log_service.go`

**Commit 11: Log Ingestion Endpoint**
```
feat(api): add log ingestion endpoint for streaming data

- POST /api/logs/ingest - accept log entries
- Batch ingestion support (array or newline-delimited JSON)
- Schema validation
- Transform and enrich (add timestamp, host info)
- Store to Elasticsearch and queue to Kafka

Features:
- Accepts JSON and plaintext logs
- Automatic field parsing (JSON) or regex parsing
- Service auto-discovery from source IP
- Enrichment with timestamp, host, container ID
```
**Files:**
- `backend_go/internal/api/log_handler.go` (ingest handler)
- Updates to `internal/services/log_service.go`

**Commit 12: Alerting System - Core**
```
feat(alert): implement alert rule engine and notifications

- Create internal/services/alert_service.go
- Alert rule evaluation engine
- Threshold-based detection (error rate %, frequency)
- Alert state management (firing, resolved)
- Notification channel abstraction

Features:
- Rule DSL: {field} {operator} {threshold} {duration}
- Multiple channels: email, webhook, Slack, PagerDuty
- Alert deduplication (avoid spam)
- Alert history tracking
```
**Files:**
- `backend_go/internal/services/alert_service.go`
- `backend_go/internal/api/alert_handler.go`
- `backend_go/internal/models/types.go` (AlertRule struct)

---

### PHASE 3: API & Advanced Features (Commits 13-18)
*Goal:* Complete the API feature set

**Commit 13: Alert Management REST API**
```
feat(api): implement alert management endpoints

- GET /api/alerts - list all alerts
- GET /api/alerts/{id} - get specific alert
- POST /api/alerts - create new alert rule
- PUT /api/alerts/{id} - update rule
- DELETE /api/alerts/{id} - remove rule
- GET /api/alerts/history - alert event history

Features:
- Role-based access control (future)
- Audit logging for changes
- Alert trigger history with timestamps
```
**Files:**
- `backend_go/internal/api/alert_handler.go`

**Commit 14: System Health & Metrics**
```
feat(metrics): add Prometheus metrics and health endpoints

- Create internal/metrics/prometheus.go
- System health check endpoint (dependencies check)
- Prometheus metrics exposure (/metrics)
- Custom metrics:
  - Log ingestion rate (logs/sec)
  - Search latency (p50, p95, p99)
  - Alert trigger rate
  - Elasticsearch query time
  - Kafka consumer lag

Metrics tracked:
- http_requests_total (by endpoint, method, status)
- http_request_duration_seconds (by endpoint)
- elasticsearch_query_duration_seconds
- logs_ingested_total
- alerts_triggered_total
```
**Files:**
- `backend_go/internal/metrics/prometheus.go`
- `backend_go/internal/api/system_handler.go` (health/metrics endpoints)

**Commit 15: Request Validation & Error Handling**
```
refactor(api): standardize error handling and request validation

- Create centralized error types and response formatting
- Input validation middleware
- Structured error responses with error codes
- Request body size limits
- Rate limiting infrastructure

Error Response Format:
{
  "error": "INVALID_REQUEST",
  "message": "Required field 'service' missing",
  "request_id": "req-12345",
  "timestamp": "2024-01-01T00:00:00Z"
}
```
**Files:**
- `backend_go/internal/api/helpers.go` (error handling)
- `backend_go/internal/middleware/` (validation middleware)

**Commit 16: Graceful Shutdown & Signal Handling**
```
feat(backend): implement graceful shutdown with signal handling

- Capture SIGTERM, SIGINT signals
- Close database connections cleanly
- Drain Kafka consumer gracefully
- Flush Prometheus metrics
- Wait for in-flight requests to complete (timeout)

Features:
- 30-second grace period for in-flight requests
- Prevent new connections after signal
- Log shutdown process
- Exit code 0 for clean shutdown
```
**Files:**
- `backend_go/cmd/backend/main.go` (signal handling)

**Commit 17: Application Logging Integration**
```
feat(backend): integrate structured logging throughout application

- Add logging to all major functions
- Log API requests and responses
- Log Elasticsearch queries and response times
- Log Kafka consumer errors and recovery
- Structured fields for context (trace_id, request_id, service)

Logging Levels:
- DEBUG: Elasticsearch queries, consumer offsets, configuration
- INFO: API requests, alerts triggered, connections established
- WARN: Slow queries, consumer lag, retries
- ERROR: Request failures, connection errors, data inconsistencies
```
**Files:**
- Updates throughout `internal/` (all handlers and services)

**Commit 18: Comprehensive API Documentation**
```
docs(api): add complete API documentation

- Document all endpoints with examples
- Request/response schemas
- Error codes and meanings
- Query parameter documentation
- Rate limits and best practices
- cURL examples

Endpoints documented:
- GET /health - Health check
- GET /metrics - Prometheus metrics
- POST /api/logs/ingest - Log ingestion
- POST /api/logs/search - Log search
- GET|POST|PUT|DELETE /api/alerts - Alert management
- GET /api/alerts/history - Alert history
```
**Files:**
- `docs/API_GUIDE.md`
- Update `backend_go/README.md`

---

### PHASE 4: Frontend Dashboard (Commits 19-26)
*Goal:* Build interactive React UI

**Commit 19: Layout & Styling Setup**
```
feat(frontend): establish base layout and styling system

- Create App.tsx with main layout structure
- Header, sidebar, main content area
- Dark mode support with CSS variables
- Tailwind CSS configuration review
- Global styles (globals.css)

Components:
- Navigation header with branding
- Sidebar with navigation links
- Content area wrapper
- Responsive grid system
```
**Files:**
- `frontend/src/App.tsx`
- `frontend/src/globals.css`
- `frontend/src/index.css`

**Commit 20: Dashboard Page Structure**
```
feat(dashboard): create main dashboard page

- Create pages/Dashboard.tsx
- Grid layout for dashboard cards
- Real-time stats section
- Recent alerts section
- Log stream section

Sections:
- Stats Cards (logs/hour, alerts/hour, services, errors)
- Alert Widget (recent triggered alerts)
- Log Stream (latest 10 logs)
- Search Panel (quick access)
```
**Files:**
- `frontend/src/pages/Dashboard.tsx`
- `frontend/src/components/StatsCard.tsx`

**Commit 21: Stats Card Component**
```
feat(components): implement reusable stats card component

- Display metric with title and value
- Optional trend indicator (up/down)
- Loading state
- Error handling
- Responsive sizing

Props:
- title: string
- value: number
- unit?: string
- trend?: number
- loading?: boolean
- error?: string
```
**Files:**
- `frontend/src/components/StatsCard.tsx`

**Commit 22: Log Viewer Component**
```
feat(components): build log viewer with filtering and search

- Display log entries in table format
- Columns: timestamp, service, level, message
- Color-coded log levels (ERROR=red, WARN=yellow, INFO=blue)
- Expandable row for full message
- Copy to clipboard functionality
- Pagination controls

Features:
- Service filter dropdown
- Log level checkboxes (ERROR, WARN, INFO, DEBUG)
- Time range picker
- Search/filter input
```
**Files:**
- `frontend/src/components/LogViewer.tsx`

**Commit 23: Alerts Panel Component**
```
feat(components): create alerts management panel

- Display alert rules in table
- Create new alert form
- Edit/delete functionality
- Enable/disable toggle
- Alert history view

Alert Rule Form:
- Name
- Service selector
- Condition (error rate, frequency, etc.)
- Threshold value
- Duration (e.g., errors in last 5min)
- Notification channels
- Enable/disable
```
**Files:**
- `frontend/src/components/AlertsPanel.tsx`

**Commit 24: State Management with Custom Hook**
```
feat(hooks): implement state management with useStore hook

- Create custom useStore hook for global state
- Manage active filters (service, level, time range)
- Manage dashboard stats
- Manage alert rules
- Manage UI state (theme, sidebar collapse)

State Structure:
{
  filters: { service, level, dateRange },
  stats: { logsPerHour, alertsPerHour, errorCount },
  alerts: AlertRule[],
  theme: 'light' | 'dark',
  loading: { [key]: boolean }
}
```
**Files:**
- `frontend/src/hooks/useStore.ts`

**Commit 25: API Service Layer**
```
feat(services): create API client with request handling

- Create services/api.ts with HTTP client setup
- Implement error handling and retry logic
- Add request/response interceptors
- Type-safe API endpoints
- Local storage caching for offline support

Endpoints:
- fetchLogs(filters, pagination)
- searchLogs(query, filters)
- createAlert(rule)
- updateAlert(id, rule)
- deleteAlert(id)
- fetchAlerts()
- getMetrics()
```
**Files:**
- `frontend/src/services/api.ts`
- Create `types/` directory for API response types

**Commit 26: Integration & Real-time Updates**
```
feat(frontend): integrate with backend API and add auto-refresh

- Connect Dashboard to API endpoints
- Fetch stats on component mount
- Auto-refresh stats every 30 seconds
- Connect LogViewer to search API
- Connect AlertsPanel to alert management API
- Error handling and user feedback
- Loading states for all async operations

Features:
- Graceful fallbacks for API errors
- Toast notifications for user feedback
- Loading skeletons during data fetch
- Auto-refresh with visual indication
```
**Files:**
- Updates to `frontend/src/pages/Dashboard.tsx`
- Updates to `frontend/src/components/LogViewer.tsx`
- Updates to `frontend/src/components/AlertsPanel.tsx`

---

### PHASE 5: Infrastructure & DevOps (Commits 27-33)
*Goal:* Containerization and deployment setup

**Commit 27: Docker Compose for Local Development**
```
chore(docker): create Docker Compose for full local development stack

Services:
- Backend API (Go) - port 8000
- Frontend (Nginx) - port 3000
- Elasticsearch - port 9200
- Kafka - port 9092
- Zookeeper - port 2181
- Postgres (for metadata) - port 5432

Features:
- Volume mounts for code hot-reloading
- Network isolation
- Health checks for all services
- Seed data initialization
```
**Files:**
- `docker/docker-compose.yml`
- `docker/Dockerfile.backend`
- `docker/Dockerfile.frontend`

**Commit 28: Dockerfile Optimization**
```
chore(docker): optimize Dockerfiles with multi-stage builds

Backend:
- Stage 1: Build stage (golang:1.21-alpine)
- Stage 2: Runtime stage (alpine)
- CGO disabled for static builds
- Minimal final image (~20MB)

Frontend:
- Stage 1: Build stage (node:20)
- Stage 2: Runtime stage (nginx:alpine)
- Nginx config with proper routing
- Final image (~50MB)
```
**Files:**
- `docker/Dockerfile.backend` (optimization)
- `docker/Dockerfile.frontend` (optimization)

**Commit 29: Kubernetes Manifests**
```
chore(k8s): create Kubernetes deployment manifests

Deployments:
- backend-deployment.yaml (3 replicas, resource limits)
- frontend-deployment.yaml (2 replicas, CDN ready)
- elasticsearch-statefulset.yaml (3 nodes, PVC)
- kafka-statefulset.yaml (3 brokers, ZK)

Features:
- Service definitions for networking
- ConfigMaps for configuration
- Secrets for sensitive data
- Resource requests and limits
- Health check probes
```
**Files:**
- `kubernetes/backend-deployment.yaml`
- `kubernetes/frontend-deployment.yaml`
- `kubernetes/elasticsearch-statefulset.yaml`
- `kubernetes/kafka-statefulset.yaml`

**Commit 30: Configuration Management**
```
chore(config): create environment configurations

Files:
- config/elasticsearch-mapping.json - Index mapping for logs
- config/fluentd.conf - Fluentd agent config
- config/logstash.conf - Logstash agent config
- config/kafka-topics.yaml - Kafka topic definitions

Topics to create:
- logs (3 partitions, replication factor 3)
- alerts (1 partition, replication factor 3)
- metrics (2 partitions, replication factor 3)
```
**Files:**
- `config/elasticsearch-mapping.json`
- `config/fluentd.conf`
- `config/logstash.conf`
- `config/kafka-topics.yaml`

**Commit 31: Build & Deployment Scripts**
```
chore(scripts): add build and deployment automation

Scripts:
- scripts/setup.sh - Initialize development environment
- scripts/create-kafka-topics.sh - Create Kafka topics
- scripts/generate-sample-logs.sh - Generate test data

Features:
- Dependency checking
- Database migrations
- Elasticsearch index creation
- Data seeding
```
**Files:**
- `scripts/setup.sh`
- `scripts/create-kafka-topics.sh`
- `scripts/generate-sample-logs.sh`

**Commit 32: Documentation & Quickstart**
```
docs: add comprehensive deployment and quickstart guides

Documents:
- QUICKSTART.md - 5-minute setup guide
- DEPLOYMENT.md - Production deployment steps
- INFRASTRUCTURE_UPDATES.md - Infrastructure changelog

Coverage:
- Local development setup
- Docker Compose quick start
- Kubernetes deployment
- AWS/GCP deployment examples
- Troubleshooting guide
```
**Files:**
- `docs/DEPLOYMENT.md`
- `backend_go/QUICKSTART.md`
- `INFRASTRUCTURE_UPDATES.md`

**Commit 33: CI/CD Pipeline**
```
chore(ci): add GitHub Actions CI/CD pipeline

Workflows:
- Build and test (on PR)
- Test coverage reporting
- Docker image build (on merge to main)
- Push to registry
- Automated deployment (optional)

Features:
- Linting (Go: golangci-lint, TS: ESLint)
- Unit tests (Go & TS)
- Integration tests
- Security scanning
- SonarQube analysis
```
**Files:**
- `.github/workflows/test.yml`
- `.github/workflows/build.yml`
- `.github/workflows/deploy.yml`

---

### PHASE 6: Testing & Quality (Commits 34-38)
*Goal:* Add comprehensive testing

**Commit 34: Backend Unit Tests - API Layer**
```
test(api): add unit tests for HTTP handlers

Tests:
- Test log search with various filters
- Test invalid input handling
- Test response format validation
- Test error responses
- Test pagination

Coverage target: 80% of api/ package
```
**Files:**
- `backend_go/internal/api/handlers_test.go`

**Commit 35: Backend Unit Tests - Services**
```
test(services): add unit tests for business logic

Tests:
- Elasticsearch query building
- Alert condition evaluation
- Log parsing and transformation
- Error handling and retries
- Kafka message handling

Coverage target: 75% of services/ package
```
**Files:**
- `backend_go/internal/services/services_test.go`

**Commit 36: Backend Integration Tests**
```
test(integration): add integration tests with test containers

Tests:
- Full log search flow (write + search)
- Alert creation and triggering
- Kafka consumer and processing
- Elasticsearch index creation and updates

Setup:
- Testcontainers for Elasticsearch, Kafka
- Test database for metadata
- Cleanup after each test
```
**Files:**
- `backend_go/internal/api/integration_test.go`

**Commit 37: Frontend Component Tests**
```
test(frontend): add React component tests with Vitest

Tests:
- StatsCard rendering with different props
- LogViewer filtering and pagination
- AlertsPanel form validation
- Error state handling
- Loading states

Coverage target: 70% of components/
```
**Files:**
- `frontend/src/components/__tests__/LogViewer.test.tsx`
- `frontend/src/components/__tests__/AlertsPanel.test.tsx`
- `frontend/src/components/__tests__/StatsCard.test.tsx`

**Commit 38: Code Quality & Performance**
```
chore(quality): add linting, formatting, and performance checks

Tools:
- Go: golangci-lint, gofmt
- TypeScript: ESLint, Prettier
- Security scanning: trivy (Docker images)
- Performance: benchmarks for critical paths
- Code coverage: >75% target

Automated:
- Pre-commit hooks
- GitHub Actions checks
- Coverage badges in README
```
**Files:**
- `.golangci.yml` (Go linting config)
- `.eslintrc.json` (TS linting config)
- `.prettierrc` (formatting)
- Update CI/CD pipeline

---

### PHASE 7: Performance & Optimization (Commits 39-42)
*Goal:* Production-ready performance

**Commit 39: Database Query Optimization**
```
perf(elasticsearch): optimize search queries and indexing

Optimizations:
- Query result caching (Redis)
- Bulk indexing for ingestion
- Index lifecycle management (daily rollover)
- Field analyzer optimization
- Aggregation pre-computation

Measurable improvements:
- Search latency: <500ms (p95) for 100M documents
- Ingestion throughput: >10k logs/sec
- Index size reduction: 30%
```
**Files:**
- Updates to `internal/services/elasticsearch_service.go`
- New caching layer documentation

**Commit 40: Connection Pooling & Concurrency**
```
perf(backend): implement connection pooling and concurrency limits

Features:
- Elasticsearch connection pool (size: 10-50)
- Kafka producer batching (batch size: 100)
- Goroutine limits to prevent resource exhaustion
- Buffered channels for async processing
- Circuit breaker for downstream services

Metrics:
- Connection pool utilization
- Goroutine count monitoring
- Circuit breaker state changes
```
**Files:**
- Updates to `internal/services/clients.go`
- Updates to `internal/worker/kafka_consumer.go`

**Commit 41: Frontend Performance**
```
perf(frontend): optimize React rendering and bundle size

Optimizations:
- Code splitting and lazy loading
- Memoization of expensive components
- Virtualized list for large log datasets
- Image optimization and lazy loading
- Reduce bundle size to <200KB gzipped

Results:
- Initial load time: <2s
- Lighthouse score: >90
```
**Files:**
- Updates to `frontend/src/components/LogViewer.tsx`
- Updates to `frontend/vite.config.ts`
- Create `frontend/src/components/VirtualizedLogList.tsx`

**Commit 42: Observability & Tracing**
```
feat(observability): add distributed tracing and enhanced monitoring

Features:
- Jaeger distributed tracing
- Trace context propagation (W3C TraceContext)
- Request tracing across backend/frontend
- Performance monitoring
- Custom metrics for business logic

Setup:
- Jaeger local deployment (docker-compose)
- Instrumentation of critical paths
- Trace sampling (10% in prod, 100% in dev)
```
**Files:**
- New tracing middleware
- Updates to key services
- `docker-compose.yml` (Jaeger service)

---

### PHASE 8: Security & Hardening (Commits 43-46)
*Goal:* Production security

**Commit 43: Authentication & Authorization**
```
feat(auth): add JWT-based authentication and RBAC

Features:
- JWT token generation and validation
- Role-based access control (admin, analyst, viewer)
- API key support for service-to-service auth
- Token refresh mechanism
- Logout functionality

Endpoints protected:
- Alert creation/modification/deletion (admin)
- Configuration access (admin)
- Log search (analyst, viewer)
```
**Files:**
- New `internal/auth/jwt.go`
- New `internal/middleware/auth.go`
- Updates to route registration

**Commit 44: API Security**
```
feat(api): implement security headers and rate limiting

Features:
- HTTPS enforcement (Strict-Transport-Security)
- CORS configuration
- Rate limiting (100 req/min per IP)
- Request validation and sanitization
- CSRF token protection (for forms)
- Input validation to prevent injection

Headers:
- Content-Security-Policy
- X-Frame-Options
- X-Content-Type-Options
- Strict-Transport-Security
```
**Files:**
- Updates to `internal/api/router.go`
- New `internal/middleware/security.go`

**Commit 45: Data Privacy & Encryption**
```
feat(security): add encryption and data privacy measures

Features:
- Sensitive data masking (PII redaction)
- Encryption at rest for credentials
- TLS for all connections
- Data retention policies
- GDPR/compliance support (data export)

Redaction patterns:
- Email addresses
- Credit card numbers
- API keys and tokens
- Phone numbers
```
**Files:**
- New `internal/security/encryption.go`
- New `internal/security/redaction.go`
- Configuration for policies

**Commit 46: Security Testing**
```
test(security): add security testing and vulnerability checks

Tests:
- SQL injection prevention (N/A for ES, but for DB)
- XSS prevention in frontend
- CSRF token validation
- Authentication bypass attempts
- Permission escalation tests

Tools:
- OWASP ZAP scanning
- Trivy container scanning
- Snyk dependency vulnerability scanning
```
**Files:**
- Security test suite
- GitHub Actions security workflow

---

### PHASE 9: Documentation & Community (Commits 47-50)
*Goal:* Maintainable, documented project

**Commit 47: Architecture Deep Dive**
```
docs: create comprehensive architecture documentation

Sections:
- System design decisions and rationale
- Component interaction diagrams
- Data flow documentation
- Scaling strategies
- High availability setup
- Disaster recovery procedures

Diagrams:
- C4 architecture diagrams
- Sequence diagrams for key flows
- Deployment topologies
```
**Files:**
- Enhanced `docs/ARCHITECTURE.md`
- `docs/DESIGN_DECISIONS.md`
- Add architecture diagrams (PNG/SVG)

**Commit 48: Developer Guide**
```
docs: add complete developer guide for contributors

Sections:
- Development environment setup
- Code structure explanation
- How to add new features
- How to write tests
- Debugging guide
- Contributing guidelines

Topics:
- Git workflow
- Code review process
- Release process
- Dependency management
```
**Files:**
- `CONTRIBUTING.md`
- `DEVELOPER_GUIDE.md`
- Enhanced `backend_go/IMPLEMENTATION_SUMMARY.md`

**Commit 49: Troubleshooting & FAQ**
```
docs: create troubleshooting guide and FAQ

Common issues:
- Service won't start
- Elasticsearch connection errors
- Kafka connection issues
- High memory usage
- Slow queries
- Docker networking problems

Debug steps:
- How to check logs
- How to verify connections
- How to reset state
- Performance profiling
```
**Files:**
- `docs/TROUBLESHOOTING.md`
- `docs/FAQ.md`

**Commit 50: Final Polish & Versioning**
```
chore: prepare project for release v1.0.0

Changes:
- Version bump in all relevant files (go.mod, package.json)
- Create CHANGELOG.md with all features
- Add version endpoint to API
- Create release notes
- Tag release in Git
- Update documentation with installation instructions

Release contents:
- Docker images tagged with version
- Release notes on GitHub
- Deployment guides
```
**Files:**
- Update `go.mod`, `package.json` with version
- Create `CHANGELOG.md`
- Create release tag and notes

---

---

## 📈 Commit Statistics

- **Total Commits:** 50
- **By Category:**
  - Docker Infrastructure: 4 commits
  - Backend API (in containers): 14 commits
  - Frontend (in containers): 8 commits
  - Integration & Real-time: 5 commits
  - Testing: 5 commits
  - Configuration & Data: 4 commits
  - Kubernetes: 4 commits
  - Documentation: 4 commits
  - Optimization & Security: 3 commits
  - Release & Polish: 1 commit

- **Expected Code Distribution:**
  - Backend Go: ~3500 LOC
  - Frontend React/TS: ~1200 LOC
  - Configuration: ~800 LOC
  - Tests: ~1500 LOC
  - **Total:** ~7000 LOC

---

## 🚀 How to Execute This Plan (Docker-First Approach)

### Prerequisites (All you need!)
```bash
# Install Docker and Docker Compose
# Download from: https://www.docker.com/products/docker-desktop

# Verify installation
docker --version       # Docker 24.0+
docker-compose --version  # Docker Compose 2.20+
```

### Complete Development Workflow

```bash
# 1. Clone or initialize the repository
git init distributed_log_monitoring
cd distributed_log_monitoring

# 2. For each commit phase:

# Phase 1: Set up Docker infrastructure
git add .
git commit -m "chore: initialize distributed log monitoring project with Docker-first approach"
git commit -m "chore(docker): set up Docker Compose with all services"
git commit -m "chore(docker): create optimized Go development container"
git commit -m "chore(docker): create Node development container with Vite"

# 3. Start development environment
docker-compose up
# This starts ALL services with mounted volumes for live development

# 4. In another terminal, develop the backend
# Edit files in backend_go/ → Changes auto-reload in container
docker-compose exec backend bash  # If you need shell access
docker-compose logs -f backend    # View backend logs

# 5. In another terminal, develop the frontend
# Edit files in frontend/ → Changes auto-reload in browser
docker-compose logs -f frontend   # View frontend logs

# 6. Continue with remaining commits (5-50)
# For each feature, edit code, test in containers, commit

# 7. Run tests inside containers
docker-compose exec backend go test ./...
docker-compose exec frontend npm test

# 8. When done, stop services
docker-compose down
```

### Local Development Setup (No Installation Needed!)

```bash
# Everything you edit is inside containers
# Your workflow:
1. Edit backend_go/... (Go code)
   → Auto-compiles in backend container via air
   → Hot-reload takes ~2-5 seconds

2. Edit frontend/src/... (React/TypeScript)
   → Vite dev server auto-updates in browser
   → Changes appear instantly (HMR)

3. Check API responses
   curl http://localhost:8000/health
   curl http://localhost:8000/api/logs/search

4. View frontend
   Browser: http://localhost:3000

5. Inspect services
   Elasticsearch: http://localhost:9200/_health
   Kafka: docker-compose exec kafka kafka-topics --list
```

### Development Environment Structure

```
distributed_log_monitoring/
├── docker/
│   ├── docker-compose.yml          # All services, volumes, networking
│   ├── Dockerfile.backend.dev      # Go development image (air, hot reload)
│   └── Dockerfile.frontend.dev     # Node development image (Vite)
├── backend_go/
│   ├── cmd/backend/main.go         # Entry point (edit and auto-reloads)
│   ├── internal/                   # Packages (edit and auto-reloads)
│   └── .air.toml                   # Air config for hot reload
├── frontend/
│   ├── src/                        # React components (edit and auto-reloads)
│   ├── package.json                # Node dependencies
│   └── vite.config.ts              # Vite dev server config
├── config/                         # Elasticsearch, Kafka, Fluentd configs
├── kubernetes/                     # K8s manifests (for later production)
├── scripts/                        # Helper scripts
└── docs/                           # Documentation
```

### What Happens When You Run docker-compose up

```
✓ Backend container starts
  - Builds Go binary (first time: ~30s)
  - air watches go files
  - Listening on :8000
  
✓ Frontend container starts
  - Installs npm dependencies
  - Starts Vite dev server
  - Listening on :3000 with HMR enabled
  
✓ Elasticsearch starts
  - Initializes on port :9200
  - Ready for indexing
  
✓ Kafka + Zookeeper start
  - Ready for log streaming
  - Topics auto-created on first use
  
✓ All services networked together
  - Backend can reach Elasticsearch via 'elasticsearch:9200'
  - Frontend can reach Backend via 'http://backend:8000'
  - All on internal Docker network
```

### Commit at Each Phase

After completing each phase:
```bash
git add .
git commit -m "feat(scope): meaningful commit message

- What you implemented
- How it works
- Why this approach

Developed and tested in Docker containers."
```

---

## 💡 Benefits of This Docker-First Approach

### For Development
✅ **No installation hell** - Just Docker  
✅ **Consistent environment** - Same as production  
✅ **Hot reload** - See changes instantly  
✅ **Easy onboarding** - New developers: `docker-compose up`  
✅ **Team alignment** - Everyone uses same containers  
✅ **Reproducible bugs** - Environment is identical  

### For Deployment
✅ **Development → Production** - Minimal config changes  
✅ **CI/CD ready** - Same containers in pipeline  
✅ **Kubernetes ready** - Manifests already prepared  
✅ **Scaling simple** - Just orchestrate containers  
✅ **Version control** - Exact versions in docker-compose  

### For Learning
✅ **See full system** - All services running together  
✅ **Understand dependencies** - Services interact in real-time  
✅ **DevOps maturity** - Learn containerization from day 1  
✅ **Production-like** - Not "works on my machine"  

---

## 🛠️ Common Development Tasks in Docker

### Add a New Go Dependency
```bash
docker-compose exec backend go get github.com/user/package
# Your go.mod/go.sum automatically updated
# Restart backend (or air auto-detects)
```

### Add a New Node Package
```bash
docker-compose exec frontend npm install package-name
# Your package.json/package-lock.json automatically updated
# Vite auto-reloads
```

### Run Backend Tests
```bash
docker-compose exec backend go test ./... -v
docker-compose exec backend go test -cover ./...
```

### Run Frontend Tests
```bash
docker-compose exec frontend npm test
docker-compose exec frontend npm run test:coverage
```

### Debug Backend
```bash
# View logs
docker-compose logs -f backend

# View specific service
docker-compose logs backend | tail -50

# Real-time logs with timestamps
docker-compose logs -f --timestamps backend
```

### Inspect Services
```bash
# Connect to Elasticsearch
curl http://localhost:9200/_cat/indices

# Check Kafka topics
docker-compose exec kafka kafka-topics --list --bootstrap-server kafka:9092

# View running containers
docker-compose ps

# Check service health
docker-compose exec backend curl http://localhost:8000/health
```

### Reset Everything
```bash
# Stop all services
docker-compose down

# Remove volumes (fresh data)
docker-compose down -v

# Rebuild containers (fresh images)
docker-compose build --no-cache

# Start fresh
docker-compose up
```

---

## ✅ Development Checklist for Each Commit

Before committing, ensure:

- [ ] Code works in the container
- [ ] Tests pass: `docker-compose exec backend go test ./...`
- [ ] No linting errors: `docker-compose exec backend golangci-lint run`
- [ ] Frontend builds: `docker-compose exec frontend npm run build`
- [ ] API tested: `curl http://localhost:8000/health`
- [ ] Logs are clear (check `docker-compose logs`)
- [ ] Git status clean: `git status`
- [ ] Commit message is descriptive
- [ ] Feature/fix is self-contained (atomic)

---

## 🎯 Time Estimates (Working in Containers)

**Per Phase Development Time:**
- Phase 1 (Docker setup): 1-2 hours (first time understanding)
- Phase 2 (Backend API): 4-6 hours (features + testing)
- Phase 3 (Metrics): 2-3 hours
- Phase 4 (Frontend): 4-6 hours (components + styling)
- Phase 5 (Integration): 2-3 hours
- Phase 6 (Testing): 3-4 hours
- Phase 7 (Config): 1-2 hours
- Phase 8 (Kubernetes): 3-4 hours (understanding manifests)
- Phase 9 (Docs): 2-3 hours
- Phase 10 (Optimization): 2-3 hours

**Total:** 24-36 hours (part-time development)

---

## 📚 Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [React Best Practices](https://react.dev/reference/rules)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [12 Factor App](https://12factor.net/)

---

## 🎉 Result

After completing all 50 commits:
- ✨ Fully containerized development environment
- 🐳 Docker-based entire system
- 🚀 Production-ready Kubernetes manifests
- 📖 Comprehensive documentation
- 🧪 Well-tested codebase
- 📊 Professional commit history showing Docker-first approach
- 🔒 Security-hardened system
- ⚡ Performance-optimized code

**You'll have a portfolio project that demonstrates:**
1. Docker and containerization expertise
2. Full-stack development (Go + React)
3. DevOps and infrastructure thinking
4. Professional development practices
5. Scalable system design
6. Container orchestration knowledge

---

**Total Development Timeline:** 6-9 weeks (part-time, learning) or 3-4 weeks (full-time, experienced)

**Good luck! Your Docker-first development approach will impress! 🚀**
