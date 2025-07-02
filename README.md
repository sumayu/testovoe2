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
