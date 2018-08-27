CREATE TABLE coordinates (
  id serial primary key,
  user_id INTEGER not null,
  item_ids VARCHAR(32) not null
);
