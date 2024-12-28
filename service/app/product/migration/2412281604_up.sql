CREATE TABLE  komo_category (
    slug varchar(255) NOT NULL,
    category_name varchar(255) NOT NULL,
    state varchar(16) NOT NULL,
    created_at timestamptz NOT NULL,
    data jsonb NOT NULL,
    PRIMARY KEY (slug)
);