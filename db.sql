-- Table: public.auth
-- DROP TABLE public.auth;
CREATE TABLE IF NOT EXISTS auth
(
    login character(100) PRIMARY KEY NOT NULL,
    email character(100) NOT NULL UNIQUE,
    password character(100) NOT NULL,
    phone character(100) NOT NULL UNIQUE
);


INSERT INTO auth (login, email, password, phone) VALUES
('user1', 'user1@test.com', 'cGFzc3dvcmQ=', '+79999999999'),
('user2', 'user2@test.com', 'cXdlcnR5', '+78888888888');