-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
 id UID,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL,
 patronymic VARCHAR(255),
 CONSTRAINT persons_pk PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS cars (
 regNum VARCHAR(9) NOT NULL,
 mark VARCHAR(255) NOT NULL,
 model VARCHAR(255) NOT NULL,
 owner UID,
 year int,
 CONSTRAINT cars_pk PRIMARY KEY(regNum),
 CONSTRAINT cars_fk1 FOREIGN KEY (owner) REFERENCES persons (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
DROP TABLE persons;
-- +goose StatementEnd
