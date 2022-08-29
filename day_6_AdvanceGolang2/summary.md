# Summary Day 6

REST API

## HTTP/HTTPS

HTTP Messages/body/payload merupakan bagaimana data dikirim antara server dan klien.
HTTP Header berisi informasi-informasi tambahan yang nantinya dibutuhkan pada sisi server.

## Methods

A. GET, untuk mengambil data pada server
B. POST, digunakan untuk create resource
C. PUT, digunakan untuk update data
D. DELETE, menghapus data

## Status Codes

2xx: Success - Berarti request yang dikirimkan oleh klien terus diterima.
200 - OK - Request dari sisi klien sukses.
201 - Created - Berhubungan dengan proses create.
202 - Accepted - Mengindikasi bahwa request telah diterima namun belum selesai.

4xx: Client Error - Request yang dikirimkan dari sisi klien tidak valid.
400 - Bad Request - Kalau server request dari klien tidak sesuai dari speksifikasi sisi server.
401 - Unauthorized - Memerlukan autentikasi informasi user.
403 - Forbidden - Unauthorized request.
404 - Not Found - Server tidak menemukan reource yang dituju.
405 - Method Not Allowed - Salah method.

5xx: Server Error - Error yang ada pada sisi server.
501 - Internal Server Error -  
502 - Bad Gateway - Invealid response pada saat gateway mendapatkan response yang diperlukan dalam menghandle request.
504 - Gateway Timeout

## JSON

Merupakan format data yang digunakan pada REST API. Javascript Object Notation (JSON)

## Authorization

Menggunakan JWT (JSON Web Token).

`<header>.<payload>.<signature>`

Header - Berisi

REST merupakan sebuah arsitektur yang digunakan untuk membuat web service, karena berjalan di protokol HTTP, yang mana HTTP ini biasa digunakan pada level WEB.

Tools Pemanggil API
How to use Postman
