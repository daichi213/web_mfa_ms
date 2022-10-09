CREATE ROLE root LOGIN PASSWORD 'password';
CREATE DATABASE schedule_test;
GRANT ALL PRIVILEGES ON DATABASE schedule_test TO root;
\c root;

-- CREATE DATABASE IF NOT EXISTS schedule_test;
-- -- ALTER DATABASE schedule_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- CREATE DATABASE IF NOT EXISTS schedule_production;
-- -- ALTER DATABASE schedule_production CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;9