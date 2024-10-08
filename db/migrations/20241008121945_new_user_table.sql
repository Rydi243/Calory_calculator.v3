-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS calorie (
    id            serial primary key,
    product       text     not null,
    weight        int      not null,
    calorie       int      not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS calorie;
-- +goose StatementEnd
