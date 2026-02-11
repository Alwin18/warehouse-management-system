# [MASTER-01] Warehouse Management CRUD

**Status**: TODO  
**Priority**: P1 (High)  
**Estimated Time**: 3-4 days  
**Assignee**: TBD

## Description
Implement full CRUD untuk Warehouse, WarehouseZone, dan Location.

## Requirements
- [ ] Repository layer untuk Warehouse, WarehouseZone, Location
- [ ] Use case layer dengan business logic
- [ ] Handler & routes dengan proper validation
- [ ] Filter & pagination untuk list endpoints
- [ ] Soft delete untuk warehouses (ubah is_active)

## Acceptance Criteria
- ✅ POST `/warehouses` - create warehouse
- ✅ GET `/warehouses` - list dengan filter & pagination
- ✅ GET `/warehouses/:id` - detail warehouse dengan zones & locations
- ✅ PUT `/warehouses/:id` - update warehouse
- ✅ DELETE `/warehouses/:id` - soft delete (set is_active = false)
- ✅ POST `/warehouses/:id/zones` - create zone di warehouse
- ✅ POST `/warehouses/:id/locations` - create location di zone
- ✅ Validation: code unique, required fields

## API Examples
```bash
# Create warehouse
POST /api/v1/warehouses
{
  "code": "WH-JKT-01",
  "name": "Jakarta Main Warehouse",
  "address": "Jl. Sudirman No. 123",
  "city": "Jakarta",
  "country": "Indonesia",
  "time_zone": "Asia/Jakarta"
}

# List warehouses with filter
GET /api/v1/warehouses?is_active=true&city=Jakarta&page=1&limit=20

# Get warehouse detail with zones
GET /api/v1/warehouses/1?include=zones,locations
```

## Files to Create
- `internal/repository/warehouse_repository.go`
- `internal/infrastructure/persistence/warehouse_repository_impl.go`
- `internal/usecase/warehouse_usecase.go`
- `internal/delivery/http/handler/warehouse_handler.go`
- `internal/dto/warehouse.go`

## Dependencies
- SETUP-01, SETUP-02

## Notes
Zone types: RECEIVING, STORAGE, PICKING, PACKING, SHIPPING, QC
Location types: RACK, BIN, FLOOR, STAGING
