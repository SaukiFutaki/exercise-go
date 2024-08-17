# Pemrograman Backend Lanjutan

## Student Portal Concurrent

### Description

Program ini adalah sebuah student portal sederhana yang digunakan untuk melakukan registrasi mahasiswa baru, login, mengambil program studi dari suatu mahasiswa, dan mengubah program studi dari suatu mahasiswa.

Implementasi mencakup struct `InMemoryStudentManager`, yang memiliki daftar internal mahasiswa dan program studinya, interface `StudentManager` yang mendefinisikan fungsionalitas student portal, dan kami merekomendasikan untuk penggunaan teknik concurrency pada proyek ini agar proses yang berat tidak membebani CPU dan memori.

#### Get Students

Fungsi `GetStudents` tidak menerima input dan akan mengembalikan seluruh data mahasiswa dalam bentuk slice `[]Student`.

- Jika tidak ada mahasiswa yang terdaftar, fungsi akan mengembalikan slice kosong.

#### Login

Fungsi `Login` menerima input `id` dan `name` ber-_type_ `string` yang merupakan id dan nama mahasiswa. Fungsi ini akan memeriksa apakah data mahasiswa yang terdapat dalam slice `students`.

- Jika input `id` atau `name` kosong, fungsi akan mengembalikan pesan `""` dan error `"ID or Name is undefined!"`.
- Jika ditemukan, fungsi akan mengembalikan pesan `"Login berhasil: Selamat datang <nama>! Kamu terdaftar di program studi: <program studi>.
"` beserta error `nil`.
- Jika tidak ditemukan, fungsi akan mengembalikan pesan `""` dan error `"Login gagal: data mahasiswa tidak ditemukan"`.

Contoh:

```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

- Kamu perlu melakukan pengecekan percobaan login dengan ID yang sama, jika sudah melebihi batas maksimum, ID tersebut akan diblokir dari proses login.

Untuk menjaga keamanan data mahasiswa, maka proses login yang gagal kita batasi maksimum 3 kali saja. Setelah 3 kali gagal, ID akan diblokir dan proses login dengan ID tersebut akan ditolak.

Apabila sebelum 3 kali proses login, ada 1 login yang berhasil, maka kita akan reset percobaan gagal ID tersebut menjadi 0.

Berikut langkah-langkah yang perlu dilakukan:

- Kamu perlu menambahkan sebuah `map[string]int` dengan nama `failedLoginAttempts` pada `InMemoryStudentManager`.
- Kamu juga perlu menginisialisasi `failedLoginAttempts` di function `NewInMemoryStudentManager`.
- Tambahkan pada function `Login` fungsionalitas `failedLoginAttempts` untuk menyimpan ID dan berapa kali proses login sudah dilakukan untuk ID tersebut.
- Jika proses login untuk sebuah ID sudah lebih dari 3 kali, maka berikan pesan error `Login gagal: Batas maksimum login terlampaui`.
- Jika proses login berhasil, maka reset data ID tersebut di `failedLoginAttempts` jika ada.

#### Register

Fungsi `Register` menerima input `id`, `name`, dan `studyProgram` ber-_type_ `string` yang merupakan id, nama, dan program studi mahasiswa yang akan diregistrasikan. Fungsi ini akan menambahkan data mahasiswa baru ke dalam slice `Students`.

- Jika input `id`, `name`, atau `studyProgram` kosong, program mengembalikan pesan `""` dan error `"ID, Name or StudyProgram is undefined!"`
- Jika kode program studi yang dimasukkan tidak terdapat dalam map `studentStudyPrograms`, program mengembalikan pesan `""` dan error `"Study program <studyProgram> is not found"`.
- Jika id mahasiswa yang dimasukkan sudah terdapat di dalam slice `Students`, program mengembalikan pesan `""` dan error `"Registrasi gagal: id sudah digunakan"`.
- Jika berhasil, fungsi akan mengembalikan pesan `"Registrasi berhasil: name (studyProgram)"` dan error `nil`.

Contoh penggunaan:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Register("C1234", "Cindy", "SI") // Output: "Registrasi berhasil: Cindy (Sistem Informasi)", nil
```

#### Get Study Program

Fungsi `GetStudyProgram` menerima input code yang merupakan kode program studi. Fungsi ini akan mencari nama program studi dari map `studentStudyPrograms` berdasarkan kode program studi yang dimasukkan.

- Jika input code kosong, program mengembalikan pesan `""` dan error `"Code is undefined!"`
- Jika kode program studi yang dimasukkan tidak ditemukan, program mengembalikan pesan `""` dan error `"Kode program studi tidak ditemukan"`.
- Jika berhasil, fungsi akan mengembalikan nama program studi yang dicari dan error `nil`.

Contoh penggunaan:

```go
sm := NewInMemoryStudentManager()
message, err := sm.GetStudyProgram("TK") // Output: "Teknik Komputer", nil
```

#### Modify Student

Fungsi `ModifyStudent` menerima input `name` yang merupakan nama mahasiswa dan `fn` yang merupakan sebuah fungsi `studentModifier`. Fungsi ini akan mencari data mahasiswa berdasarkan nama yang dimasukkan, dan akan memodifikasinya menggunakan fungsi `fn`.

- Jika mahasiswa dengan nama yang dimasukkan tidak ditemukan, program mengembalikan pesan `""` dan error `"Mahasiswa tidak ditemukan."`.
- Jika terdapat error pada fungsi `fn`, program akan mengembalikan pesan `""` dan error tersebut.
- Jika berhasil, fungsi akan mengembalikan pesan `"Program studi mahasiswa berhasil diubah."` dan error `nil`.

Contoh penggunaan:

```go
sm := NewInMemoryStudentManager()
message, err := sm.ModifyStudent("Dito", sm.ChangeStudyProgram("SI")) // Output: "Program studi mahasiswa berhasil diubah.", error
```

#### Change Study Program

Fungsi `ChangeStudyProgram` menerima input `programStudi` yang merupakan kode program studi yang akan diganti. Fungsi ini akan mengembalikan sebuah fungsi dengan type `model.tudentModifier` yang akan memodifikasi program studi dari mahasiswa yang diberikan.

- Jika kode program studi yang dimasukkan tidak ditemukan, program akan mengembalikan error `"Kode program studi tidak ditemukan"`.
- Jika berhasil, fungsi dengan type `model.StudentModifier` yang dihasilkan akan mengubah program studi dari mahasiswa yang diberikan.

Contoh penggunaan:

```go
sm := NewInMemoryStudentManager()
modifier := sm.ChangeStudyProgram("SI")
sm.ModifyStudent("Dito", modifier) // "Program studi mahasiswa berhasil
```

#### Import Student

Kamu diberikan sebuah function `ReadStudentsFromCSV()` yang menerima satu parameter yaitu `filename` dan hasil dari function tersebut dalam bentuk slice `[]Student` yang disimpan di variable `students`

Data untuk import sudah diberikan dalam 3 file, masing-masing berisi 1.000 data yaitu `students1.csv`, `students2.csv`, `students3.csv` dan diterima di function `ImportStudents` sebagai slice `filenames`.

Disarankan untuk menggunakan teknik `concurrency` yang sudah kamu pelajari di proyek ini.

Silahkan kamu lanjutkan function `ImportStudents` agar data bisa di-registrasi dengan cepat dan tepat. Dimana kalian diharapkan melakukan import data yang berada di slice `filenames` menggunakan `ReadStudentsFromCSV()` dan melakukan registrasi menggunakan Fungsi `Register` yang sudah kamu miliki.

### Submit Assignment

Setiap siswa dapat mengirimkan tugas ke portal. Hanya saja dikarenakan masing-masing tugas merupakan file yang cukup besar, maka kamu akan diharapkan untuk menggunakan teknik `job queue` untuk meng-handle pengiriman tugas dari siswa.

`job queue` secara sederhana adalah ketika kita membuat jumlah `goroutine` yang terbatas sehingga misal kita memiliki 1000 tugas yang harus dikirimkan dan jumlah `goroutine` yang kita buat hanya 10, maka 10 `goroutine` tersebut akan meng-handle 10 tugas pertama, ketika salah satu `goroutine` sudah selesai maka `goroutine` tersebut akan meng-handle tugas selanjutnya dalam antrian. Ini ditujukan agar proses yang berat tidak membebani CPU dan memori.

Kamu ditugaskan untuk membuat 3 goroutine saja untuk meng-handle pengiriman tugas dari siswa, masing-masing goroutine wajib memanggil ke `SubmitAssignmentLongProcess()` untuk pemrosesan pengiriman tugas.

### Constraints

- `ID` mahasiswa harus unik, tidak boleh ada dua mahasiswa dengan `ID` yang sama.
- Saat melakukan login, program hanya memeriksa apakah `ID` dan `name` yang dimasukkan sudah ada di slice `students`. Program tidak memeriksa apakah `ID` dan `name` tersebut merupakan pasangan yang valid.
- Kode program studi (`studyProgram`) harus valid, artinya harus terdaftar di map `studentStudyPrograms`.
- Saat melakukan registrasi, program memeriksa apakah `ID` yang dimasukkan sudah ada di slice `students`. Jika sudah ada, registrasi akan gagal.
- Saat melakukan modifikasi data mahasiswa, program hanya memeriksa apakah nama mahasiswa sudah terdaftar di slice `students`. Jika nama mahasiswa ditemukan, program akan memanggil fungsi dengan type `model.StudentModifier` yang diberikan sebagai argumen dan mengirimkan data mahasiswa yang bersesuaian.

### Test Case Examples

#### Test Case 1

**Input**:

```go
sm := NewInMemoryStudentManager()
sp, err := sm.GetStudyProgram("TI")
```

**Expected Output / Behavior**:

```bash
"Teknik Informatika", nil
```

**Explanation**:

Method `GetStudyProgram` harus mengembalikan `"Teknik Informatika"` untuk kode masukan `"TI"`.

#### Test Case 2

**Input**:

```go
sm := NewInMemoryStudentManager()
res, err := sm.Register("C12345", "Charlie", "SI")
```

**Expected Output / Behavior**:

```bash
"Registrasi berhasil: Charlie (SI)", nil
```

**Explanation**:

Method `Register` harus mendaftarkan mahasiswa baru dengan ID `C12345`, name `Charlie`, dan studyProgram `SI`. Sehingga output yang diharapkan adalah `"Registrasi Berhasil: Charlie (SI)", nil`.

#### Test Case 3

**Input**:

```go
sm := NewInMemoryStudentManager()
res, err := sm.ModifyStudent("Aditira", sm.ChangeStudyProgram("SI"))
```

**Expected Output / Behavior**:

```bash
"Program studi mahasiswa berhasil diubah.", nil
```

**Explanation**:

Metode `ModifyStudent` harus mengubah program studi siswa dengan nama "Aditira" menjadi "SI". Oleh karena itu output yang diharapkan adalah `"Program studi mahasiswa berhasil diubah.", nil`.

#### Test Case 4

**Input**:

```go
sm := NewInMemoryStudentManager()
err := sm.ImportStudents([]string{"students1.csv", "students2.csv", "students3.csv"})
```

**Expected Output / Behavior**:

Tidak ada error, i.e. err should be nil.

**Explanation**:

Fungsi `ImportStudents` akan membaca isi dari semua file CSV dan mendaftarkan data student tersebut menggunakan fungsi `Register`. Kalau tidak ada error, maka `err` akan berisi `nil`, dan jumlah mahasiswa di `manager` setelah import akan bertambah sesuai dengan jumlah dari dari file CSV yaitu sebanyak 3000+3 data bawaan awal. Proses import sendiri harus berlangsung cepat, kurang dari 100ms (100 milidetik).

#### Test Case 5

**Input**:

```go
sm := NewInMemoryStudentManager()
sm.SubmitAssignments(10)
```

**Expected Output / Behavior**:

```bash
=== Submit Assignment ===
Enter the number of assignments you want to submit: 10
Worker 3: Processing assignment 1
Worker 1: Processing assignment 2
Worker 2: Processing assignment 3
Worker 2: Finished assignment 3
Worker 3: Finished assignment 1
Worker 3: Processing assignment 5
Worker 1: Finished assignment 2
Worker 1: Processing assignment 6
Worker 2: Processing assignment 4
Worker 2: Finished assignment 4
Worker 2: Processing assignment 7
Worker 3: Finished assignment 5
Worker 3: Processing assignment 8
Worker 1: Finished assignment 6
Worker 1: Processing assignment 9
Worker 1: Finished assignment 9
Worker 1: Processing assignment 10
Worker 2: Finished assignment 7
Worker 3: Finished assignment 8
Worker 1: Finished assignment 10
Submitting 10 assignments took 122.333986ms
Press any key to continue...
```

**Explanation**:

Fungsi `SubmitAssignments` akan berjalan sebanyak jumlah assignment yang diberikan di user input melalui terminal. Disimulasikan bahwa setiap assignment perlu diproses selama 3 detik. Perhatikan bahwa output di atas hanyalah contoh, kemungkinan besar urutan dari Finished Assignment akan acak.

#### Test Case 6

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "Aditira")
```

**Expected Output / Behavior**:

```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

**Explanation**:

Sudah cukup jelas, lihat bagian Description.

#### Test Case 7

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: Batas maksimum login terlampaui
```

**Expected Output / Behavior**:

Pada login pertama sampai ketiga:

```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

Pada login keempat dan seterusnya:

```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: Batas maksimum login terlampaui

Press any key to continue...
```

**Explanation**:

Pada saat login pertama sampai ketiga, dengan ID yang sama, namun dengan nama yang salah, sistem masih memperbolehkan login. Namun saat login keempat dan seterusnya, ID tersebut tidak bisa login.

#### Test Case 8

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "Aditira") // Output: Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
```

**Expected Output / Behavior**:

Pada login pertama sampai kedua:

```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

Pada login ketiga:

```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

Pada login keempat dan kelima:

```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

**Explanation**:

Pada saat login pertama dan kedua, login dilakukan dengan ID yang sama, namun dengan nama yang salah. Pada login ketiga, login dilakukan dengan ID dan nama yang benar dan mahasiswa berhasil login. Kemudian pada login keempat dan kelima, ID tersebut melakukan login namun dengan nama yang salah, di sini tidak terjadi pemblokiran login, karena pada saat login ketiga dengan ID dan nama yang benar, sudah kita reset menjadi 0, sehingga percobaan berikutnya dihitung dari 1, bukan dari 3.
