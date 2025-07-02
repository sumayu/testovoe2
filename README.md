docker compose up --build 
запросы
post http://localhost:8080/api/v1/wallet
#зачислить деньги

{
  "walletId": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
  "operationType": "DEPOSIT",
  "amount": 500.50
}

#снять деньги
{
  "walletId": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
  "operationType": "WITHDRAW",
  "amount": 200.50
}

get http://localhost:8080/api/v1/wallets/a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11


в configs/postgres/init.sql тут создается бд в докере (если что-то  нужно поменять к примеру название кошелька то сюда)

немного намусорил в коде бд функцией getconfigpath она на основе env (которое я передаю в docker-compose.yml)
она проверяет находится ли сервер в докере или нет и строит пути на основе этого  я пишу это из-за того что можно сразу не понять откуда должно берется isdocker т.к в config.env его нет

создание таблицы для запуска вне докера ЭТО НЕ НУЖНО ЕСЛИ ЗАПУСКАШЬ В ДОКЕРЕ
CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0)
);
INSERT INTO wallets (id) 
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11') 
RETURNING id; 

я удалил все to-do если интересен ход моих мыслей смотри 1 commit там в файле bd и router я все записал
для теста роутера нужно поднять бд (я не стал использовать бд т.к из-за этого пришлось бы переделывать все слои)