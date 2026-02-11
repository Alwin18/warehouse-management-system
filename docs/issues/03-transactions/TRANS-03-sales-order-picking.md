# [TRANS-03] Sales Order & Picking Flow

**Status**: TODO  
**Priority**: P2 (Medium)  
**Estimated Time**: 7-8 days  
**Assignee**: TBD

## Description
Implement Sales Order creation, allocation, dan Picking task generation.

## Requirements
- [ ] Create SO dengan multiple lines
- [ ] SO status workflow: DRAFT → SUBMITTED → ALLOCATED → PICKING → PACKED → SHIPPED
- [ ] Inventory allocation (reserve stock)
- [ ] Generate picking waves & tasks
- [ ] Assign tasks to pickers
- [ ] Record picked quantities
- [ ] Update inventory movements

## Acceptance Criteria
- ✅ POST `/sales-orders` - create SO
- ✅ PUT `/sales-orders/:id/submit` - submit SO
- ✅ PUT `/sales-orders/:id/allocate` - allocate inventory
- ✅ POST `/picking-waves` - create wave with multiple SOs
- ✅ POST `/picking-tasks` - generate picking tasks
- ✅ PUT `/picking-tasks/:id/assign` - assign to picker
- ✅ PUT `/picking-tasks/:id/pick` - record picked items
- ✅ PUT `/picking-tasks/:id/complete` - complete picking
- ✅ Inventory reserved_qty updated on allocation
- ✅ Inventory on_hand_qty decreased on pick completion

## API Examples
```bash
# Create SO
POST /api/v1/sales-orders
{
  "customer_id": 1,
  "warehouse_id": 1,
  "order_date": "2025-12-12",
  "requested_ship_date": "2025-12-15",
  "priority": "HIGH",
  "shipping_address": "Jl. Customer No. 123",
  "lines": [
    {
      "product_id": 1,
      "uom": "PCS",
      "ordered_qty": 50,
      "unit_price": 175000
    }
  ]
}

# Allocate inventory
PUT /api/v1/sales-orders/1/allocate

# Create picking wave
POST /api/v1/picking-waves
{
  "warehouse_id": 1,
  "sales_order_ids": [1, 2, 3]
}

# Record pick
PUT /api/v1/picking-tasks/1/pick
{
  "lines": [
    {
      "picking_task_line_id": 1,
      "picked_qty": 50,
      "from_location_id": 10,
      "batch_id": 5
    }
  ]
}
```

## Files to Create
- `internal/repository/sales_order_repository.go`
- `internal/repository/picking_repository.go`
- `internal/infrastructure/persistence/*_impl.go`
- `internal/usecase/sales_order_usecase.go`
- `internal/usecase/picking_usecase.go`
- `internal/delivery/http/handler/sales_order_handler.go`
- `internal/delivery/http/handler/picking_handler.go`
- `internal/dto/sales_order.go`, `picking.go`

## Dependencies
- MASTER-01 (Warehouse)
- MASTER-02 (Product)
- MASTER-03 (Customer)
- TRANS-02 (need inventory from putaway)
- SETUP-03 (RBAC - picker role)

## Business Rules
- SO number: SO-YYYYMMDD-XXXX
- Wave number: WAVE-YYYYMMDD-XXXX
- Allocation checks available_qty per location
- FIFO strategy for batch selection (oldest batch first)
- Picking sequence optimized by location proximity
