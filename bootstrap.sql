create database `exchange` character set utf8mb4 collate utf8mb4_unicode_ci;

use `exchange`;

create table `currency`(
  `cur_id` int not null primary key,
  `date` datetime not null,
  `cur_abbreviation` varchar(3) not null,
  `cur_scale` int not null,
  `cur_name` varchar(256) not null,
  `cur_officialrate` double not null
);

create or replace user 'exchange'@'127.0.0.1' identified by 'exchange';

grant select, insert on `exchange`.* to 'exchange'@'127.0.0.1';

flush privileges;