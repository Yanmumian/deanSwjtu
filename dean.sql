create table students(
id int primary key auto_increment,
`name` varchar(20) not null,
number varchar(11) unique not null,
sex char(2) not null,
email varchar(50),
phone varchar(11),
address varchar(50),
class varchar(20),
department varchar(20),
regist_time datetime default current_timestamp comment '注册时间',
lastlog_time datetime default current_timestamp on update current_timestamp
)

