# syntax=docker/dockerfile:1
FROM postgres

ENV POSTGRES_PASSWORD abc123
ENV POSTGRES_DB shakespeare_plays

COPY init.sql /docker-entrypoint-initdb.d/