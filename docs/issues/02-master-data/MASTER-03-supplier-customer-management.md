# [MASTER-03] Supplier & Customer Management CRUD

**Status**: TODO  
**Priority**: P1 (High)  
**Estimated Time**: 2-3 days  
**Assignee**: TBD

## Description
Implement CRUD untuk Supplier, Customer, dan Carrier.

## Requirements
- [ ] Repository layer untuk Supplier, Customer, Carrier
- [ ] Use case dengan validation
- [ ] Handler & routes
- [ ] Search & filter
- [ ] Pagination

## Acceptance Criteria
- ✅ POST `/suppliers` - create supplier
- ✅ GET `/suppliers` - list dengan filter
- ✅ GET `/suppliers/:id` - detail supplier
- ✅ PUT `/suppliers/:id` - update supplier
- ✅ DELETE `/suppliers/:id` - soft delete
- ✅ Similar endpoints untuk `/customers` dan `/carriers`
- ✅ Validation: code unique, email format valid

## API Examples
```bash
# Create supplier
POST /api/v1/suppliers
{
  "code": "SUP-001",
  "name": "PT Sumber Makmur",
  "address": "Jl. Raya Industri No. 45",
  "city": "Tangerang",
  "country": "Indonesia",
  "phone": "+62 21 1234567",
  "email": "sales@sumbermakmur.co.id",
  "tax_id": "01.234.567.8-901.000"
}

# List active suppliers
GET /api/v1/suppliers?is_active=true&city=Jakarta
```

## Files to Create
- `internal/repository/supplier_repository.go`
- `internal/repository/customer_repository.go`
- `internal/repository/carrier_repository.go`
- `internal/infrastructure/persistence/*_repository_impl.go`
- `internal/usecase/supplier_usecase.go`
- `internal/usecase/customer_usecase.go`
- `internal/delivery/http/handler/supplier_handler.go`
- `internal/delivery/http/handler/customer_handler.go`
- `internal/dto/supplier.go`, `customer.go`

## Dependencies
- SETUP-01, SETUP-02

## Notes
Pattern sama untuk ketiga entities. Bisa duplicate code dulu, nanti refactor jika ada common pattern.
