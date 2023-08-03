CREATE TABLE specialities (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(100) NOT NULL
);

CREATE TABLE organizations (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(100) NOT NULL,
                               location VARCHAR(255)
);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       role VARCHAR(50),
                       first_name VARCHAR(100),
                       last_name VARCHAR(100),
                       about TEXT,
                       organization_id INTEGER,
                       speciality_id INTEGER,
                       FOREIGN KEY (organization_id) REFERENCES organizations(id),
                       FOREIGN KEY (speciality_id) REFERENCES specialities(id)
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER NOT NULL,
                          title VARCHAR(255) NOT NULL,
                          description TEXT,
                          file_path VARCHAR(255),
                          uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL,
                       post_text TEXT,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          parent_id INTEGER,
                          comment_text TEXT,
                          author_id INTEGER NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          is_review BOOLEAN,
                          product_id INTEGER,
                          grade INTEGER,
                          post_id INTEGER,
                          FOREIGN KEY (parent_id) REFERENCES comments(id),
                          FOREIGN KEY (author_id) REFERENCES users(id),
                          FOREIGN KEY (product_id) REFERENCES products(id),
                          FOREIGN KEY (post_id) REFERENCES posts(id)
);
