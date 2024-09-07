### JAWABAN SOAL TEORI SUDAH ADA DI REPO INI
1. **Jawaban Soal Teori**:
    ```bash
    git clone https://github.com/gemgum/qoin-technical_test.git

### Deskripsi APLIKASI PRAKTEK DICE-GAME

Permainan dadu ini adalah sebuah aplikasi berbasis Golang yang mengimplementasikan permainan sederhana untuk beberapa pemain. Setiap pemain akan melempar sejumlah dadu pada setiap giliran, melakukan evaluasi hasil lemparan, dan mengikuti aturan khusus yang berlaku pada angka dadu tertentu. Permainan ini didasarkan pada konsep yang dibuat oleh Gilang, dan berakhir ketika hanya tersisa satu pemain yang masih memiliki dadu. Pemain dengan poin terbanyak akan menjadi pemenang.

### Fitur Utama
- **Input pemain dan jumlah dadu**: Menentukan jumlah pemain (`N`) dan jumlah dadu per pemain (`M`).
- **Evaluasi hasil lemparan dadu**:
  - Angka 6 menjadi poin bagi pemain.
  - Angka 1 diberikan ke pemain sebelah.
  - Angka 2, 3, 4, dan 5 tetap dimainkan.
- **Permainan berlangsung hingga tersisa satu pemain dengan dadu**.
- **Pemenang ditentukan berdasarkan poin terbanyak**.

### Cara Instalasi
1. **Clone Repository**:
   Clone proyek ini ke mesin lokal Anda dengan perintah:
   ```bash
   git clone https://github.com/gemgum/qoin-technical_test.git
Navigasi ke Direktori Proyek: Pindah ke direktori proyek:

### Cara Run Aplikasi
1. **Pindah Direktori**:
    ```bash    
    cd dicegame
Jalankan Permainan: Pastikan Golang terinstal di komputer Anda. Jalankan perintah di bawah ini untuk memulai permainan:

2. **Run main.go**:
    ```bash
    go run main.go
Cara Menggunakan
Jalankan Aplikasi: Ketika menjalankan aplikasi, Anda akan diminta untuk memasukkan jumlah pemain (N) dan jumlah dadu (M).
Lihat Proses Permainan: Setiap giliran akan menampilkan lemparan dadu, hasil evaluasi, dan status setiap pemain.
Menangkan Permainan: Pemain yang memiliki poin terbanyak setelah semua pemain lain kehabisan dadu dinyatakan sebagai pemenang.
Contoh Penggunaan
Jika permainan dimulai dengan 3 pemain dan setiap pemain memiliki 4 dadu, permainan akan berjalan seperti ini:

3. **Contoh**:
    ```bash
    Pemain = 3, Dadu = 4
    ==================
    Contoh:
    Pemain = 3, Dadu = 4
    ==================
    Giliran 1 lempar dadu:
    Pemain #1 (0): 3,6,1,3
    Pemain #2 (0): 2,4,5,5
    Pemain #3 (0): 1,2,5,6
    Setelah evaluasi:
    Pemain #1 (1): 3,3,1
    Pemain #2 (0): 2,4,5,5,1
    Pemain #3 (1): 2,5
    ==================
    Giliran 2 lempar dadu:
    Pemain #1 (1): 1,2,6
    Pemain #2 (0): 4,3,1,3,3
    Pemain #3 (1): 1,6
    Setelah evaluasi:
    Pemain #1 (2): 2,1
    Pemain #2 (0): 4,3,3,3,1
    Pemain #3 (2): 1
    ==================
    Giliran 3 lempar dadu:

    Pemain #1 (2): 6,1
    Pemain #2 (0): 2,5,6,4,6
    Pemain #3 (2): 1
    Setelah evaluasi:
    Pemain #1 (3): 1
    Pemain #2 (2): 2,5,4,1
    Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
    ==================
    Giliran 4 lempar dadu:
    Pemain #1 (3): 1
    Pemain #2 (2): 3,4,5,5
    Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
    Setelah evaluasi:
    Pemain #1 (3): _ (Berhenti bermain karena tidak memiliki dadu)
    Pemain #2 (2): 3,4,5,5
    Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
    ==================
Game berakhir karena hanya pemain #2 yang memiliki dadu.
Game dimenangkan oleh pemain #1 karena memiliki poin lebih banyak dari pemain lainnya.