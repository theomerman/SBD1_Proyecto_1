package db

var QueryVector = []string{
	`CREATE TABLE IF NOT EXISTS sistemadevotos.PARTIDO (
		id_partido INT AUTO_INCREMENT NOT NULL,
		nombre_partido VARCHAR(100) NOT NULL,
		siglas VARCHAR(30) NOT NULL,
		fundacion DATE NOT NULL,
		PRIMARY KEY (id_partido)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.CARGO(
		id_cargo INT AUTO_INCREMENT NOT NULL,
		cargo VARCHAR(50) NOT NULL,
		PRIMARY KEY (id_cargo)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.DEPARTAMENTO(
		id_departamento INT AUTO_INCREMENT NOT NULL,
		nombre VARCHAR(20) NOT NULL,
		PRIMARY KEY (id_departamento)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.CIUDADANO(
		dpi VARCHAR(13) NOT NULL,
		nombre VARCHAR(50) NOT NULL,
		apellido VARCHAR(50) NOT NULL,
		direccion VARCHAR(100) NOT NULL,
		telefono VARCHAR(10) NOT NULL,
		edad int NOT NULL,
		genero VARCHAR(1) NOT NULL,
		PRIMARY KEY (dpi)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.CANDIDATO (
		id_candidato INT AUTO_INCREMENT NOT NULL,
		nombres VARCHAR(150) NOT NULL,
		fecha_nacimiento DATE NOT NULL ,
		PARTIDO_id_partido INT NOT NULL, 
		CARGO_id_cargo INT NOT NULL,
		PRIMARY KEY (id_candidato),
		FOREIGN KEY (PARTIDO_id_partido) REFERENCES PARTIDO(id_partido),
		FOREIGN KEY (CARGO_id_cargo) REFERENCES CARGO(id_cargo)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.MESA(
		id_mesa INT AUTO_INCREMENT NOT NULL,
		DEPARTAMENTO_id_departamento INT NOT NULL,
		PRIMARY KEY (id_mesa),
		FOREIGN KEY (DEPARTAMENTO_id_departamento) REFERENCES DEPARTAMENTO(id_departamento)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.VOTO(
		id_voto INT AUTO_INCREMENT NOT NULL,
		fecha_hora DATE NOT NULL,
		CIUDADANO_dpi VARCHAR(13) NOT NULL,
		MESA_id_mesa INT NOT NULL,
		PRIMARY KEY (id_voto),
		FOREIGN KEY (CIUDADANO_dpi) REFERENCES CIUDADANO(dpi),
		FOREIGN KEY (MESA_id_mesa) REFERENCES MESA(id_mesa)
	)`,
	`CREATE TABLE IF NOT EXISTS sistemadevotos.DETALLE_VOTO(
		id_detalle_voto INT AUTO_INCREMENT NOT NULL,
		VOTO_id_voto INT NOT NULL,
		CANDIDATO_id_candidato INT NOT NULL,
		PRIMARY KEY (id_detalle_voto),
		FOREIGN KEY (VOTO_id_voto) REFERENCES VOTO(id_voto),
		FOREIGN KEY (CANDIDATO_id_candidato) REFERENCES CANDIDATO(id_candidato)
	)`,
	// `
	// CREATE TEMPORARY TABLE sistemadevotos.temp_csv_data (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	password VARCHAR(255)
	// 	-- Add other columns as needed
	// )`,
}
