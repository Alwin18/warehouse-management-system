# [TRANS-02] Goods Receipt & Putaway Flow

**Status**: TODO  
**Priority**: P2 (Medium)  
**Estimated Time**: 6-7 days  
**Assignee**: TBD

## Description
Implement Goods Receipt (inbound receiving) dan Putaway task generation.

## Requirements
- [ ] Create GR dari PO atau standalone
- [ ] Record received quantities per line
- [ ] Support batch/lot tracking
- [ ] QC status per line
- [ ] Auto-generate putaway tasks
- [ ] Update PO received_qty
- [ ] Create inventory movement records

## Acceptance Criteria
- ✅ POST `/goods-receipts` - create GR from PO
- ✅ PUT `/goods-receipts/:id/receive` - record received items
- ✅ POST `/goods-receipts/:id/qc` - update QC status
- ✅ PUT `/goods-receipts/:id/complete` - complete GR & generate putaway
- ✅ GET `/goods-receipts` - list dengan filter
- ✅ Update PO line received_qty automatically
- ✅ Create batch records for batch-managed products
- ✅ Generate putaway tasks to staging locations

## API Examples
```bash
# Create GR from PO
POST /api/v1/goods-receipts
{
  "purchase_order_id": 1,
  "warehouse_id": 1,
  "received_at": "2025-12-12T10:00:00Z"
}

# Receive items
PUT /api/v1/goods-receipts/1/receive
{
  "lines": [
    {
      "purchase_order_line_id": 1,
      "product_id": 1,
      "received_qty": 95,
      "batch_number": "BATCH-20251212-001",
      "expiry_date": "2026-12-12",
      "qc_status": "PENDING"
    }
  ]
}

# Complete GR (generates putaway tasks)
PUT /api/v1/goods-receipts/1/complete
```

## Files to Create
- `internal/repository/goods_receipt_repository.go`
- `internal/repository/putaway_task_repository.go`
- `internal/infrastructure/persistence/*_impl.go`
- `internal/usecase/goods_receipt_usecase.go`
- `internal/usecase/putaway_usecase.go`
- `internal/delivery/http/handler/goods_receipt_handler.go`
- `internal/dto/goods_receipt.go`

## Dependencies
- TRANS-01 (Purchase Order)
- MASTER-01 (Warehouse - need staging locations)
- MASTER-02 (Product - batch management)

## Business Rules
- GR number: GR-YYYYMMDD-XXXX
- QC Status: PENDING, PASSED, FAILED
- Received qty can be less than ordered (partial receipt)
- Staging location should be zone type = RECEIVING
- Putaway destinations should be zone type = STORAGE
