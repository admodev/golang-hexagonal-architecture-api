-- +goose Up
CREATE TABLE IF NOT EXISTS reports (
    id int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    start_date timestamp NOT NULL,
    end_date timestamp NOT NULL,
    job_address varchar(100) NOT NULL,
    code_charge varchar(50) NOT NULL,
    quantity int(10) NOT NULL,
    units_description varchar(150) NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE reports;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
