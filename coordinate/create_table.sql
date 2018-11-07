CREATE TABLE coordinates (
  id serial primary key,
  user_id INTEGER not null
);

CREATE TABLE coordinate_items (
  coordinate_id INTEGER not null references coordinates(id),
  item_id INTEGER not null,
  primary key (coordinate_id, item_id)
);
