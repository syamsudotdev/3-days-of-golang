package servicereport

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"ijahinventory/data/database"
	"ijahinventory/data/model/log"
	"ijahinventory/data/model/product"
	"ijahinventory/util"
)

var (
	db  *gorm.DB
	err error
)

func openDb() {
	db, err = database.OpenDb()
}

func closeDb() {
	db.Close()
}

func GenerateCsvProductCount(path string) error {
	openDb()
	defer closeDb()

	var products []product.Product
	if err := db.Find(&products).Error; err != nil {
		return err
	}

	var data [][]string
	data = append(data, []string{"SKU", "Nama Item", "Jumlah Sekarang"})
	for _, value := range products {
		data = append(
			data,
			[]string{
				value.Sku,
				value.Name,
				strconv.FormatInt(int64(value.StockCount), 10),
			},
		)
	}

	if err := util.WriteToCsv(data, path); err != nil {
		return err
	}

	return nil
}

func GenerateCsvProductIngoing(path string) error {
	openDb()
	defer closeDb()

	var logs []log.LogIngoing
	if err := db.Preload("Product").Find(&logs).Error; err != nil {
		return err
	}

	var data [][]string
	data = append(data,
		[]string{
			"Waktu", "SKU", "Nama Barang",
			"Jumlah Pemesanan", "Jumlah Diterima", "Harga Beli",
			"Total", "Nomor Kwitansi", "Catatan"})
	for _, value := range logs {
		data = append(
			data,
			[]string{
				value.Timestamp.Format(time.UnixDate),
				value.Product.Sku,
				value.Product.Name,
				util.IntToString(value.CountOrder),
				util.IntToString(value.CountReceived),
				util.IntToString(value.BuyPrice),
				util.IntToString(value.TotalPrice),
				value.ReceiptNumber,
				value.Note,
			},
		)
	}

	if err := util.WriteToCsv(data, path); err != nil {
		return err
	}

	return nil
}

func GenerateCsvProductOutgoing(path string) error {
	openDb()
	defer closeDb()

	var logs []log.LogOutgoing
	if err := db.Preload("Product").Find(&logs).Error; err != nil {
		return err
	}

	var data [][]string
	data = append(data, []string{
		"Waktu", "SKU", "Nama Barang", "Jumlah Keluar",
		"Harga Jual", "Total", "Catatan",
	})
	for _, value := range logs {
		data = append(
			data,
			[]string{
				value.Timestamp.Format(time.UnixDate),
				value.Product.Sku,
				value.Product.Name,
				util.IntToString(value.CountOutgoing),
				util.IntToString(value.SalePrice),
				util.IntToString(value.TotalPrice),
				value.Note,
			},
		)
	}

	if err := util.WriteToCsv(data, path); err != nil {
		return err
	}

	return nil
}

func GenerateCsvProductValue(path string) error {
	openDb()
	defer closeDb()

	var countSku int
	if err := db.
		Model(&product.Product{}).
		Count(&countSku).Error; err != nil {
		return err
	}

	var countProductStock int
	row := db.Raw("SELECT SUM(stock_count) from products").Row()
	row.Scan(&countProductStock)

	var data [][]string
	data = append(data, []string{"Laporan Nilai Barang"})
	data = append(data, []string{" "})
	data = append(data, []string{
		"Tanggal Cetak",
		time.Now().Format(time.UnixDate),
	})
	data = append(data, []string{
		"Jumlah SKU",
		util.IntToString(countSku),
	})
	//total quantity in inventory
	data = append(data, []string{
		"Jumlah Total Barang",
		util.IntToString(countProductStock),
	})

	rows, err := db.
		Raw("SELECT " +
			" p.sku, p.name, p.stock_count," +
			" aa.qty, aa.avgbuy, (aa.qty * aa.avgbuy) AS total " +
			" FROM products p " +
			" LEFT JOIN ( " +
			"SELECT " +
			" l.product_id, SUM(l.count_received) AS qty, " +
			" (sum(l.total_price)/sum(l.count_received)) as avgbuy " +
			"FROM log_ingoing l " +
			"GROUP BY l.product_id) aa " +
			" ON p.id = aa.product_id").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	var items [][]string
	var totalValue int
	for rows.Next() {
		var sku, productName, stockCount string
		var quantity, avgBuy, total int
		rows.Scan(&sku, &productName, &stockCount, &quantity, &avgBuy, &total)
		totalValue += total
		items = append(
			items,
			[]string{
				sku,
				productName,
				stockCount,
				util.IntToString(quantity),
				util.IntToString(avgBuy),
				util.IntToString(total)},
		)
	}

	data = append(data, []string{"Total Nilai", util.IntToString(totalValue)})
	data = append(data, []string{" "})
	data = append(data, []string{"SKU", "Nama Item", "Jumlah",
		"Rata-Rata Harga Beli", "Total"})
	data = append(data, items...)

	if err := util.WriteToCsv(data, path); err != nil {
		return err
	}

	return nil
}

func GenerateCsvSales(path string, start time.Time, end time.Time) error {
	openDb()
	defer closeDb()

	layout := "2006-01-02 15:04:05"
	startString := start.Format(layout)
	end = end.AddDate(0, 0, 1)
	endString := end.Format(layout)

	var data [][]string
	data = append(data, []string{"Laporan Penjualan"})
	data = append(data, []string{" "})
	data = append(data, []string{
		"Tanggal Cetak",
		time.Now().Format(time.UnixDate),
	})
	data = append(data, []string{"Tanggal", startString + " - " + endString})

	rows, err := db.
		Raw("SELECT " +
			"l.note, l.timestamp, p.sku, p.name, l.count_outgoing, " +
			"l.sale_price, l.total_price, aa.avgbuy, " +
			"(l.total_price - aa.avgbuy) AS profit " +
			"FROM log_outgoing l " +
			"JOIN (" +
			"SELECT " +
			"l.product_id, SUM(l.count_received) AS qty, " +
			"(SUM(l.total_price) / SUM(l.count_received)) AS avgbuy " +
			"FROM log_ingoing l GROUP BY l.product_id) aa " +
			"ON l.product_id = aa.product_id " +
			"JOIN products p ON l.product_id = p.id " +
			"WHERE l.timestamp " +
			"BETWEEN datetime('" + startString + "') " +
			"AND datetime('" + endString + "')").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	var items [][]string
	var totalTurnover, totalProfit, totalQuantity int
	for rows.Next() {
		var note, timestamp, sku, productName string
		var quantity, salePrice, totalPrice, avgBuy, profit int
		rows.Scan(&note, &timestamp, &sku, &productName, &quantity, &salePrice,
			&totalPrice, &avgBuy, &profit)
		totalTurnover += totalPrice
		totalProfit += profit
		totalQuantity += quantity
		items = append(
			items,
			[]string{
				note,
				timestamp,
				sku, productName,
				util.IntToString(quantity),
				util.IntToString(salePrice),
				util.IntToString(totalPrice),
				util.IntToString(avgBuy),
				util.IntToString(profit)},
		)
	}

	data = append(
		data,
		[]string{"Total Omzet", util.IntToString(totalTurnover)},
	)

	data = append(
		data,
		[]string{"Total Laba Kotor", util.IntToString(totalProfit)},
	)

	data = append(
		data,
		[]string{"Total Penjualan", util.IntToString(len(items))},
	)

	data = append(
		data,
		[]string{"Total Barang", util.IntToString(totalQuantity)},
	)

	data = append(data, []string{" "})
	data = append(data, []string{"ID-Pesanan", "Waktu", "SKU", "Nama Barang",
		"Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})
	data = append(data, items...)

	if err := util.WriteToCsv(data, path); err != nil {
		return err
	}

	return nil
}
