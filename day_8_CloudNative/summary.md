# Summary Day 8

## Cloud Native

Sebuah paradigma baru yang hanya dioptimisasi untuk modern distributed system (server), yang mana aplikasi tersebut bisa disclaing hingga ribuan.

### Modern Design

Sebuah aplikasi yang dibuat yang dengan konsep the twelve factor application

1. Codebase - terkait version control, buat sejelas mungkin agar tidak susah nantinya.
2. Dependencies - depensensi dideklarasikan secara khusus.
3. Configuration - untuk menyimpan config dari environment yang ada.
4. Backing service - sebisa mungkin kita tidak menggunakan IP.
5. Build, Release, and Run - pisahkan build, release, dan run.
6. Process - perubahan data harus disimpan di tempat yang seharusnya, sebisa mungkin proses itu stateless.
7. Port Binding - export servies via port binding.
8. Concurrency - diperlukan untuk menambah performance.
9. Disposability - bikin agar sistem dapat shutdown dengan sempurna, sehingga tidak mengganggu experience dari user.
10. Dev/Prod Parity - meskipun development, stagging, dan production berbeda, namun tetap life cyclenya sebisa mungkin harus sama.
11. Logging - bikin log sebaik mungkin, dan efisien.
12. Admin Process - sebisa mungkin pada saat kita menjalankan sebuah aplikasi tidak menggunakan admin.

### Containers

Ibaratnya kontainer pada umumnya, dockes ini bisa membungkus seluruh kebutuhan kode kita, sehingga nantinya dapat digunakan oleh developer lain, atau lebih mudahnya sebuah abstraksi pada layer aplikasi yang menjadi satu antara code dan dependensi lainnya.
