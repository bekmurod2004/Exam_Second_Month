package postgresql

import (
	"app/api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
	"time"
)

type R_Repo struct {
	db *pgxpool.Pool
}

func NewCodeRepo(db *pgxpool.Pool) *R_Repo {
	return &R_Repo{db: db}
}

func (r *R_Repo) Exam(req *models.StoreChange) (string, error) {
	apiCount, _ := strconv.Atoi(req.Count)
	var (
		give models.ReadFrom
		get  models.ReadTo
	)

	a, _ := r.ReadStocksF(context.Background(), req.Give_StoreId, req.ProdId)
	b, _ := r.ReadStocksG(context.Background(), req.Get_StoreId, req.ProdId)

	for _, v := range a {
		give = v
	}
	for _, v := range b {
		get = v
	}

	fmt.Println()
	fmt.Println("qaysi store dan olinadi = ", give)
	fmt.Println("qaysi store ga qoshiladi = ", get)
	fmt.Println()

	if give.Count-apiCount < 0 {
		fmt.Println("-- cant minus because count few null --")
		return "-- cant minus because count few null --", nil
	}

	// validation , bormi etgan store yomi
	vld, _ := r.Validator(context.Background())
	mapa := make(map[int]int)

	for _, v := range vld {
		mapa[v.Store] = v.Prod

	}

	nextStep := false

	for i, _ := range mapa {
		if i == get.Get_StoreId {
			nextStep = true
			break

		}

	}
	if nextStep == false {
		fmt.Println("stock unaqa store yoq")
		return "stock unaqa store yoq", nil

	}

	if nextStep == true {
		if give.Give_StoreId == get.Get_StoreId {
			fmt.Println("store oziga narsa jonatomidi")
			return "store oziga narsa jonatomidi", nil
		}
		get.Count += apiCount
		give.Count -= apiCount

		r.WriteChanged(context.Background(), give, get)

		fmt.Println("Changed")
		fmt.Println(r.ReadStocksF(context.Background(), req.Give_StoreId, req.ProdId))
		fmt.Println(r.ReadStocksG(context.Background(), req.Get_StoreId, req.ProdId))

	}

	return "all good", nil

}

func (r *R_Repo) ReadStocksF(ctx context.Context, idGive string, prod string) (from []models.ReadFrom, err error) {
	// Give
	rows, err := r.db.Query(ctx, ` SELECT stores.store_name,stocks.store_id,products.product_name,stocks.product_id, stocks.quantity
 FROM stocks
 JOIN products ON stocks.product_id = products.product_id
 JOIN stores ON stocks.store_id = stores.store_id
 WHERE stocks.product_id = $1  AND  stocks.store_id = $2 `, prod, idGive)

	if err != nil {
		fmt.Println("errore")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.ReadFrom
		err = rows.Scan(
			&a.Give_StoreName,
			&a.Give_StoreId,
			&a.ProdName,
			&a.ProdId,
			&a.Count,
		)

		from = append(from, a)
		if err != nil {
			return nil, err
		}

	}

	return from, err

}

func (r *R_Repo) ReadStocksG(ctx context.Context, idTo string, prod string) (to []models.ReadTo, err error) {
	// Get
	rows, err := r.db.Query(ctx, ` SELECT stores.store_name,stocks.store_id,products.product_name,stocks.product_id, stocks.quantity
 FROM stocks
 JOIN products ON stocks.product_id = products.product_id
 JOIN stores ON stocks.store_id = stores.store_id
 WHERE stocks.product_id = $1  AND  stocks.store_id = $2 `, prod, idTo)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.ReadTo
		err = rows.Scan(
			&a.Get_StoreName,
			&a.Get_StoreId,
			&a.ProdName,
			&a.ProdId,
			&a.Count,
		)

		to = append(to, a)
		if err != nil {
			return nil, err
		}

	}

	return to, err

}

func (r *R_Repo) WriteChanged(ctx context.Context, give models.ReadFrom, get models.ReadTo) (err error) {
	// Give logic
	queryGive := `UPDATE stocks SET quantity = $1 WHERE store_id = $2 AND product_id = $3`

	_, err = r.db.Exec(ctx, queryGive,
		give.Count,
		give.Give_StoreId,
		give.ProdId)

	if err != nil {
		return nil
	}

	// Get logic
	queryGet := `UPDATE stocks SET quantity = $1 WHERE store_id = $2 AND product_id = $3`

	_, err = r.db.Exec(ctx, queryGet,
		get.Count,
		get.Get_StoreId,
		get.ProdId)

	if err != nil {
		return nil
	}

	return err
}

func (r *R_Repo) Validator(ctx context.Context) (from []models.Valid, err error) {
	// Give
	rows, err := r.db.Query(ctx, ` SELECT store_id, product_id FROM stocks`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.Valid
		err = rows.Scan(
			&a.Store,
			&a.Prod,
		)

		from = append(from, a)
		if err != nil {
			return nil, err
		}

	}

	return from, err

}

////////////////////////////////////////////////////////// ikkinchi task

func (r R_Repo) GetDate(ctx context.Context, req *models.GiveMe) (res []models.Answer, err error) {
	query := `SELECT
    staffs.first_name || ' ' || staffs.last_name AS "employe",  categories.category_name AS "category",
       products.product_name AS "product",   order_items.quantity AS "quantity",   order_items.list_price * order_items.quantity AS "summ"
FROM orders
         JOIN order_items ON orders.order_id = order_items.order_id
         JOIN products ON order_items.product_id = products.product_id
         JOIN categories ON products.category_id = categories.category_id
         JOIN staffs ON orders.staff_id = staffs.staff_id
WHERE orders.order_date = $1`

	var hh string

	if req.Day == "" {
		dt := time.Now()
		hh = dt.Format("2006-02-01")
	} else {
		hh = req.Day
	}

	fmt.Println("This time -----> ", hh)

	date, error := time.Parse("2006-01-02", req.Day)
	if error != nil {
		fmt.Println(error)
		return
	}

	rows, err := r.db.Query(ctx, query, date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var s models.Answer
		err = rows.Scan(
			&s.StaffName,
			&s.Category,
			&s.Product,
			&s.Quantity,
			&s.Summ,
		)
		res = append(res, s)
		if err != nil {
			return res, err
		}
	}
	return res, nil

}

// ////////////////////////////////////////////////////// Uchinchi task
func (r R_Repo) Create(ctx context.Context, req *models.PromoCreate) (int, error) {
	var id int

	query := `INSERT INTO
    promo(promo_name,is_percent,discount,order_limit_price)
    values($1,$2,$3,$4) RETURNING promo_id`

	fmt.Println(query)

	err := r.db.QueryRow(ctx, query,
		req.PromoName,
		req.IsPercent,
		req.Discount,
		req.Limit_Price,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r R_Repo) GetByID(ctx context.Context, req *models.PromoPrimaryKey) (*models.Promo, error) {
	var (
		query string
		promo models.Promo
	)
	query = `select promo_id , promo_name , is_percent , discount , order_limit_price from promo WHERE promo_id = $1`

	err := r.db.QueryRow(ctx, query, req.Promo_id).Scan(
		&promo.Promo_id,
		&promo.PromoName,
		&promo.IsPercent,
		&promo.Discount,
		&promo.Limit_Price,
	)

	if err != nil {
		return nil, err
	}

	return &promo, nil

}

func (r R_Repo) GetList(ctx context.Context, req *models.Query) (resp []models.Promo, err error) {
	query := `select promo_id , promo_name , is_percent , discount , order_limit_price from promo OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(ctx, query, req.Offset, req.Limit)

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.Promo
		err = rows.Scan(
			&a.Promo_id,
			&a.PromoName,
			&a.IsPercent,
			&a.Discount,
			&a.Limit_Price,
		)

		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r R_Repo) Delete(ctx context.Context, req *models.PromoPrimaryKey) (int64, error) {
	query := `DELETE FROM promo WHERE promo_id = $1`

	result, err := r.db.Exec(ctx, query, req.Promo_id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil

}

////////////////////////////////////////// tortinchi Task

func (r R_Repo) PromoView(ctx context.Context, req *models.StigmaApi) (res models.SigmaSql, err error) {

	query := `select order_id, sum(list_price) AS "list_price" , sum(discount) AS "discount"
from order_items
WHERE order_id = $1 GROUP BY  order_id`

	err = r.db.QueryRow(ctx, query, req.Order_id).Scan(
		&res.Order_id,
		&res.List_price,
		&res.Discount,
	)

	if err != nil {
		return res, err
	}

	if req.Promo_Code == "" {
		return res, nil
	}

	res.List_price -= res.Discount

	return res, nil
}
