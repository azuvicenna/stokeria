package app

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"stokeria/model/domain"
)

func RunSeeder(db *gorm.DB) error {
	now := time.Now()
	var err error

	var userIDs, categoryIDs, supplierIDs, warehouseIDs, productIDs []string
	var inboundIDs, outboundIDs []string

	// 1. Users
	var users []domain.User
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		userIDs = append(userIDs, id)
		users = append(users, domain.User{ID: id, Email: fmt.Sprintf("user%d@stokeria.com", i), PasswordHash: "$2a$10$7Z01yDkWwP14/dK2eQ9yEOm5H22b40J1Yq.uS3rO40Q35hJz19yC2", Name: fmt.Sprintf("User %d", i), Role: "admin"})
	}
	if err = db.Create(&users).Error; err != nil {
		return err
	}

	// 2. Categories
	var categories []domain.Category
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		categoryIDs = append(categoryIDs, id)
		categories = append(categories, domain.Category{ID: id, Name: fmt.Sprintf("Kategori %d", i), Description: "Deskripsi Kategori"})
	}
	if err = db.Create(&categories).Error; err != nil {
		return err
	}

	// 3. Suppliers
	var suppliers []domain.Supplier
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		supplierIDs = append(supplierIDs, id)
		suppliers = append(suppliers, domain.Supplier{ID: id, Name: fmt.Sprintf("Supplier PT %d", i), ContactPerson: fmt.Sprintf("Kontak %d", i), Phone: "08123456789", Address: "Alamat Supplier"})
	}
	if err = db.Create(&suppliers).Error; err != nil {
		return err
	}

	// 4. Warehouses
	var warehouses []domain.Warehouse
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		warehouseIDs = append(warehouseIDs, id)
		warehouses = append(warehouses, domain.Warehouse{ID: id, Code: fmt.Sprintf("WH-%03d", i), Name: fmt.Sprintf("Gudang %d", i), Address: "Alamat Gudang"})
	}
	if err = db.Create(&warehouses).Error; err != nil {
		return err
	}

	// 5. Products
	var products []domain.Product
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		productIDs = append(productIDs, id)
		products = append(products, domain.Product{ID: id, SKU: fmt.Sprintf("SKU-%04d", i), CategoryID: categoryIDs[i-1], Name: fmt.Sprintf("Produk %d", i), Price: 50000.0})
	}
	if err = db.Create(&products).Error; err != nil {
		return err
	}

	// 6. Inventory
	var inventories []domain.Inventory
	for i := 1; i <= 5; i++ {
		inventories = append(inventories, domain.Inventory{ID: uuid.New().String(), ProductID: productIDs[i-1], WarehouseID: warehouseIDs[i-1], Quantity: 100})
	}
	if err = db.Create(&inventories).Error; err != nil {
		return err
	}

	// 7. Inbound Transactions
	var inbounds []domain.InboundTransaction
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		inboundIDs = append(inboundIDs, id)
		inbounds = append(inbounds, domain.InboundTransaction{
			ID:              id,
			ReferenceCode:   fmt.Sprintf("INB-%04d", i),
			SupplierID:      supplierIDs[i-1],
			WarehouseID:     warehouseIDs[i-1],
			UserID:          userIDs[i-1],
			TransactionDate: now,
			Status:          "COMPLETED",
		})
	}
	if err = db.Create(&inbounds).Error; err != nil {
		return err
	}

	// 8. Outbound Transactions
	var outbounds []domain.OutboundTransaction
	for i := 1; i <= 5; i++ {
		id := uuid.New().String()
		outboundIDs = append(outboundIDs, id)
		outbounds = append(outbounds, domain.OutboundTransaction{
			ID:              id,
			ReferenceCode:   fmt.Sprintf("OUT-%04d", i),
			WarehouseID:     warehouseIDs[i-1],
			Destination:     "Tujuan Pengiriman",
			UserID:          userIDs[i-1],
			TransactionDate: now,
			Status:          "COMPLETED",
		})
	}
	if err = db.Create(&outbounds).Error; err != nil {
		return err
	}

	// 9. Transaction Details
	var details []domain.TransactionDetail
	for i := 1; i <= 5; i++ {
		transactionType := "INBOUND"
		transactionID := inboundIDs[i-1]
		if i > 3 {
			transactionType = "OUTBOUND"
			transactionID = outboundIDs[i-1]
		}
		details = append(details, domain.TransactionDetail{
			ID:              uuid.New().String(),
			TransactionID:   transactionID,
			TransactionType: transactionType,
			ProductID:       productIDs[i-1],
			Quantity:        10,
			UnitPrice:       50000.0,
		})
	}
	if err = db.Create(&details).Error; err != nil {
		return err
	}

	// 10. Stock Adjustments
	var adjustments []domain.StockAdjustment
	for i := 1; i <= 5; i++ {
		adjustments = append(adjustments, domain.StockAdjustment{
			ID:          uuid.New().String(),
			ProductID:   productIDs[i-1],
			WarehouseID: warehouseIDs[i-1],
			UserID:      userIDs[i-1],
			OldQuantity: 100,
			NewQuantity: 95,
			Reason:      "Barang rusak",
		})
	}
	if err = db.Create(&adjustments).Error; err != nil {
		return err
	}

	return nil
}
