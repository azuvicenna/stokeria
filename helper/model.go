package helper

import (
	"stokeria/model/domain"
	"stokeria/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func ToSupplierResponse(supplier domain.Supplier) web.SupplierResponse {
	return web.SupplierResponse{
		ID:            supplier.ID,
		Name:          supplier.Name,
		ContactPerson: supplier.ContactPerson,
		Phone:         supplier.Phone,
		Address:       supplier.Address,
		CreatedAt:     supplier.CreatedAt,
		UpdatedAt:     supplier.UpdatedAt,
	}
}

func ToWarehouseResponse(warehouse domain.Warehouse) web.WarehouseResponse {
	return web.WarehouseResponse{
		ID:        warehouse.ID,
		Code:      warehouse.Code,
		Name:      warehouse.Name,
		Address:   warehouse.Address,
		CreatedAt: warehouse.CreatedAt,
		UpdatedAt: warehouse.UpdatedAt,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	response := web.ProductResponse{
		ID:        product.ID,
		SKU:       product.SKU,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	// Menangani relasi: Jika GORM melakukan Preload Category, maka ikut dikonversi
	if product.Category.ID != "" {
		categoryResponse := ToCategoryResponse(product.Category)
		response.Category = &categoryResponse
	}

	return response
}

// Helper untuk mengonversi slice/array dari Product (Berguna untuk API List/Get All)
func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToInventoryResponse(inventory domain.Inventory) web.InventoryResponse {
	response := web.InventoryResponse{
		ID:        inventory.ID,
		Quantity:  inventory.Quantity,
		UpdatedAt: inventory.UpdatedAt,
	}

	if inventory.Product.ID != "" {
		productResp := ToProductResponse(inventory.Product)
		response.Product = &productResp
	}
	if inventory.Warehouse.ID != "" {
		warehouseResp := ToWarehouseResponse(inventory.Warehouse)
		response.Warehouse = &warehouseResp
	}

	return response
}

func ToTransactionDetailResponse(detail domain.TransactionDetail) web.TransactionDetailResponse {
	response := web.TransactionDetailResponse{
		ID:        detail.ID,
		Quantity:  detail.Quantity,
		UnitPrice: detail.UnitPrice,
	}

	if detail.Product.ID != "" {
		productResp := ToProductResponse(detail.Product)
		response.Product = &productResp
	}

	return response
}

func ToTransactionDetailResponses(details []domain.TransactionDetail) []web.TransactionDetailResponse {
	var responses []web.TransactionDetailResponse
	for _, detail := range details {
		responses = append(responses, ToTransactionDetailResponse(detail))
	}
	return responses
}

func ToInboundTransactionResponse(trx domain.InboundTransaction, details []domain.TransactionDetail) web.InboundTransactionResponse {
	response := web.InboundTransactionResponse{
		ID:              trx.ID,
		ReferenceCode:   trx.ReferenceCode,
		TransactionDate: trx.TransactionDate,
		Status:          trx.Status,
		CreatedAt:       trx.CreatedAt,
		UpdatedAt:       trx.UpdatedAt,
		Details:         ToTransactionDetailResponses(details),
	}

	if trx.Supplier.ID != "" {
		supplierResp := ToSupplierResponse(trx.Supplier)
		response.Supplier = &supplierResp
	}
	if trx.User.ID != "" {
		userResp := ToUserResponse(trx.User)
		response.User = &userResp
	}

	return response
}

func ToOutboundTransactionResponse(trx domain.OutboundTransaction, details []domain.TransactionDetail) web.OutboundTransactionResponse {
	response := web.OutboundTransactionResponse{
		ID:              trx.ID,
		ReferenceCode:   trx.ReferenceCode,
		Destination:     trx.Destination,
		TransactionDate: trx.TransactionDate,
		Status:          trx.Status,
		CreatedAt:       trx.CreatedAt,
		UpdatedAt:       trx.UpdatedAt,
		Details:         ToTransactionDetailResponses(details),
	}

	if trx.User.ID != "" {
		userResp := ToUserResponse(trx.User)
		response.User = &userResp
	}

	return response
}

func ToStockAdjustmentResponse(adj domain.StockAdjustment) web.StockAdjustmentResponse {
	response := web.StockAdjustmentResponse{
		ID:          adj.ID,
		OldQuantity: adj.OldQuantity,
		NewQuantity: adj.NewQuantity,
		Reason:      adj.Reason,
		CreatedAt:   adj.CreatedAt,
	}

	if adj.Product.ID != "" {
		productResp := ToProductResponse(adj.Product)
		response.Product = &productResp
	}
	if adj.Warehouse.ID != "" {
		warehouseResp := ToWarehouseResponse(adj.Warehouse)
		response.Warehouse = &warehouseResp
	}
	if adj.User.ID != "" {
		userResp := ToUserResponse(adj.User)
		response.User = &userResp
	}

	return response
}