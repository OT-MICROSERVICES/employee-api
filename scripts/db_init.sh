#!/usr/bin/env bash

SCYLLADB_HOST="${SCYLLADB_HOST:=localhost}"
SCYLLADB_USERNAME="${SCYLLADB_USERNAME:=cassandra}"
SCYLLADB_PASSWORD="${SCYLLADB_PASSWORD:=cassandra}"
REPLICATION_FACTOR="${DATABASE_REPLICATION_FACTOR:=1}"

cqlsh "${SCYLLADB_HOST}" -u "${SCYLLADB_USERNAME}" -p "${SCYLLADB_PASSWORD}" \
-e "CREATE KEYSPACE IF NOT EXISTS employee_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : ${REPLICATION_FACTOR}};"

cqlsh "${SCYLLADB_HOST}" -u "${SCYLLADB_USERNAME}" -p "${SCYLLADB_PASSWORD}" \
-e "CREATE TABLE IF NOT EXISTS employee_db.employee_info (id text, name text, designation text, department text, joining_date date, address text, office_location text, status text, email text, annual_package float, phone_number text, PRIMARY KEY (id, joining_date)) WITH CLUSTERING ORDER BY (joining_date DESC);"
