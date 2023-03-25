create database `sample` default character set utf8mb4;

use `sample`;

create table `users` (
  id int auto_increment primary key,
  username varchar(255) not null,
  password varchar(255) not null,
  age int,
  created_at datetime not null,
  updated_at datetime not null
);

