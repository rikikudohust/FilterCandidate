-- +migrate Up

CREATE TABLE cv (
  id SERIAL PRIMARY KEY,
  name varchar,
  education int,
  language varchar,
  language_point decimal(78,0),
  experience int,
  gender  int
);


CREATE TABLE skill (
  id SERIAL PRIMARY KEY,
  cv_id int REFERENCES cv(id) ON DELETE CASCADE,   
  skill_type_id int REFERENCES cv(id) ON DELETE CASCADE
);

CREATE TABLE personal_skill (
  id SERIAL PRIMARY KEY,
  skill_name varchar
);

-- ALTER TABLE skill FOREIGN KEY (cv_id) REFERENCES cv (id);
-- ALTER TABLE skill FOREIGN KEY (skill_type_id) REFERENCES personal_skill (id);
