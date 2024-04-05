-- logs table stores all logs produced since the system startup.
CREATE TABLE logs
(
	`id`        INT AUTO_INCREMENT NOT NULL,
	`timestamp` TEXT          NOT NULL,
	`log_type`  ENUM (
		'room_entry',
		'room_exit',
		'temperature',
		'humidity',
		'emergency'
		)                          NOT NULL,
	`data`      JSON               NOT NULL NOT NULL,

	INDEX (`log_type`),
	PRIMARY KEY (`id`)
);

-- employees represents table with employees data.
CREATE TABLE employees
(
	`id`            INT AUTO_INCREMENT NOT NULL,
	`firstname`     VARCHAR(64)        NOT NULL,
	`lastname`      VARCHAR(64)        NOT NULL,
	`department_id` VARCHAR(64)        NOT NULL,
	`rfid`          TEXT               NOT NULL,

	PRIMARY KEY (`id`)
);