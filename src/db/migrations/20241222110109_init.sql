-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    main_category SMALLINT NOT NULL,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(1200)
);
CREATE TABLE category_results (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    category_id UUID NOT NULL,
    result_type SMALLINT NOT NULL,
    result_int SMALLINT,
    result_description VARCHAR(2000),
    date DATE NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
CREATE INDEX idx_category_results_user_category_date
ON category_results (user_id, category_id, date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_category_results_user_category_date;
DROP TABLE category_results;
DROP TABLE categories;
-- +goose StatementEnd
