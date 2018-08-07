FROM Postgres:10.3

COPY create_table.sql /docker-enterpoint-initdb.d/

CMD postgres
