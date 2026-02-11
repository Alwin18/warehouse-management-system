# [SETUP-02] Master Data Seeding

**Status**: TODO  
**Priority**: P0 (Critical)  
**Estimated Time**: 3-4 days  
**Assignee**: TBD

## Description
Create database seeder untuk initial master data yang diperlukan sistem bisa berjalan.

## Requirements
- [ ] Create seeder package di `internal/infrastructure/database/seeders/`
- [ ] Seed default roles (Admin, Warehouse Manager, Picker, Receiver)
- [ ] Seed default admin user
- [ ] Seed sample warehouse dengan zones
- [ ] Seed sample locations
- [ ] Seed sample products (10-20 SKU)
- [ ] Seed sample customers & suppliers
- [ ] Create seed command/flag untuk dev environment

## Acceptance Criteria
- ✅ Command `go run cmd/api/main.go --seed` berfungsi
- ✅ Minimal 1 admin user bisa login
- ✅ Minimal 1 warehouse dengan 3 zones tersedia
- ✅ Minimal 10 products dengan UOM tersedia
- ✅ Data seed idempotent (bisa run ulang tanpa duplicate error)

## Technical Notes
```go
// Example seeder structure
package seeders

type Seeder interface {
    Seed(db *gorm.DB) error
}

type RoleSeeder struct{}
func (s *RoleSeeder) Seed(db *gorm.DB) error {
    roles := []models.Role{
        {Code: "ADMIN", Name: "Administrator"},
        {Code: "WH_MGR", Name: "Warehouse Manager"},
        // ...
    }
    return db.FirstOrCreate(&roles).Error
}
```

## Files to Create
- `internal/infrastructure/database/seeders/role_seeder.go`
- `internal/infrastructure/database/seeders/user_seeder.go`
- `internal/infrastructure/database/seeders/warehouse_seeder.go`
- `internal/infrastructure/database/seeders/product_seeder.go`
- `internal/infrastructure/database/seeders/seeder.go` (orchestrator)

## Dependencies
- SETUP-01 (database migration)

## Notes
Gunakan `FirstOrCreate` untuk idempotency. Password default untuk dev: `password123`.
