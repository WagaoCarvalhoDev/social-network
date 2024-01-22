CREATE TABLE users(
    id int auto_increment primary key,
    nameUser varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    passwd varchar(20) not null,
    createAt timestamp default current_timestamp()
) ENGINE=INNODB;