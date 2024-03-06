CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "email" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "date_registered" datetime,
  "last_login" datetime,
  "is_active" bool
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "role_name" varchar,
  "description" varchar
);

CREATE TABLE "user_roles" (
  "id" int PRIMARY KEY,
  "user_id" varchar,
  "role_id" int
);

CREATE TABLE "permissions" (
  "id" int PRIMARY KEY,
  "permission_name" varchar,
  "description" varchar
);

CREATE TABLE "role_permissions" (
  "id" int PRIMARY KEY,
  "role_id" int,
  "permission_id" int
);

CREATE TABLE "audit_logs" (
  "id" int PRIMARY KEY,
  "user_id" varchar,
  "activity_type" varchar,
  "timestamp" datetime,
  "description" varchar
);

CREATE TABLE "password_reset_tokens" (
  "id" int PRIMARY KEY,
  "user_id" varchar,
  "token_value" varchar,
  "expiry_date" datetime
);

CREATE TABLE "two_factor_authentication" (
  "id" int PRIMARY KEY,
  "user_id" varchar,
  "method" varchar,
  "verification_code" varchar,
  "expiry_date" datetime
);

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "password_reset_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "two_factor_authentication" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
