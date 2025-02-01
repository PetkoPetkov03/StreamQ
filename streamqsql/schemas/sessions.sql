CREATE TABLE IF NOT EXISTS USERSESSIONS (
    ID INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    TOKEN CLOB NOT NULL UNIQUE,
    CREATED_FROM CLOB NOT NULL UNIQUE,
    CREATED_AT DATE NOT NULL DEFAULT now,
    EXPIRE_DATE DATE NOT NULL DEFAULT now,
    USER_ID INTEGER NOT NULL UNIQUE,
    FOREIGN KEY(USER_ID) REFERENCES USERS(id)
);