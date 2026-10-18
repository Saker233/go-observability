FROM grafana/grafana:latest

COPY grafana/dashboards.yml \
    /etc/grafana/provisioning/dashboards/dashboards.yml

COPY grafana/go-observability.json \
    /var/lib/grafana/dashboards/go-observability.json