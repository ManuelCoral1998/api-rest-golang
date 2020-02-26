CREATE DATABASE apirest;

CREATE USER IF NOT EXISTS apirestuser;

CREATE TABLE apirest.domainservers (domain STRING, servers STRING, sslgrade STRING, timeAccess STRING, PRIMARY KEY (domain, servers));

GRANT ALL ON DATABASE apirest TO apirestuser;
