-- CREATE TABLES
CREATE TABLE IF NOT EXISTS locations (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(30) NOT NULL UNIQUE
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS tasks (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  description VARCHAR(100) NOT NULL,
  location_id INT(11) NOT NULL,
  completed BOOLEAN DEFAULT FALSE,
  FOREIGN KEY(location_id) REFERENCES locations(id) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS workers (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(30) NOT NULL UNIQUE,
  hourly_wage DECIMAL(5, 2) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS logged_time (
  id INT(11) AUTO_INCREMENT PRIMARY KEY,
  time_hours INT(11) NOT NULL,
  task_id INT(11) NOT NULL,
  worker_id INT(11) NOT NULL,
  FOREIGN KEY(task_id) REFERENCES tasks(id) ON DELETE CASCADE,
  FOREIGN KEY(worker_id) REFERENCES workers(id) ON DELETE CASCADE
) ENGINE=INNODB;

-- TEST DATA
INSERT INTO locations (name) VALUES 
('Office A'),
('Warehouse B'),
('Store C'),
('Site D'),
('Branch E');

INSERT INTO tasks (description, location_id, completed) VALUES 
('Clean windows', 1, FALSE),
('Organize inventory', 2, TRUE),
('Restock shelves', 3, FALSE),
('Repair roof', 4, FALSE),
('Update computers', 5, TRUE);

INSERT INTO workers (username, hourly_wage) VALUES 
('john_doe', 15.50),
('jane_smith', 16.75),
('bob_johnson', 14.25),
('alice_brown', 17.00),
('charlie_davis', 15.00);

INSERT INTO logged_time (time_hours, task_id, worker_id) VALUES 
-- Worker 1
(2, 2, 1),
(3, 3, 1),
(1.5, 4, 1),
(4, 5, 1),
-- Worker 2
(2.5, 1, 2),
(1, 3, 2),
(3.5, 4, 2),
(2, 5, 2),
-- Worker 3
(2, 1, 3),
(1.5, 2, 3),
(2.5, 4, 3),
(3, 5, 3),
-- Worker 4
(1, 1, 4),
(2, 2, 4),
(3, 3, 4),
(1.5, 5, 4),
-- Worker 5
(2.5, 1, 5),
(3, 2, 5),
(1.5, 3, 5),
(4, 4, 5);
