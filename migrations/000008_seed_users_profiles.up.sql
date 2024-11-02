INSERT INTO users (id, name, email, password)
VALUES
    ('b36c7205-e063-4d56-8c41-e7a504e0e0c4', 'Alice Smith', 'alice@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('c1a2b2b1-9811-45c7-9a1b-913cba50d1c0', 'Bob Johnson', 'bob@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('4f52456e-bb8b-45a1-9f58-ff8e4b2d87bc', 'Charlie Brown', 'charlie@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('02f3eafe-3e2e-4487-9447-05480ef63c7e', 'David Wilson', 'david@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('f64190d2-9b8f-4732-a50d-e0032898db26', 'Eva Green', 'eva@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('d8e0cbfa-9b85-4bc4-8485-4a174e8ee9d4', 'Frank White', 'frank@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('6b6cb630-abb8-4711-a8d8-fb42f0f6b8e5', 'Grace Black', 'grace@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('f4e9c7c2-909f-418c-bfbc-2a824bf5939e', 'Hank Grey', 'hank@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('30c5e2f7-8148-470b-9d8d-165fd4e689e7', 'Ivy Adams', 'ivy@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('e9dff5e4-6f1f-473c-b12b-e51f36cfd76e', 'Jack Blue', 'jack@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('1e3c6e08-603c-4b6b-9625-5a9dbf6e236f', 'Karen Pink', 'karen@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('d7c9cc54-38f1-4a4d-8c8f-53cf1f635261', 'Leo Silver', 'leo@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('b5b1e6ec-f7a6-409c-814b-5a8e586c57f9', 'Mia Orange', 'mia@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('c67e3e0b-8b1d-4972-a0a1-e115eacbf2a8', 'Nina Violet', 'nina@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('c1a2b8a1-3124-42bb-81cf-dc80fbb1ae59', 'Oscar Gold', 'oscar@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('2e46d9b4-e74e-4cc7-bf77-dbb19a65d5da', 'Paul Copper', 'paul@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('12e69d77-62b8-45c9-843f-3f1b4c0cd4c9', 'Quinn Steel', 'quinn@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq'),
    ('3f21b74f-6727-42ec-b8b9-23867ef8c314', 'Rose Iron', 'rose@example.com', '$2a$10$nosfFOy/i4OTtTVwn6eaQed2kTmvQSny2JFZdfhG0gfuTDLZKZIxq');

INSERT INTO profiles (id, user_id, bio, age, location, profile_pic_url)
VALUES
    ('6048ce30-d213-47dd-a2fe-f91b262650fc', 'b36c7205-e063-4d56-8c41-e7a504e0e0c4', 'I am a software developer.', 28, 'Bandung', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('1b5ba771-4ddd-4b35-b32f-5386ed57c230', 'c1a2b2b1-9811-45c7-9a1b-913cba50d1c0', 'Lover of music and art.', 32, 'Jakarta', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('d0216eae-a6f4-4a87-a1ed-cdb053ac215d', '4f52456e-bb8b-45a1-9f58-ff8e4b2d87bc', 'Enjoys hiking and outdoor activities.', 25, 'Surabaya', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('95ce4e9b-c69d-4b62-8c7f-8e2c13d5ec46', '02f3eafe-3e2e-4487-9447-05480ef63c7e', 'Passionate about technology.', 30, 'Yogyakarta', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('0a57efc0-18d1-49f7-9e14-f720dfa69ddd', 'f64190d2-9b8f-4732-a50d-e0032898db26', 'Coffee enthusiast.', 27, 'Bandung', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('b0927586-4151-4b7d-b8af-b3f68b71fb08', 'd8e0cbfa-9b85-4bc4-8485-4a174e8ee9d4', 'Traveler and foodie.', 35, 'Bali', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('c7e2d47c-2e94-49ca-8977-e4a7f5aa9fa2', '6b6cb630-abb8-4711-a8d8-fb42f0f6b8e5', 'Avid reader and writer.', 29, 'Medan', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('dbbab3e0-8207-4054-b004-9f93c715f5a1', 'f4e9c7c2-909f-418c-bfbc-2a824bf5939e', 'Nature lover.', 34, 'Malang', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('faabeee0-89c3-4113-873b-379fa1623ab0', '30c5e2f7-8148-470b-9d8d-165fd4e689e7', 'Fitness freak.', 22, 'Banjarmasin', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('91e4267c-58a7-4bdf-94c5-91a7ddfdd02c', 'e9dff5e4-6f1f-473c-b12b-e51f36cfd76e', 'Music producer.', 31, 'Semarang', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('00d11b9b-d8be-4e6e-99cc-6aaf856762e4', '1e3c6e08-603c-4b6b-9625-5a9dbf6e236f', 'Gamer and streamer.', 24, 'Jakarta', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('c08e5f90-7a1d-486d-a412-af5443a9df68', 'd7c9cc54-38f1-4a4d-8c8f-53cf1f635261', 'Food blogger.', 26, 'Bandung', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('ee156e6b-fb67-49f9-98d6-99a8861cd1c0', 'b5b1e6ec-f7a6-409c-814b-5a8e586c57f9', 'Photographer.', 33, 'Bali', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('633164b0-1b8d-486a-b5a6-7e929c1bcab3', 'c67e3e0b-8b1d-4972-a0a1-e115eacbf2a8', 'Motivational speaker.', 37, 'Yogyakarta', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('d2a24204-ae00-41e7-9d82-43945bfd10b3', 'c1a2b8a1-3124-42bb-81cf-dc80fbb1ae59', 'Online entrepreneur.', 29, 'Surabaya', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('0e25b958-e402-477f-9561-c88827184f4d', '2e46d9b4-e74e-4cc7-bf77-dbb19a65d5da', 'Sports enthusiast.', 36, 'Medan', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('1393dd89-f38a-4522-bfc3-52089770cfe6', '12e69d77-62b8-45c9-843f-3f1b4c0cd4c9', 'Fashion designer.', 30, 'Malang', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png'),
    ('a34c4c76-35e8-4eb3-9c47-35aa78ee3166', '3f21b74f-6727-42ec-b8b9-23867ef8c314', 'Health coach.', 34, 'Banjarmasin', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ac/No_image_available.svg/800px-No_image_available.svg.png');