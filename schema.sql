CREATE TABLE students (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  clas text,
  scool text,
  order_day SMALLINT,
  order_time TIME,
  order_cost SMALLINT
);