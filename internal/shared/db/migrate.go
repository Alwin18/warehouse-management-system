package db

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"gorm.io/gorm"
)

// AutoMigrate runs GORM auto migration for all models
func AutoMigrate(db *gorm.DB, log logger.Logger) error {
	log.Info("Starting database auto migration...")

	err := db.AutoMigrate(
		// User Management
		&models.User{},
		&models.Role{},
		&models.UserRole{},

		// Warehouse Management
		&models.Warehouse{},
		&models.WarehouseZone{},
		&models.Location{},

		// Product Management
		&models.UnitOfMeasure{},
		&models.Product{},
		&models.ProductUOM{},
		&models.ProductBatch{},

		// Supplier & Customer
		&models.Supplier{},
		&models.Customer{},
		&models.Carrier{},

		// Purchase Orders
		&models.PurchaseOrder{},
		&models.PurchaseOrderLine{},

		// Goods Receipt (Inbound)
		&models.GoodsReceipt{},
		&models.GoodsReceiptLine{},

		// Putaway Tasks
		&models.PutawayTask{},
		&models.PutawayTaskLine{},

		// Sales Orders
		&models.SalesOrder{},
		&models.SalesOrderLine{},

		// Picking Operations
		&models.PickingWave{},
		&models.PickingTask{},
		&models.PickingTaskLine{},

		// Shipment (Outbound)
		&models.Shipment{},
		&models.ShipmentOrder{},
		&models.ShipmentPackage{},
		&models.ShipmentPackageItem{},

		// Inventory Management
		&models.InventoryBalance{},
		&models.InventoryMovement{},

		// Stock Operations
		&models.StockCount{},
		&models.StockCountLine{},
		&models.StockAdjustment{},
		&models.StockAdjustmentLine{},

		// Returns
		&models.CustomerReturn{},
		&models.CustomerReturnLine{},
		&models.SupplierReturn{},
		&models.SupplierReturnLine{},
	)

	if err != nil {
		log.Error("Failed to run auto migration:", err)
		return err
	}

	log.Info("Database auto migration completed successfully")
	return nil
}
