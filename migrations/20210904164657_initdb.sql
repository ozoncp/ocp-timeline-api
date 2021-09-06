-- +goose Up
-- +goose StatementBegin
CREATE TABLE timeline(
    timeline_id SERIAL PRIMARY KEY,
    user_id INT,
    type_id INT,
    start_timeline TIMESTAMP (0) WITH TIME ZONE,
    end_timeline TIMESTAMP(0) WITH TIME ZONE
);

insert into timeline (user_id, type_id, start, end)
values
    (1,3,now(), now()),
    (2,2,now(), now()),
    (3,6,now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE timeline
-- +goose StatementEnd
