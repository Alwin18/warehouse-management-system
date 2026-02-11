# [SETUP-01] Database Migration & Initial Setup

**Status**: TODO  
**Priority**: P0 (Critical)  
**Estimated Time**: 2-3 days  
**Assignee**: Backend Lead

## Description
Setup database migration system dan pastikan semua 45 tables ter-create dengan benar di PostgreSQL.

## Requirements
- [ ] Verify PostgreSQL connection ke database
- [ ] Run auto migrate untuk semua models
- [ ] Verify semua tables, indexes, dan foreign keys ter-create
- [ ] Create migration rollback strategy
- [ ] Document database schema di ERD

## Acceptance Criteria
- ✅ Semua 45 tables ada di database
- ✅ Foreign key constraints berfungsi
- ✅ Unique indexes ter-create dengan benar
- ✅ No migration errors di logs
- ✅ ERD diagram tersedia di dokumentasi

## Technical Notes
```bash
# Run migration
go run cmd/api/main.go

# Verify tables
psql -U postgres -d warehouse_hub -c "\dt"

# Check foreign keys
SELECT constraint_name, table_name 
FROM information_schema.table_constraints 
WHERE constraint_type = 'FOREIGN KEY';
```

## Files to Modify
- `internal/infrastructure/database/postgres.go` (already done)
- `cmd/api/main.go` (verify migration call)

## Dependencies
- None (first task)

## Notes
Pastikan semua GORM tags di models sudah benar sebelum run migration.
