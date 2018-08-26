FROM postgres:10.3

COPY create_table.sql /docker-entrypoint-initdb.d/1.sql

USER postgres

EXPOSE 5432

CMD  ["postgres"]
