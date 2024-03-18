CREATE TABLE EndUser (
    IDUser INT PRIMARY KEY, Nama VARCHAR(255), Alamat VARCHAR(255), Email VARCHAR(255), Deskripsi TEXT, PengalamanKerja TEXT
);

CREATE TABLE LowonganPekerjaan (
    IDLowongan INT PRIMARY KEY, Judul VARCHAR(255), Deskripsi TEXT, JenisBidangPekerjaan VARCHAR(255)
);

CREATE TABLE Lamaran (
    IDLamaran INT PRIMARY KEY, IDUser INT, IDLowongan INT, DokumenCV BLOB, CoverLetter TEXT, StatusLamaran VARCHAR(50), FOREIGN KEY (IDUser) REFERENCES EndUser (IDUser), FOREIGN KEY (IDLowongan) REFERENCES LowonganPekerjaan (IDLowongan)
);

CREATE TABLE Rekruter (
    IDRekruter INT PRIMARY KEY, NamaPerusahaan VARCHAR(255), EmailPerusahaan VARCHAR(255)
);