#!/bin/sh

echo ""
echo "Running migrations..."
/goose -dir=/app/migrations/ postgres "${POSTGRES_URL}" up
echo "Finished running migrations"
echo ""

exec "$@"
