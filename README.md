This project demonstrates full integration between a Go web server and the ELK stack (Elasticsearch, Logstash, Kibana) using 
structured JSON logs over UDP. 
It's designed for learning and experimentation.

---

## 📦 Stack Overview

| Component     | Purpose                                |
|---------------|----------------------------------------|
| **Go App**    | HTTP server that logs structured events |
| **Logstash**  | Parses and forwards logs to Elasticsearch |
| **Elasticsearch** | Stores and indexes log data            |
| **Kibana**    | Visualizes log data through dashboards   |

---

## 🚀 Quick Start

### 1. Clone & Run

```bash
git clone https://github.com/ervand7/go-elk-kibana-lab.git
cd go-elk-kibana-lab
./script.sh  # or: docker compose up --build
