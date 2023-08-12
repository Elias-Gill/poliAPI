-- Script de creacion de la base de datos
-- tablas
CREATE TABLE usuarios (
    id_usuario serial4 NOT NULL,
    nombre varchar NOT NULL,
    email varchar NOT NULL,
    imagen varchar NOT NULL, -- se usando conversion a string
    password varchar NOT NULL,
    CONSTRAINT pk_usuario PRIMARY KEY (id_usuario)
);

CREATE TABLE horarios (
    id_horario serial4 NOT NULL,
    id_usuario integer NOT NULL,
    CONSTRAINT pk_horarios PRIMARY KEY (id_horario),
    CONSTRAINT fk_usuario_horario FOREIGN KEY (id_usuario) REFERENCES usuarios (id_usuario) on delete cascade
);

CREATE TABLE hojas_excel (
    id_excel serial NOT NULL,
    fecha_carga date NOT NULL,
    nombre_archivo varchar NOT NULL,
    CONSTRAINT pk_excel PRIMARY KEY (id_excel)
);

create table carreras (
    id_carrera serial not null,
    nombre varchar not null,
    archivo_excel integer not null,
    n_hoja int not null, -- numero de hoja en el excel
    CONSTRAINT pk_carrera PRIMARY KEY(id_carrera),
    CONSTRAINT fk_exel_carreras FOREIGN KEY (archivo_excel) REFERENCES hojas_excel (id_excel) on delete cascade
);

CREATE TABLE materias (
    id_materia serial4 NOT NULL,
    descripcion varchar NOT NULL,
    archivo_excel integer NOT NULL,
    carrera integer NOT NULL,
    -- horarios de clase
    lunes varchar,
    martes varchar,
    miercoles varchar,
    jueves varchar,
    viernes varchar,
    -- horarios de examenes
    parcial1 varchar,
    parcial2 varchar,
    final1 varchar,
    final2 varchar,
    --
    CONSTRAINT pk_materias PRIMARY KEY (id_materia),
    CONSTRAINT fk_materias_excel FOREIGN KEY (archivo_excel) REFERENCES hojas_excel (id_excel) on delete cascade,
    CONSTRAINT fk_materias_carrera FOREIGN KEY (carrera) REFERENCES carreras (id_carrera) on delete cascade
);

CREATE TABLE detalle_horario (
    id_horario integer NOT NULL,
    id_materia integer NOT NULL,
    CONSTRAINT pk_detalles PRIMARY KEY (id_horario, id_materia),
    CONSTRAINT fk_horario_detalle FOREIGN KEY (id_horario) REFERENCES horarios (id_horario) on delete cascade,
    CONSTRAINT fk_materia_detalle FOREIGN KEY (id_materia) REFERENCES materias (id_materia) on delete cascade
);
