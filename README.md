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


***Flow Aplikasi***
1. Registrasi User
   
   Endpoint: POST /api/auth/register

   Request:
   ```
   {
       "first_name": "Yudistira",
       "last_name": "Rivaldi",
       "email": "yudistira@gmail.com",
       "password": "password",
       "date_of_birth": "1995-07-01",
       "gender": "L"
   }
   ```
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Registration successful"
   }
   ```
 
2. Login User
   
   User login untuk mendapatkan token JWT yang digunakan untuk semua endpoint selanjutnya.

   Endpoint: POST /api/auth/login
   
   Request: 
   ```
   {
       "email": "yudistira@gmail.com",
       "password": "password"
   }
   ```
   
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Login successful",
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb25zdW1lcl9pZCI6NiwiZXhwIjoxNzUwMzMyMTE0fQ.T2kVE3uPL6TK6jQk7TuLSRnyFSg_pnnDBn9JEWDD31U"
   }
   ```

3. GET Profil User
   
   Endpoint: GET /api/users

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```
   
   Response:
   ```
   {
       "responseCode": "00",
       "message": "User profile retrieved successfully",
       "data": {
           "ID": 1,
           "FirstName": "Yudistira",
           "LastName": "Rivaldi",
           "Email": "yudistira@gmail.com",
           "Password": "",
           "DateOfBirth": "1995-07-01T00:00:00Z",
           "Gender": "L",
           "CreatedAt": "0001-01-01T00:00:00Z",
           "UpdatedAt": "0001-01-01T00:00:00Z"
       }
   }
   ```

4. Update Data User
   
   Endpoint: PUT /api/users

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "first_name": "Yudistira Update",
       "last_name": "Rivaldi",
       "email": "yudistira@gmail.com",
       "password": "password",
       "date_of_birth": "1995-07-01",
       "gender": "L"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "User updated successfully"
   }
   ```

5. GET Data Categories
   
   Endpoint: PUT /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Success",
       "data": {
           "ID": 1,
           "Name": "ATK",
           "Description": "Alat Tulis Kantor",
           "CreatedAt": "0001-01-01T00:00:00Z",
           "UpdatedAt": "0001-01-01T00:00:00Z"
       }
   }
   ```

6. GET Data Categories Detail
   
   Endpoint: PUT /api/categories

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Success",
       "data": [
           {
               "ID": 1,
               "Name": "ATK",
               "Description": "Alat Tulis Kantor",
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z"
           }
       ]
   }
   ```

7. Create Data Categories
   
   Endpoint: POST /api/categories

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "name": "ATK",
       "Description": "Alat Tulis Kantor"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category created successfully"
   }
   ```

8. Update Data Categories
   
   Endpoint: PUT /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "name": "ATK Anak",
       "Description": "Alat Tulis Kantor Anak"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category updated successfully"
   }
   ```

9. Delete Data Categories
   
   Endpoint: DELETE /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category deleted successfully"
   }
   ```

10. GET Data Categories
   
   Endpoint: PUT /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Success",
       "data": {
           "ID": 1,
           "Name": "ATK",
           "Description": "Alat Tulis Kantor",
           "CreatedAt": "0001-01-01T00:00:00Z",
           "UpdatedAt": "0001-01-01T00:00:00Z"
       }
   }
   ```

11. GET Data Categories Detail
   
   Endpoint: PUT /api/categories

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Success",
       "data": [
           {
               "ID": 1,
               "Name": "ATK",
               "Description": "Alat Tulis Kantor",
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z"
           }
       ]
   }
   ```

12. Create Data Categories
   
   Endpoint: POST /api/categories

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "name": "ATK",
       "Description": "Alat Tulis Kantor"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category created successfully"
   }
   ```

13. Update Data Categories
   
   Endpoint: PUT /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "name": "ATK Anak",
       "Description": "Alat Tulis Kantor Anak"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category updated successfully"
   }
   ```

14. Delete Data Categories
   
   Endpoint: DELETE /api/categories/:id

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Category deleted successfully"
   }
   ```

   


