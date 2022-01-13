## TODO API Apps

Pada aplikasi ini kita dapat memasukan satu atau semua project pada sebuh ToDo list.

# endpoint pada fitur user :

```
http://localhost:8080/register
```
`POST`     Pada method ini kita bisa mendaftarkan akun untuk dapat mengakses fitur ToDo dan fitur Project.

```
http://localhost:8080/login
```
`POST`     Pada method ini kita melakukan login agar dapat melakukan autentikasi pada user yang sudah kita daftarkan dan dapat mengakses semua fitur.

```
http://localhost:8080/users/profile
```
`GET`     endpoint ini digunakan untuk dapat melihat data user yang telah kita daftarkan.

```
http://localhost:8080/users/delete
```
`DELETE`  ini digunakan jika kita ingin menghapus data user kita.
```

http://localhost:8080/users/update
```
`PUT`     fitur ini digukan untuk melakukan perubahan pada data user kita.

# endpoint pada fitur ToDo :
```
http://localhost:8080/todos
```
`GET`     fitur ini digunakan unutuk mengumpulkan semua todo yang telah ditambahkan.

`POST`    fitur ini digunakan untuk menambahkan Todo.

```
http://localhost:8080/todo/[:id]
```
`GET`     digunakan untuk menampilkan data pada Todo.

`PUT`     digukanan untuk mengedit Todo.

`DELETE`  digunakan untuk menghapus Todo.

```
http://localhost:8080/:id/complete
```
`POST`     endpoint yg digunakan untuk set status Todo sudah terselesaikan.

```
http://localhost:8080/[:id]/reopen
```
`POST`     endpoint yg digunakan untuk set status 'Reopen' Todo yang sudah terselesaikan.

# endpoint pada fitur project :
```
http://localhost:8080/projects
```
`GET`   fitur ini digunakan unutuk mengumpulkan semua project yang telah dibuat.

`POST`  fitur ini digunakan untuk menambahkan project.

```
http://localhost:8080/projects/[:id]
```
`GET`     digunakan untuk menampilkan data project.

`PUT`     digukanan untuk mengedit data project.

`DELETE`  digunakan untuk menghapus project.

