--Membuat database
CREATE DATABASE alexander_50420116_pert7;

--menggunakan database
use alexander_50420116_pert7;

--membuat tabel mahasiswa
CREATE TABLE mahasiswa (
    id int(6) unsigned AUTO_INCREMENT primary key,
    npm char(8) NOT NULL,
    nama varchar(30) NOT NULL,
    kelas char(5) NOT NULL,
    profile varchar(30)NOT NULL);

--menampilkan tabel mahasiswa
select * from mahasiswa;

--menampilkan database mahasiswa
desc mahasiswa;
