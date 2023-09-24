
-- /consulta1

USE sistemadevotos
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
SELECT c.nombres AS Presidente, c2.nombres AS "Vice Presidente", p.nombre_partido AS Partido
FROM candidato c
INNER JOIN partido p 
ON c.PARTIDO_id_partido = p.id_partido
LEFT JOIN candidato c2 
ON c.PARTIDO_id_partido = c2.PARTIDO_id_partido
AND c2.CARGO_id_cargo = 2
WHERE c.CARGO_id_cargo = 1
;


-------------------------------------------------
-- /consulta2

SELECT nombre_partido AS Partido,count(siglas) AS Cantidad
FROM CANDIDATO c
INNER JOIN PARTIDO p
ON c.PARTIDO_id_partido = p.id_partido
WHERE CARGO_id_cargo = 3
OR CARGO_id_cargo = 4
OR CARGO_id_cargo = 5
group by nombre_partido
;
-------------------------------------------------
-- /consulta3

SELECT nombres AS Nombre, nombre_partido AS Partido
FROM candidato
INNER JOIN partido
ON PARTIDO_id_partido = id_partido
WHERE CARGO_id_cargo = 6
ORDER BY nombres ASC
;

-------------------------------------------------
-- /consulta4

SELECT count(PARTIDO_id_partido) AS Cantidad, nombre_partido AS Partido
FROM candidato
INNER JOIN partido
ON PARTIDO_id_partido = id_partido
WHERE CARGO_id_cargo != '-1' 
GROUP BY nombre_partido
ORDER BY Cantidad
;

-------------------------------------------------
-- /consulta5

SELECT nombre AS Departamento, count(nombre) AS Voto
FROM voto
INNER JOIN MESA
ON MESA_id_mesa = id_mesa
INNER JOIN departamento
ON DEPARTAMENTO_id_departamento = id_departamento
GROUP BY nombre
ORDER BY Voto ASC
;

-------------------------------------------------
-- /consulta6
 
SELECT siglas as Votos, ceil(count(id_partido)/5) AS Cantidad
FROM voto
INNER JOIN detalle_voto
ON id_voto = VOTO_id_voto
INNER JOIN candidato
ON CANDIDATO_id_candidato = id_candidato
INNER JOIN PARTIDO
ON PARTIDO_id_partido = id_partido
group by CANDIDATO_id_candidato
order by CANDIDATO_id_candidato
limit 1
;

-------------------------------------------------
-- /consulta7

SELECT edad AS Edad, ceil(count(id_voto)/5) AS Cantidad
FROM detalle_voto
INNER JOIN voto
ON VOTO_id_voto = id_voto
INNER JOIN ciudadano
ON CIUDADANO_dpi = dpi
GROUP BY edad
order by Cantidad desc
limit 10
;

-------------------------------------------------
-- /consulta8

SELECT
	c.nombres AS Presidente, 
	c2.nombres AS "Vice Presidente",
	count(VOTO_id_voto) AS Votos
FROM candidato c
INNER JOIN detalle_voto
ON c.id_candidato = CANDIDATO_id_candidato
LEFT JOIN candidato c2 
ON c.PARTIDO_id_partido = c2.PARTIDO_id_partido
AND c2.CARGO_id_cargo = 2
WHERE c.CARGO_id_cargo = 1
GROUP BY c.nombres, c2.nombres
ORDER BY Votos desc
limit 10
;

-------------------------------------------------
-- /consulta9

SELECT id_mesa AS "No. Mesa", count(id_voto) AS Votos, nombre AS Nombre
FROM detalle_voto
INNER JOIN voto
ON VOTO_id_voto = id_voto
INNER JOIN mesa
ON MESA_id_mesa = id_mesa
INNER JOIN departamento
ON DEPARTAMENTO_id_departamento = id_departamento
GROUP BY MESA_id_mesa
order by Votos asc
limit 5
; 

-------------------------------------------------
-- /consulta10

SELECT fecha_hora AS Hora, ceil(count(id_detalle_voto)/5) AS Votos
FROM voto
INNER JOIN detalle_voto
ON VOTO_id_voto = id_voto
group by fecha_hora
order by Votos desc
limit 5
;

-------------------------------------------------
-- /consulta11

SELECT genero AS Genero, ceil(count(id_detalle_voto)/5) AS Votos
FROM ciudadano
INNER JOIN voto
ON CIUDADANO_dpi = dpi
INNER JOIN detalle_voto
ON VOTO_id_voto = id_voto
GROUP BY genero
;