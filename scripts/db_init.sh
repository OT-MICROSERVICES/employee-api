#!/usr/bin/env bash

SCYLLADB_HOST="${SCYLLADB_HOST:=localhost}"
SCYLLADB_USERNAME="${SCYLLADB_USERNAME:=cassandra}"
SCYLLADB_PASSWORD="${SCYLLADB_PASSWORD:=cassandra}"
REPLICATION_FACTOR="${DATABASE_REPLICATION_FACTOR:=1}"

cqlsh "${SCYLLADB_HOST}" -u "${SCYLLADB_USERNAME}" -p "${SCYLLADB_PASSWORD}" \
-e "CREATE KEYSPACE employee_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : ${REPLICATION_FACTOR}};"
