***Flow Instalasi Aplikasi***
1. Clone Repository

   ```
   git clone https://github.com/yudistirarivaldi/technical-test-deeptech.git
   cd technical-test-kreditplus
   ```

2. Siapkan File .env
   Buat file .env atau gunakan yang sudah disediakan:
   
   DB_HOST=mysql untuk running host via docker

   ```
   SERVER_PORT=8080
   DB_HOST=mysql 
   DB_PORT=3306
   DB_USER=dev
   DB_PASS=dev123
   DB_NAME=technical_deep_tech
   JWT_SECRET=supersecretkey123
   ```

3. Jalankan via Docker Compose
   ```
   docker compose up --build
   ```
   Ini akan Build image aplikasi Jalankan container MySQL + aplikasi Otomatis baca .env dan mengatur konfigurasi

4. Stop Container
   ```
   docker compose down
   ```
