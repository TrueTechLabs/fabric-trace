-- Active: 1708403864220@@127.0.0.1@3337@fabrictrace
-- 用户ID 用户名 密码 实名信息
CREATE DATABASE fabrictrace
CREATE TABLE users (
    user_id INT PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    realInfo VARCHAR(100) 
);
