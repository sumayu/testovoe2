создание таблицы для запуска вне докера 
CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
);
NewWalletRepository (хранилище пулов сейчас тут хранится только sql подключение)
to-do не забудь удалить все to-do из проекта 