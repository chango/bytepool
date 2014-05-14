package bytepool

type ArbitraryItem struct {
	pool *ArbitraryPool
	data interface{}
}

func newArbitraryItem(pool *ArbitraryPool, data interface{}) *ArbitraryItem {
	return &ArbitraryItem{
		pool: pool,
		data: data,
	}
}

func (item *ArbitraryItem) Close() error {
	var err error
	if item.pool != nil {
		err = item.pool.resetItem(item.data)
	}
	if err != nil {
		return err
	} else {
		if item.pool != nil {
			item.pool.list <- item
		}
		return nil
	}
}
