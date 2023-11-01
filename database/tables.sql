-- Tabla de Clientes
CREATE TABLE clientes (
    cliente_id SERIAL PRIMARY KEY,
    cliente_nombre VARCHAR(255),
    cliente_direccion VARCHAR(255),
    cliente_email VARCHAR(255),
    cliente_password VARCHAR(255),
    cliente_telefono VARCHAR(15)
);

-- Tabla de Mascotas
CREATE TABLE mascotas (
    mascota_id SERIAL PRIMARY KEY,
    mascota_nombre VARCHAR(255),
    mascota_especie VARCHAR(50)
);

-- Tabla intermedia entre Clientes y Mascotas (N:N)
CREATE TABLE clientes_mascotas (
    cliente_mascota_id SERIAL PRIMARY KEY,
    cliente_id INT,
    mascota_id INT,
    CONSTRAINT fk_client_id
        FOREIGN KEY (cliente_id)
            REFERENCES clientes (cliente_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE,
    CONSTRAINT fk_mascota_id
        FOREIGN KEY (mascota_id)
            REFERENCES mascotas (mascota_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE
);

-- Tabla de Servicios
CREATE TABLE servicios (
    servicio_id SERIAL PRIMARY KEY,
    servicio_nombre VARCHAR(255),
    servicio_costo DECIMAL(10, 2)
);

-- Tabla intermedia entre Mascotas y Servicios (N:N)
CREATE TABLE mascotas_servicios (
    mascota_servicio_id SERIAL PRIMARY KEY,
    mascota_id INT,
    servicio_id INT,
    CONSTRAINT fk_mascota_id
        FOREIGN KEY (mascota_id)
            REFERENCES mascotas (mascota_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE,
    CONSTRAINT fk_servicio_id_mascotas
        FOREIGN KEY (servicio_id)
            REFERENCES servicios (servicio_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE
);

-- Tabla de Pagos
CREATE TABLE pagos (
    pago_id SERIAL PRIMARY KEY,
    pago_fecha DATE,
    pago_cantidad DECIMAL(10, 2)
);

-- Tabla intermedia entre Pagos y Servicios (N:N)
CREATE TABLE pagos_servicios (
    pago_servicio_id SERIAL PRIMARY KEY,
    pago_id INT,
    servicio_id INT,
    CONSTRAINT fk_pago_id
        FOREIGN KEY (pago_id)
            REFERENCES pagos (pago_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE,
    CONSTRAINT fk_servicio_id_pagos
        FOREIGN KEY (servicio_id)
            REFERENCES servicios (servicio_id)
            ON DELETE SET NULL
            ON UPDATE CASCADE
);
