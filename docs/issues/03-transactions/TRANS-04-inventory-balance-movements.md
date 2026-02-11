# [TRANS-04] Inventory Balance & Movements

**Status**: TODO  
**Priority**: P2 (Medium)  
**Estimated Time**: 4-5 days  
**Assignee**: TBD

## Description
Implement inventory balance tracking dan movement audit log.

## Requirements
- [ ] Real-time inventory balance updates
- [ ] Movement logging untuk semua transactions
- [ ] Inventory query by product, location, batch
- [ ] Status tracking (AVAILABLE, RESERVED, QC, DAMAGED)
- [ ] Audit trail lengkap

## Acceptance Criteria
- ✅ GET `/inventory/balance` - query balance dengan filter
- ✅ GET `/inventory/movements` - movement history
- ✅ GET `/inventory/balance/:product_id` - balance per product across locations
- ✅ POST `/inventory/adjust` - manual adjustment (link to stock adjustment)
- ✅ Auto-create movement records dari putaway, picking, adjustment
- ✅ Balance updates atomic (transaction-safe)
- ✅ Movement records immutable (append-only)

## API Examples
```bash
# Query inventory balance
GET /api/v1/inventory/balance?warehouse_id=1&product_id=1&status=AVAILABLE

Response:
{
  "data": [
    {
      "warehouse_id": 1,
      "location_id": 10,
      "product_id": 1,
      "batch_id": 5,
      "status": "AVAILABLE",
      "on_hand_qty": 100,
      "reserved_qty": 20,
      "available_qty": 80
    }
  ]
}

# Get movement history
GET /api/v1/inventory/movements?product_id=1&date_from=2025-12-01&date_to=2025-12-31

Response:
{
  "data": [
    {
      "movement_type": "PUTAWAY",
      "product_id": 1,
      "qty": 100,
      "from_location_id": null,
      "to_location_id": 10,
      "reference_type": "PUTAWAY_TASK",
      "reference_id": 5,
      "created_at": "2025-12-12T14:30:00Z",
      "created_by": 2
    }
  ]
}
```

## Files to Create
- `internal/repository/inventory_repository.go`
- `internal/infrastructure/persistence/inventory_repository_impl.go`
- `internal/usecase/inventory_usecase.go`
- `internal/delivery/http/handler/inventory_handler.go`
- `internal/dto/inventory.go`

## Dependencies
- TRANS-02 (Putaway creates movements)
- TRANS-03 (Picking updates balances)

## Business Rules
- Movement types: PUTAWAY, PICK, ADJUSTMENT, TRANSFER, RETURN
- Balance formula: available_qty = on_hand_qty - reserved_qty
- Movement records contain reference_type & reference_id untuk traceability
- Use database transactions untuk ensure consistency
