## Uy vazifasi: Dockerfiles bilan Go va Environment variables ishlatishi

### Maqsad:
`Docker` konteyneriga environment variablesni uzatish va ularni Go dasturida ishlatishni o'rganish.

### Talablar:

1. **Go dasturini yaratish**:
    - Port `9000` da tinglaydigan va xabarni qaytaradigan oddiy `Go` ilovasini rivojlantiring. Xabar environment variable orqali olinishi kerak.

2. **Dockerfile yozish**:
    - `Go` ilovasini konteynerlash uchun `Dockerfile` yarating.
    - `Dockerfile` `Go` ilovasini tuzish va ishga tushirish uchun kerakli ko'rsatmalarni o'z ichiga olganligiga ishonch hosil qiling.

3. **Environment variablesni Uzatish**:
    - Docker konteyneriga environment variablesni uzating.

4. **Docker Image Qurish**:
    - `Dockerfile` dan `Docker` image ni yaratish uchun `docker build` buyrug'idan foydalaning.

5. **Docker Konteynerini Ishga Tushirish**:
    - Qurilgan image dan `Docker` konteynerini ishga tushirish uchun `docker run` buyrug'idan foydalaning.
    - Konteynerning `9000` portini xost mashinaning `9000` portiga ulang.

6. **Ilovani Sinovdan O'tkazish**:
    - Ilovani `http://localhost:9000` ga kirib tekshiring.

7. **Topshiriqlar**:
    - `Go` ilovasi uchun `Dockerfile`.
    - `Docker` konteynerini qurish va ishga tushirish uchun ishlatilgan buyruqlar.
    - Ilova muvaffaqiyatli ishlayotganini ko'rsatadigan skrinshotlar.
    - `docker-compose.yml` fayli agar qo'llanilgan bo'lsa.



