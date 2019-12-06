CREATE TABLE account(
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (355) UNIQUE NOT NULL
);
insert into account(username, password, email) values('ed.c', '12345', 'eddie.christian@iron.net');
insert into account(username, password, email) values('bill.d', '12345', 'bill.davis@iron.net');
insert into account(username, password, email) values('frank.b', '12345', 'frank.beamer@iron.net');