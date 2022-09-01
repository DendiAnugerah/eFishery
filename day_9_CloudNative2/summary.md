# Summary Day 9

## Container Orchestration

Sebuah automation tools untuk mengurangi operasional deployment, scaling, dan lain-lain. Salah satu container orchestration yang paling basic adalah docker compose. Docker compose memungkinkan kita menjalankan beberapa container sekaligus.

Docker Compose:

- Memungkinkan kita menjalankan beberapa aplikasi menggunakan docker-compose.yaml
- Menggunakan single command untuk mendeploy aplikasi

Service:
Service ini didefinisikan oleh Docker Image dan runtime argument.

Volume:
Tempat untuk penyimpanan data sementara pada docker.

Network:
Layer yang mana memperbolehkan serives berkomunikasi satu sama lain.

### Kelebihan Container Orchestration

1. Simplified operations
2. Resilience
3. Added security

Hashicorp Stack
Nomad
Consul
Vault

## CI/CD

Kelebihannya, feedback dan mendeteksi dapat berjalan lebih cepat.

Salah satu contoh aplikasi CI/CD adalah Github Action.

Di dalam github action terdapat github runner merupakan sebuah aplikasi yang dapat menjalankan logic pada github actionnya. Workflow dapat mengkonfigurasi mana saja kegiatan yang sedang berjalan.

## Observability

Suatu kegiatan untuk mengukur matriks keadaan suatu sistem. Nah beberapa kelebihan observability ialah,

- Monitoring performa suatu aplikasi
- DevSecOps dan SRE
- Monitoring inftrastruktur, cloud, dan kubernetes
- End-user experience
- Analisis Bisnis

Tantangan dari observability

- Data silos, data ini hanya tersimpan di suatu tempat saja
- Volume, velocity, variety, dan complexity
- Instrumentasi dan configurasi masih manual
- Susah untuk dites
- Menghabiskan waktu pada saat troubleshooting

Beberapa tools observability adalah, New Relic dan Datadog
