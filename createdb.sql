BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `eventi` (
	`idevento`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`codcliente`	INTEGER NOT NULL,
	`tipoevento`	TEXT,
	`timestamp`	TEXT,
	`tipoprodotto`	TEXT,
	`nomeprodotto`	TEXT,
	`pagato`	NUMERIC
);
CREATE TABLE IF NOT EXISTS `clienti` (
	`idcliente`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`codcliente`	INTEGER NOT NULL UNIQUE,
	`titolo`	TEXT,
	`cognome`	TEXT,
	`nome`	TEXT,
	`mail`	TEXT,
	`telefono1`	TEXT,
	`telefono2`	TEXT,
	`telefono3`	TEXT
);
COMMIT;
