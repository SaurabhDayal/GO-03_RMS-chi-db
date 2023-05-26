BEGIN;

INSERT INTO users(user_name, user_password, user_email, credit)
VALUES ('Saurabh Dayal', 'saurabh','saurabh@gmail.com', 5000 ),
       ('Divyank surname', 'divyank','divi@gmail.com', 2500 ),
       ('Vivek surname', 'vivek','vivi@gmail.com', 3500 ),
       ('Manish Singh', 'manish','mani@gmail.com', 1000 ),
       ('Nitin Kumar', 'nitin','niti@gmail.com', 1000 ),
       ('Akash Trivedi', 'akash','aki@gmail.com', 1000 ),
       ('Histesh Kumar', 'hitubaba','hitu@gmail.com', 1000 );

INSERT INTO user_roles(user_id, role)
VALUES (1, 'admin'),
       (2, 'subAdmin'),
       (3, 'subAdmin'),
       (4, 'user'),
       (5, 'user'),
       (6, 'user'),
       (7, 'user');

INSERT INTO addresses(address_name,address_lat, address_lng,user_id)
VALUES ('Noida, Uttar Pradesh', 20.59, 78.12, 1),
       ('Little Noida , Uttar Pradesh', 20.10, 77.02, 2),
       ('New Delhi, NCR', 18.50, 77.19, 3),
       ('Gurugram, Haryana', 10.59, 76.12, 4),
       ('Faridabad, Haryana', 18.90, 75.12, 5),
       ('Ghaziabad, Uttar Pradesh', 19.10, 78.17, 6),
       ('Greater Noida, Uttar Pradesh', 19.50, 76.72, 7);

INSERT INTO addresses(address_name,address_lat, address_lng)
VALUES ('Delhi, Uttar Pradesh', 20.50, 76.72),
       ('New Delhi, Uttar Pradesh', 31.50, 76.72),
       ('Lucknow, Uttar Pradesh', 21.50, 76.72),
       ('Noida, Uttar Pradesh', 19.90, 76.72);

INSERT INTO auths(user_id, user_token)
VALUES (1,'123412341234'),
       (2,'aaaaaaaaaaaa'),
       (3,'bbbbbbbbbbbb'),
       (4,'50b0edb6f1e3'),
       (5,'50b0edb6f1e3');

INSERT INTO restaurants(restaurant_name, restaurant_address, user_id)
VALUES ('Biryani By Kilo',8,2),
       ('China Town',9,2),
       ('KFC',10,3),
       ('Mc Donald',11,3);

INSERT INTO dishes(dish_name, dish_cost, restaurant_id, user_id, preparing_time)
VALUES ('Hakka Noodles', 200, 2, 2,'00:30:00'),
       ('Shezwaan Noodles', 250, 2, 2,'00:50:00'),
       ('Lucknow Biryani', 400, 1, 2,'00:20:00'),
       ('Hyderabad Biryani', 350, 1, 2,'00:30:00'),
       ('12 piece Chicken Bucket', 550, 3, 3,'00:10:00'),
       ('Mc Chicken Burger', 220, 4, 3,'00:30:00');


INSERT INTO orders(dish_id, delivery_time, user_id)
VALUES (1,'00:20:00', 4),
       (2,'00:40:00', 5),
       (3,'01:10:00', 5),
       (4,'01:30:00', 6),
       (5,'00:50:00', 7);

COMMIT;