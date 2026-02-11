### 📋 Master Data Tables (Data Referensi)
Tabel yang berisi data relatif statis, jarang berubah, dan menjadi referensi untuk transaksi:

👥 User & Role
- users
- roles
- user_roles

🏭 Warehouse & Location
- warehouses
- warehouse_zones
- locations

📦 Product
- products
- product_uoms
- product_batches (semi-master, bisa juga dianggap transaksional)
🤝 Partners
- suppliers
- customers
- carriers

## 📊 Transactional Data Tables (Data Operasional)
Tabel yang berisi data transaksi harian, sering berubah:

🛒 Purchase Orders
- purchase_orders
- purchase_order_lines.

📥 Goods Receipt (Inbound)
- goods_receipts
- goods_receipt_lines

🛍️ Sales Orders
- sales_orders
- sales_order_lines

📋 Picking Operations
- picking_waves
- picking_tasks
- picking_task_lines

📤 Putaway Operations
- putaway_tasks
- putaway_task_lines

📊 Inventory
- inventory_balances (semi-transactional, current state)
- inventory_movements (pure transactional log)

🚚 Shipments
- shipments
- shipment_orders
- shipment_packages
- shipment_package_items

🔢 Stock Management
- stock_counts
- stock_count_lines
- stock_adjustments
- stock_adjustment_lines

↩️ Returns
- customer_returns
- customer_return_lines
- supplier_returns
- supplier_return_lines