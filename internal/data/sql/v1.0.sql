CREATE TABLE users (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    name varchar(255) NOT NULL DEFAULT '',
    age integer,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN users.id IS 'id';
COMMENT ON COLUMN users.name IS '姓名';
COMMENT ON COLUMN users.age IS '年龄';
COMMENT ON COLUMN users.created_at IS '创建时间';
COMMENT ON COLUMN users.updated_at IS '更新时间';