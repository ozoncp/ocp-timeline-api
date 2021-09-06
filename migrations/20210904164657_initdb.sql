-- +goose Up
-- +goose StatementBegin
CREATE TABLE timeline(
    timeline_id SERIAL PRIMARY KEY,
    user_id INT,
    "type_id" INT,
    "from" TIMESTAMP (0) WITH TIME ZONE,
    "to" TIMESTAMP(0) WITH TIME ZONE
);

insert into timeline (user_id, "type_id", "from", "to")
values
    (1,3,now(), now()),
    (2,2,now(), now()),
    (3,6,now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE timeline
-- +goose StatementEnd
