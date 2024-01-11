# Recycle Rush Backend

Aplikasi backend yang dibangun dengan menggunakan bahasa pemrograman Go (Golang) dan menerapkan clean architecture.

## Deskripsi

Aplikasi ini merupakan bagian dari aplikasi Recycle Rush yang bertanggung jawab untuk menyimpan data serta mengambil data dari database. Dikembangkan menggunakan bahasa pemrograman Go, aplikasi ini dirancang dengan menggunakan clean architecture untuk memastikan kelenturan, pemeliharaan, dan kejelasan struktural dalam pengembangan perangkat lunak.

## Komponen Utama

- **config**: Berisi konfigurasi aplikasi.
- **controller**: Modul untuk menangani logika pengendalian (controller) pada lapisan presentasi.
- **helper/error**: Modul berisi helper untuk penanganan kesalahan.
- **middleware**: Modul untuk menangani middleware, misalnya, otentikasi.
- **model**: Modul untuk mendefinisikan struktur data (model).
- **repository**: Modul untuk interaksi dengan sumber data eksternal (mis., database, API eksternal).
- **router**: Modul untuk menangani routing dan definisi endpoint.
- **service**: Modul untuk implementasi logika bisnis (service).
- **utils**: Modul berisi fungsi utilitas yang dapat digunakan secara umum.

