## Soal 1: A000124 Sequence Generator (Lazy Caterer Sequence)
![Screenshot output soal 1](https://github.com/user-attachments/assets/f3ccda60-de51-4a0f-b132-a2bdde3ddbe4)

#### Time Complexity: O(n)
- Dimana n adalah jumlah elemen dalam deret yang akan dihasilkan
- Kita melakukan iterasi sebanyak n kali untuk menghasilkan setiap elemen deret
- Setiap perhitungan elemen membutuhkan waktu O(1)
- Oleh karena itu, Time Complexity adalah O(n)

### Penjelasan

Program ini mengimplementasikan deret A000124 (Lazy Caterer Sequence) dengan rumus:
- a(n) = n(n+1)/2 + 1 untuk n ≥ 1
- a(0) = 1

Program menerima input berupa jumlah elemen deret yang ingin dihasilkan, kemudian menampilkan deret tersebut dengan format yang dipisahkan tanda hubung (misalnya: 1-2-4-7-11).

## Soal 2: Dense Ranking
![Screenshot output soal 2](https://github.com/user-attachments/assets/890b2b69-18b5-4493-9fab-15f711dd343b)

#### Time Complexity: O(n log n)
- Dimana n adalah jumlah skor dalam leaderboard
- Proses pengurutan skor membutuhkan waktu O(n log n)
- Proses pencarian peringkat untuk setiap skor GITS membutuhkan waktu O(m * n) dimana m adalah jumlah skor GITS
- Oleh karena itu, Time Complexity keseluruhan adalah O(n log n + m * n)

### Penjelasan

Program ini mengimplementasikan sistem Dense Ranking dimana:
- Peringkat tertinggi (1) diberikan kepada skor tertinggi
- Skor yang sama mendapatkan peringkat yang sama
- Peringkat berikutnya diberikan berdasarkan urutan nilai skor (bukan urutan jumlah pemain)

Program menerima input berupa daftar skor leaderboard dan daftar skor GITS, kemudian menghitung peringkat untuk setiap skor GITS berdasarkan sistem Dense Ranking.

## Soal 3: Balanced Brackets (Tanda Kurung Seimbang)
![Screenshot output soal 3](https://github.com/user-attachments/assets/373a3423-31fa-46f2-ba06-ea87613ee7df)

#### Time Complexity: O(n)
- Dimana n adalah panjang string input
- Kita melakukan iterasi melalui setiap karakter dalam string input tepat satu kali
- Setiap operasi di dalam loop (push ke stack, pop dari stack, pengecekan kondisi) membutuhkan waktu O(1)
- Oleh karena itu, Time Complexity adalah O(n)

#### Space Complexity: O(n)
- Dalam kasus terburuk, jika string input terdiri dari semua tanda kurung pembuka, kita perlu menyimpan semua tanda kurung pembuka dalam stack
- Oleh karena itu, Space Complexity adalah O(n)

### Penjelasan

Solusi ini menggunakan struktur data `Stack` untuk melacak tanda kurung pembuka dan mencocokkannya dengan tanda kurung penutup yang sesuai:

1. Kita melakukan iterasi melalui setiap karakter dalam string input
2. Jika kita menemukan tanda kurung pembuka, kita memasukkannya ke dalam stack
3. Jika kita menemukan tanda kurung penutup, kita:
   - Memeriksa apakah stack kosong (jika ya, kembalikan "NO" karena tidak ada tanda kurung pembuka yang cocok)
   - Mengambil elemen teratas dari stack
   - Memeriksa apakah elemen yang diambil cocok dengan tanda kurung penutup saat ini (jika tidak, kembalikan "NO")
4. Setelah memproses semua karakter, jika stack kosong, semua tanda kurung telah dicocokkan dengan benar (kembalikan "YES")
5. Jika stack tidak kosong, beberapa tanda kurung pembuka tidak ditutup (kembalikan "NO")

Pendekatan ini optimal untuk masalah ini dengan kompleksitas terendah yang mungkin, karena kita perlu memeriksa setiap karakter setidaknya satu kali untuk menentukan apakah tanda kurung seimbang.
