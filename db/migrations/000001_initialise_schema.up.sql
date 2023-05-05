CREATE EXTENSION "uuid-ossp";
-- Media table
CREATE TABLE media (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  original_filename VARCHAR(255) NOT NULL,
  original_url VARCHAR(255) NOT NULL,
  thumbnail_url VARCHAR(255) NOT NULL,
  medium_url VARCHAR(255) NOT NULL,
  wide_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- User table
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  avatar_id uuid REFERENCES media(id),
  active BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  role VARCHAR(255) NOT NULL DEFAULT 'user'
);

CREATE INDEX users_email_idx ON users (email);

-- Session table
CREATE TABLE sessions (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email varchar NOT NULL,
  refresh_token varchar NOT NULL,
  user_agent varchar NOT NULL,
  client_ip varchar NOT NULL,
  is_blocked boolean NOT NULL DEFAULT FALSE,
  expires_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);
ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");
CREATE INDEX sessions_refresh_token_idx ON sessions (refresh_token);

-- Password reset table
CREATE TABLE password_resets (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  reset_token VARCHAR(255) NOT NULL,
  expiration TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() + INTERVAL '1 day',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Blog post table
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL UNIQUE,
  content TEXT,
  excerpt TEXT,
  image_id uuid REFERENCES media(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX posts_slug_idx ON posts (slug);

-- Blog post tag table
CREATE TABLE post_tags (
  id SERIAL PRIMARY KEY,
  post_id SERIAL REFERENCES posts(id) ON DELETE CASCADE,
  tag VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


-- Ingredient table
CREATE TABLE ingredients (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX ingredients_name_idx ON ingredients (name);

-- Recipe table
CREATE TABLE recipes (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  directions TEXT,
  author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX recipes_name_idx ON recipes (title);

-- Recipe media relationship table
CREATE TABLE recipe_media (
  recipe_id SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
  media_id uuid NOT NULL REFERENCES media(id) ON DELETE CASCADE,
  PRIMARY KEY (recipe_id, media_id)
);

-- create unit enum in sql
CREATE TYPE ingredient_unit_enum AS ENUM (
  'UNIT_UNSPECIFIED',
  'GRAM',
  'KILOGRAM',
  'MILLILITER',
  'LITER',
  'TEASPOON',
  'TABLESPOON',
  'CUP',
  'PINCH',
  'PIECE'
);

-- Recipe ingredient relationship table
CREATE TABLE recipe_ingredients (
  recipe_id SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
  ingredient_id SERIAL NOT NULL REFERENCES ingredients(id) ON DELETE CASCADE,
  quantity DECIMAL(10, 2) NOT NULL,
  unit ingredient_unit_enum NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  PRIMARY KEY (recipe_id, ingredient_id)
);



CREATE INDEX recipe_ingredients_recipe_id_idx ON recipe_ingredients (recipe_id);
CREATE INDEX recipe_ingredients_ingredient_id_idx ON recipe_ingredients (ingredient_id);
