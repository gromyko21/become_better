-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    main_category SMALLINT NOT NULL,
    name VARCHAR(256) NOT NULL,
    progress_type SMALLINT NOT NULL,
    description VARCHAR(1200)
);
CREATE INDEX idx_categories_main_category
ON categories (main_category);

CREATE TABLE progress (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    category_id UUID NOT NULL,
    progress_type SMALLINT NOT NULL,
    result_int SMALLINT,
    result_description VARCHAR(2000),
    date DATE NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
CREATE INDEX idx_progress_user_category_date
ON progress (user_id, category_id, date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_progress_user_category_date;
DROP INDEX idx_categories_main_category;
DROP TABLE progress;
DROP TABLE categories;
-- +goose StatementEnd
