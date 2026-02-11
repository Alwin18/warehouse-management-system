# [MASTER-02] Product Management CRUD

**Status**: TODO  
**Priority**: P1 (High)  
**Estimated Time**: 4-5 days  
**Assignee**: TBD

## Description
Implement full CRUD untuk Product, ProductUOM, dan ProductBatch.

## Requirements
- [ ] Repository layer untuk Product, ProductUOM, ProductBatch
- [ ] Use case dengan validation
- [ ] Handler & routes
- [ ] Search by SKU, name, barcode
- [ ] Filter & pagination
- [ ] Bulk import products via CSV/Excel

## Acceptance Criteria
- ✅ POST `/products` - create product dengan UOMs
- ✅ GET `/products` - list dengan search & filter
- ✅ GET `/products/:id` - detail product dengan UOMs
- ✅ PUT `/products/:id` - update product
- ✅ DELETE `/products/:id` - soft delete
- ✅ POST `/products/:id/uoms` - add UOM conversion
- ✅ GET `/products/search?q=SKU123` - search by SKU/barcode/name
- ✅ POST `/products/import` - bulk import

## API Examples
```bash
# Create product with UOMs
POST /api/v1/products
{
  "sku": "PROD-001",
  "name": "Laptop ASUS ROG",
  "barcode": "8991234567890",
  "base_uom": "PCS",
  "weight": 2.5,
  "is_batch_managed": true,
  "is_serialized": false,
  "uoms": [
    {
      "uom": "BOX",
      "conversion_to_base": 10,
      "is_default_sales": true
    }
  ]
}

# Search products
GET /api/v1/products/search?q=laptop&is_active=true
```

## Files to Create
- `internal/repository/product_repository.go`
- `internal/infrastructure/persistence/product_repository_impl.go`
- `internal/usecase/product_usecase.go`
- `internal/delivery/http/handler/product_handler.go`
- `internal/dto/product.go`

## Dependencies
- SETUP-01, SETUP-02

## Notes
- Base UOM common values: PCS, KG, LITER, EACH
- Batch management: jika true, goods receipt harus specify batch
- Serialized: jika true, track per serial number
