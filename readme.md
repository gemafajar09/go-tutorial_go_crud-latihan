<center>
    <h1>Panduan Membuat New Project Go</h1>
</center>

- membuat folder project

- membuat go mod dengan perintah :
    ``
        go mod init github.com/nama project atau nama folder project
    ``

    contoh :
    ``
        go mod init github.com/portofolio
    ``

- package atau depedensi yg dibutuhkan
    ``
        go get github.com/jinzhu/gorm
    ``
    ``
        go get github.com/go-sql-driver/mysql
    ``
    ``
        go get github.com/gin-gonic/gin
    ``

- jika semuanya sudah silahkan lihat komentar pada contoh aplikasi di atas.