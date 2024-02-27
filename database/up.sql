DROP TABLE IF EXISTS students;

CREATE TABLE students (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
);

CREATE TABLE tests (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);

CREATE TABLE questions (
    id VARCHAR(32) PRIMARY KEY,
    test_id VARCHAR(32) NOT NULL,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(255) NOT NULL,
    FOREIGN KEY (test_id) REFERENCES tests(id)
);
