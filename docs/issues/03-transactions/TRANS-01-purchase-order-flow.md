# [TRANS-01] Purchase Order Flow

**Status**: TODO  
**Priority**: P2 (Medium)  
**Estimated Time**: 5-6 days  
**Assignee**: TBD

## Description
Implement complete Purchase Order creation dan management flow.

## Requirements
- [ ] Create PO dengan multiple line items
- [ ] PO status workflow: DRAFT → SUBMITTED → APPROVED → CLOSED
- [ ] Calculate total amount dari line items
- [ ] Validate product exists
- [ ] Track received quantity per line
- [ ] Auto-close PO ketika semua lines fully received

## Acceptance Criteria
- ✅ POST `/purchase-orders` - create PO
- ✅ GET `/purchase-orders` - list dengan filter status, supplier, date range
- ✅ GET `/purchase-orders/:id` - detail PO dengan lines
- ✅ PUT `/purchase-orders/:id/submit` - submit PO (DRAFT → SUBMITTED)
- ✅ PUT `/purchase-orders/:id/approve` - approve PO (SUBMITTED → APPROVED)
- ✅ PUT `/purchase-orders/:id/lines` - add/update line items
- ✅ GET `/purchase-orders/:id/receipt-status` - receiving progress
- ✅ Only ADMIN/WH_MGR can approve

## API Examples
```bash
# Create PO
POST /api/v1/purchase-orders
{
  "supplier_id": 1,
  "warehouse_id": 1,
  "order_date": "2025-12-12",
  "expected_date": "2025-12-20",
  "currency": "IDR",
  "lines": [
    {
      "product_id": 1,
      "uom": "PCS",
      "ordered_qty": 100,
      "unit_price": 150000,
      "tax_percent": 11
    }
  ]
}

# Submit PO untuk approval
PUT /api/v1/purchase-orders/1/submit

# Get PO with receiving status
GET /api/v1/purchase-orders/1?include=lines,receipt-status
```

## Files to Create
- `internal/repository/purchase_order_repository.go`
- `internal/infrastructure/persistence/purchase_order_repository_impl.go`
- `internal/usecase/purchase_order_usecase.go`
- `internal/delivery/http/handler/purchase_order_handler.go`
- `internal/dto/purchase_order.go`

## Dependencies
- MASTER-01 (Warehouse)
- MASTER-02 (Product)
- MASTER-03 (Supplier)
- SETUP-03 (RBAC)

## Business Rules
- PO number auto-generated: PO-YYYYMMDD-XXXX
- DRAFT status: editable
- SUBMITTED: waiting approval
- APPROVED: ready for receiving
- Line item price calculation: ordered_qty * unit_price * (1 + tax_percent/100)
