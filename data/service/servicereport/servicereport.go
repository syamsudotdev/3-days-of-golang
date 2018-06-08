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
