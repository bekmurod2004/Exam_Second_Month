# Exam_Second_Month

1. Task 1  ---> Post | http://localhost:4001/exam ,
 Json | { "give_store_id": "3", "get_store_id": "1", "prod_id": "13", "count" : "1" }

2. Task 2  ---> Get /view/2016-06-01

3. Task 3  ---> Post  /promo ,
{ "promo_name" : "SoloShop", "is_percent": true, "discount": 10, "order_limit_price": 100 },
 GetById | /promo/:id ,
 GetList | /promo, Delete | /promo/:id
 Delete | /promo/:id

4.Task 4 ---> Post /order_discount , { "order_id" : "32", "promo_code": "JUBAJUBA" },
promocode sanamsli uchun "promo_code": "" berish kere

5. Task 5 ---> Post /order_item logika qoshdim , 
{ "discount": 10, "list_price": 11000, "order_id": 999, "product_id": 2, "quantity": 1 }