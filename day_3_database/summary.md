# Summary Day 3

Hari ketiga, kita belajar mengenai database dan untuk database yang kita gunakan adalah PostgreeSQL. Kita belajar mengenai basic daripada Relational Database yang meliputi:

### **DDL dan DML**

**DDL** adalah Data Definition Language, yang pada dasarnya merupakan query-query yang digunakan untuk membuat, ataupun mengelola struktur daripada database, seperti schema, tabel, constraint, dan lain-lain.\
Beberapa contoh DDL:
|Contoh| Query | Kegunaan |
|-| --- | - |
|**Create**| `create <database_name>;`| Digunakan untuk membuat database dan juga tabel|
|**Alter**| `alter table <table_name> rename to <table_name_change>;`| Digunakan untuk menambah, menghapus, ataupun menubah kolom pada tabel yang dipilih|
|**Drop**|`drop table <table_name>;`| Digunakan untuk menghapus tabel|

**DML** merupakan Data Manipulation Language, query yang digunakan untuk memanipulasi data.\
Beberapa contoh DML:
|Contoh| Query | Kegunaan |
|-| --- | - |
|**Insert**| `insert into public.<database_name> (name, age) values ('dendi', 20);`| Memasukkan tada kedalam tabel|
|**Update**| `update public.<table_name> set age = 25 where id = 1;`| Mengupdate data yang ada|
|**Delete**|`delete from public.<table_name> where id = 1;`| Menghapus data pada tabel|

### **Join**

Join ini berfungsi untuk menggabungkan data dari dua tabel, juga pada join ini memiliki banyak jenis seperti

- Right join
- Left join
- Full join
- Inner join

### **Aggregation**

Secara garis besar, agregat ini merupakan fungsi-fungsi yang digunakan untuk menghitung nilai data pada database kita. Beberapa contoh agregat adalah count, sum, avg, dan masih sangat banyak yang lainnya.\
Teman-teman bisa mengakses lebih pada link berikut https://www.postgresql.org/docs/9.5/functions-aggregate.html

### **Subquery**

Subquery ini merupakan query dalam query, yang digunakan untuk membatasi data yang nantinya akan teman-teman ambil.\
Contoh: `update products set stock = subquery.stock - 2 from(select id, stock from products where id = 1) as subquery where products.id = 1`\
Query diatas digunakan untuk mengurangi jumlah stock sebanyak 2, yang mana id daripada produk tersebut = 1.

### **Function**

Function ini digunakan untuk menyingkat query, sama seperti halnya bahasa pemrograman lain, kita hanya perlu memanggil function yang telah kita buat sebelumnya.
