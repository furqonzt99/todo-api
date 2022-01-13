## TODO API Apps

Pada aplikasi ini kita dapat memasukan satu atau semua project pada sebuh ToDo list.

# endpoint pada fitur user :

```
http://localhost:8080/register
```
`POST` Pada method ini kita bisa mendaftarkan akun untuk dapat mengakses fitur ToDo dan fitur Project.

```
http://localhost:8080/login
```
`POST` Pada method ini kita melakukan login agar dapat melakukan autentikasi pada user yang sudah kita daftarkan dan dapat mengakses semua fitur

```
http://localhost:8080/users/profile
```
`GET` endpoint ini digunakan untuk dapat melihat data user yang telah kita daftarkan

```
http://localhost:8080/users/delete
```
`DELETE` ini digunakan jika kita ingin menghapus data user kita.
```

http://localhost:8080/users/update
```
`PUT` fitur ini digukan untuk melakukan perubahan pada data user kita.
