--删除数据库
drop database if exists littleblog;
-- 建立数据库
create database littleblog character set utf8;

-- 切换数据库

use littleblog

-- 建立用户表

drop table if exists t_user;
create table t_user (
	user_id int PRIMARY key auto_increment,
	name varchar(30) not null,
	email varchar(100) ,
	avatar varchar(200),
	passwd varchar(20) not null,
	role int,
	editor varchar(30)
);

ALTER TABLE t_user ADD unique (name);

-- 建立笔记表

create table t_note(
	note_key varchar(30),
	user_id int,
	title varchar(100),
	summary varchar(200),
	content varchar(200),
	source varchar(200),
	editor varchar(400),
	files varchar(400),
	visit int,
	praise int
);



create table t_raise_log(
	note_key    varchar(30),
	user_id int,
	type   varchar(10),
	flag   bool
);


create table t_message(
	msg_key    varchar(30),
	note_key    varchar(30),
	user_id int,
	content varchar(200),
	Praise  int
);