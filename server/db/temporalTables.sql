DROP TABLE IF EXISTS sistemadevotos.TempPartido;

CREATE TEMPORARY TABLE if not exists sistemadevotos.TempPartido (
	id_partido INT NOT NULL,
	nombrePartido VARCHAR(200) NOT NULL,
	siglas VARCHAR(40) NOT NULL,
	fundacion DATE NOT NULL
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/partidos.csv'
INTO TABLE sistemadevotos.TempPartido
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS 
(id_partido, nombrePartido,Siglas,@Fundacion) 
SET Fundacion = str_to_date(@Fundacion, '%d/%m/%Y'); 

INSERT IGNORE INTO sistemadevotos.Partido (id_partido, nombre_partido,siglas,fundacion) 
SELECT id_partido, nombrePartido,Siglas,Fundacion 
FROM sistemadevotos.TempPartido;

DROP TABLE IF EXISTS sistemadevotos.TempPartido;

----------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS sistemadevotos.TempCargo;

CREATE TEMPORARY TABLE if not exists sistemadevotos.TempCargo (
	id INT,
	cargo VARCHAR(50)
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/cargos.csv'
INTO TABLE sistemadevotos.TempCargo
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS  
(id, cargo);  

INSERT IGNORE INTO sistemadevotos.cargo (id_cargo, cargo) 
SELECT id, cargo 
FROM sistemadevotos.TempCargo;

DROP TABLE IF EXISTS sistemadevotos.tempcargo;

----------------------------------------------------------------------------------------------

DROP TABLE IF EXISTS sistemadevotos.TempDepartamento;

CREATE TEMPORARY TABLE if not exists sistemadevotos.TempDepartamento (
	id INT NOT NULL,
	nombre VARCHAR(20) NOT NULL

);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/departamentos.csv'
INTO TABLE sistemadevotos.TempDepartamento
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS  
(id, nombre); 

INSERT IGNORE INTO sistemadevotos.Departamento (id_departamento, nombre) 
SELECT id, nombre 
FROM sistemadevotos.TempDepartamento;

DROP TABLE IF EXISTS sistemadevotos.TempDepartamento;

----------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS sistemadevotos.TempCiudadano;

CREATE TEMPORARY TABLE IF NOT EXISTS sistemadevotos.TempCiudadano (
    DPI VARCHAR(13) NOT NULL,
    Nombre VARCHAR(50) NOT NULL,
    Apellido VARCHAR(50) NOT NULL,
    Direccion VARCHAR(100) NOT NULL,
    Telefono VARCHAR(10) NOT NULL,
    Edad INT NOT NULL,
    Genero VARCHAR(10) NOT NULL
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/ciudadanos.csv'
INTO TABLE sistemadevotos.TempCiudadano
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS; 

INSERT IGNORE INTO sistemadevotos.ciudadano (DPI, Nombre, Apellido, Direccion, Telefono, Edad, Genero)
SELECT DPI, Nombre, Apellido, Direccion, Telefono, Edad, Genero 
FROM sistemadevotos.TempCiudadano;

DROP TABLE IF EXISTS sistemadevotos.TempCiudadano;

----------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS sistemadevotos.TempCandidatos; 

CREATE TEMPORARY TABLE if not exists sistemadevotos.TempCandidatos (
	id_candidato INT NOT NULL,
	nombres VARCHAR(150) NOT NULL,
	fecha_nacimiento DATE NOT NULL ,
    PARTIDO_id_partido INT NOT NULL, 
	CARGO_id_cargo INT NOT NULL
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/candidatos.csv'
INTO TABLE sistemadevotos.TempCandidatos
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS  
(id_candidato, nombres,@fecha_nacimiento,PARTIDO_id_partido,CARGO_id_cargo)
SET fecha_nacimiento = str_to_date(@fecha_nacimiento, '%d/%m/%Y'); 

INSERT IGNORE INTO sistemadevotos.candidato (id_candidato, nombres,fecha_nacimiento,PARTIDO_id_partido,CARGO_id_cargo) 
SELECT id_candidato, nombres,fecha_nacimiento,PARTIDO_id_partido,CARGO_id_cargo 
FROM sistemadevotos.TempCandidatos;

DROP TABLE IF EXISTS sistemadevotos.TempCandidatos; 

----------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS sistemadevotos.TempMesa; 

CREATE TEMPORARY TABLE if NOT EXISTS sistemadevotos.TempMesa (
	id_mesa INT NOT NULL,
	DEPARTAMENTO_id_departamento INT NOT NULL
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/mesas.csv'
INTO TABLE sistemadevotos.TempMesa
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS 
(id_mesa,DEPARTAMENTO_id_departamento);

INSERT IGNORE INTO sistemadevotos.mesa (id_mesa,DEPARTAMENTO_id_departamento) 
SELECT id_mesa,DEPARTAMENTO_id_departamento
FROM sistemadevotos.TempMesa;

DROP TABLE IF EXISTS sistemadevotos.TempMesa; 


----------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS sistemadevotos.TempVotos; 

CREATE TEMPORARY TABLE IF NOT EXISTS sistemadevotos.TempVotos (
    id_voto INT NOT NULL,
    id_candidato INT NOT NULL,
    dpi_ciudadano VARCHAR(13) NOT NULL,
    mesa_id INT NOT NULL,
    fecha_hora DATETIME NOT NULL
);

LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/votaciones.csv'
INTO TABLE sistemadevotos.TempVotos
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS  
(id_voto,id_candidato,dpi_ciudadano,mesa_id,@fecha_hora)
SET fecha_hora = str_to_date(@fecha_hora, '%d/%m/%Y %H:%i');  

INSERT IGNORE INTO sistemadevotos.voto (id_voto, fecha_hora,CIUDADANO_dpi,MESA_id_mesa)
SELECT id_voto, fecha_hora,dpi_ciudadano,mesa_id 
FROM sistemadevotos.TempVotos;

INSERT INTO sistemadevotos.detalle_voto (VOTO_id_voto, CANDIDATO_id_candidato)
SELECT id_voto, id_candidato 
FROM sistemadevotos.TempVotos;

DROP TABLE IF EXISTS sistemadevotos.TempVotos; 


