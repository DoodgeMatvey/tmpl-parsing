CREATE DATABASE parsetmpl_db
USE parsetmpl_db
CREATE TABLE template (
	service_name       varchar(30) not null,
    feature_name       varchar(30) not null,
    feature_id         varchar(30) not null,
    feature_descr      varchar(60) not null,
    endpoints_path     varchar(60) not null,
    endpoints_methods  varchar(60) not null
)