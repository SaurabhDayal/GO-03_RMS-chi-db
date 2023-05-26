BEGIN;

Create table if not exists users (
    id Serial primary key,
    user_name TEXT NOT NULL,
    user_password TEXT Not null,
    user_email TEXT NOT NULL,
    credit INT not null,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_email ON users(user_email) WHERE archived_at IS NULL;

CREATE TYPE role_type AS ENUM (
    'admin',
    'subAdmin',
    'user'
);

Create table if not exists user_roles (
    id Serial primary key,
    user_id INT not null,
    role role_type not null,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_role ON user_roles(user_id, role) WHERE archived_at IS NULL;

Create table if not exists auths (
    user_id INT not null,
    user_token TEXT not null,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);

Create table if not exists addresses (
    id Serial primary key,
    address_name TEXT not null,
    address_lat NUMERIC(4,2) not null,
    address_lng NUMERIC(4,2) not null,
    user_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);

Create table if not exists restaurants (
    id Serial primary key,
    restaurant_name TEXT not null,
    restaurant_address INT not null,
    user_id INT not null,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id),
    CONSTRAINT fk_address
        FOREIGN KEY(restaurant_address)
            REFERENCES addresses(id)
);

Create table if not exists dishes (
    id Serial primary key,
    dish_name TEXT not null,
    dish_cost Int not null,
    restaurant_id INT not null,
    user_id INT not null,
    preparing_time TIME NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id),
    CONSTRAINT fk_restaurant
        FOREIGN KEY(restaurant_id)
            REFERENCES restaurants(id)
);

Create table if not exists orders (
    id Serial primary key,
    dish_id INT not null,
    delivery_time TIME NOT NULL,
    user_id INT not null,
    is_delivered BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    archived_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id),
    CONSTRAINT fk_dish
        FOREIGN KEY(dish_id)
            REFERENCES dishes(id)
);

COMMIT;