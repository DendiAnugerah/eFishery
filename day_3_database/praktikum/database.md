# Praktikum Database

## Soal

Buatlah query untuk membuat tabel seperti di bawah ini.
Setelah itu buat query untuk insert, update, dan delete di masing-masing tabel.
Serta join di tabel order.
[](../screenshot/soal.png)

**Membuat tabel**

Query yang digunakan untuk membuat **tabel customers**.

`create table customers(id int primary key not null, customer_name char(50) not null);`

Hasil query:
![](../screenshot/tabel-customer.png)

Query yang digunakan untuk membuat **tabel products**;

`create table products(id int primary key not null, name char (50) not null);'

Hasil query:
![](../screenshot/tabel-products.png)

Query yang digunakan untuk membuat **tabel order**;

Hasil query:
![](../screenshot/tabel-orders.png)
