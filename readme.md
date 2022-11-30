# Cara Pakai

Setelah melakukan clone pada project, ubah kode pada `config/config.json` untuk disesuaikan dengan database PC. Jalankan perintah `go run index.go start` untuk memulai server.
<br>
Input untuk setiap URL pada API bisa dilihat di `routes/api.go` pada bagian yang diberi comment.

<br>

# Migrasi Database

Jalankan perintah CLI (untuk mysql):

```
migrate -path db/migrations -database "mysql://DB_USERNAME:DB_PASSWORD@tcp(DB_HOST:DB_PORT)/DB_DATABASE" up
```

Pada variabel `DB_USERNAME`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT` dan `DB_DATABASE` bisa disesuaikan dengan kondisi PC.
<br>
Perintah migrate dapat dipasang dengan bantuan <a href="https://scoop.sh/">Scoop</a> untuk Windows. Dokumentasi lengkap untuk database selain mysql bisa dilihat di <a href="https://github.com/golang-migrate/migrate">golang-migrate</a>. Untuk saat ini api hanya mendukung koneksi ke database mysql yang di atur pada file `index.go`.

<br>

# DB Seeder

Sebagai permulaan untuk mencoba query perlu data dummy yang di generate oleh sistem. DB seeder bisa dilakukan di CLI dengan memasukkan perintah di bawah:

```
go run index.go seed
```

Untuk kali ini DB seeder hanya dilakukan pada tabel users, tabel books, dan tabel stocks.
<br>

# Module Error

```
go mod tidy
```
