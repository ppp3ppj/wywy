BEGIN;
INSERT INTO "users" ("username", "password", "email", "first_name", "last_name", "date_registered", "last_login", "is_active") VALUES
('admin', '$2a$08$UnB7GAN2UhRifAdlqQwLO.Ll7fWNJL166w9Rb86MhfmrESfYW/0u6', 'admin@wywy.com','admin', 'agent', NOW(), NOW(), true);
-- ('admin', 'password123', 'admin@wywy.com', 'admin', 'agent', '2024-05-15 08:00:00', '2024-05-15 08:30:00', true);
COMMIT;
