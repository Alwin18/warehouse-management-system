# 📚 Dokumentasi Database Models - Warehouse Management System

> Dokumentasi lengkap untuk semua model database dalam sistem manajemen gudang

## 📑 Daftar Isi

- [Master Data Models](#-master-data-models)
  - [User & Role Management](#1-user--role-management)
  - [Warehouse & Location](#2-warehouse--location)
  - [Product Management](#3-product-management)
  - [Partner Management](#4-partner-management)
- [Transactional Data Models](#-transactional-data-models)
  - [Purchase Orders](#1-purchase-orders)
  - [Goods Receipt (Inbound)](#2-goods-receipt-inbound)
  - [Sales Orders](#3-sales-orders)
  - [Picking Operations](#4-picking-operations)
  - [Putaway Operations](#5-putaway-operations)
  - [Inventory Management](#6-inventory-management)
  - [Shipments](#7-shipments)
  - [Stock Management](#8-stock-management)
  - [Returns](#9-returns)

---

# 📋 Master Data Models

Master data adalah data referensi yang relatif statis dan jarang berubah. Data ini menjadi fondasi untuk semua transaksi operasional.

## 1. User & Role Management

### 👤 users

**Tujuan**: Menyimpan informasi pengguna sistem yang dapat mengakses aplikasi warehouse management.

**Kolom-kolom**:
- `id` (uint, PK): ID unik pengguna
- `username` (string, unique, required): Username untuk login, harus unik
- `password_hash` (string, required): Password yang sudah di-hash untuk keamanan
- `full_name` (string, required): Nama lengkap pengguna
- `email` (string, unique, nullable): Email pengguna, harus unik jika diisi
- `phone` (string, nullable): Nomor telepon pengguna
- `is_active` (boolean, default: true): Status aktif/nonaktif pengguna
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Many-to-Many dengan `roles` melalui tabel `user_roles`

**Kegunaan**:
- Autentikasi dan otorisasi pengguna
- Tracking siapa yang membuat/mengubah transaksi
- Manajemen akses sistem

---

### 🔐 roles

**Tujuan**: Mendefinisikan peran/role yang dapat dimiliki pengguna untuk mengatur hak akses.

**Kolom-kolom**:
- `id` (uint, PK): ID unik role
- `code` (string, unique, required): Kode role (misal: ADMIN, WAREHOUSE_STAFF, MANAGER)
- `name` (string, required): Nama role yang mudah dibaca
- `description` (text, nullable): Deskripsi detail tentang role
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Many-to-Many dengan `users` melalui tabel `user_roles`

**Kegunaan**:
- Mendefinisikan level akses berbeda dalam sistem
- Implementasi Role-Based Access Control (RBAC)
- Pemisahan tugas dan tanggung jawab

---

### 🔗 user_roles

**Tujuan**: Tabel junction untuk menghubungkan users dengan roles (many-to-many relationship).

**Kolom-kolom**:
- `user_id` (uint, PK, FK): Referensi ke tabel users
- `role_id` (uint, PK, FK): Referensi ke tabel roles
- `created_at` (timestamp): Waktu assignment role
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to `users` (CASCADE on delete)
- Belongs to `roles` (CASCADE on delete)

**Kegunaan**:
- Menghubungkan user dengan satu atau lebih role
- Satu user bisa memiliki multiple roles
- Satu role bisa dimiliki oleh multiple users

---

## 2. Warehouse & Location

### 🏭 warehouses

**Tujuan**: Menyimpan informasi tentang gudang-gudang yang dikelola dalam sistem.

**Kolom-kolom**:
- `id` (uint, PK): ID unik warehouse
- `code` (string, unique, required): Kode warehouse yang unik
- `name` (string, required): Nama warehouse
- `address` (text, nullable): Alamat lengkap warehouse
- `city` (string, nullable): Kota lokasi warehouse
- `country` (string, nullable): Negara lokasi warehouse
- `time_zone` (string, nullable): Timezone warehouse untuk tracking waktu
- `is_active` (boolean, default: true): Status aktif/nonaktif warehouse
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Has Many: `warehouse_zones`, `locations`, `purchase_orders`, `goods_receipts`, `sales_orders`, `picking_waves`, `picking_tasks`, `putaway_tasks`, `inventory_balances`, `inventory_movements`, `shipments`, `stock_counts`, `stock_adjustments`, `customer_returns`, `supplier_returns`

**Kegunaan**:
- Central point untuk semua operasi warehouse
- Multi-warehouse support
- Segregasi data per warehouse
- Tracking lokasi fisik operasi

---

### 🗺️ warehouse_zones

**Tujuan**: Membagi warehouse menjadi zona-zona untuk organisasi yang lebih baik.

**Kolom-kolom**:
- `id` (uint, PK): ID unik zone
- `warehouse_id` (uint, FK, required): Referensi ke warehouse
- `code` (string, required): Kode zone
- `name` (string, required): Nama zone
- `zone_type` (string, required): Tipe zone (misal: RECEIVING, STORAGE, PICKING, SHIPPING)
- `is_active` (boolean, default: true): Status aktif/nonaktif zone
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to `warehouses` (CASCADE on delete)
- Has Many `locations`

**Kegunaan**:
- Organisasi warehouse berdasarkan fungsi
- Optimasi alur kerja (receiving area, storage area, shipping area)
- Strategi picking yang lebih efisien
- Segregasi inventory berdasarkan karakteristik

---

### 📍 locations

**Tujuan**: Menyimpan lokasi spesifik dalam warehouse untuk penyimpanan barang.

**Kolom-kolom**:
- `id` (uint, PK): ID unik location
- `warehouse_id` (uint, FK, required): Referensi ke warehouse
- `zone_id` (uint, FK, required): Referensi ke zone
- `code` (string, required): Kode lokasi (misal: A-01-01, B-02-03)
- `description` (text, nullable): Deskripsi lokasi
- `location_type` (string, required): Tipe lokasi (misal: SHELF, PALLET, FLOOR, BIN)
- `max_volume` (decimal, nullable): Kapasitas volume maksimum
- `max_weight` (decimal, nullable): Kapasitas berat maksimum
- `is_active` (boolean, default: true): Status aktif/nonaktif lokasi
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to `warehouses` (CASCADE on delete)
- Belongs to `warehouse_zones` (CASCADE on delete)
- Has Many: `inventory_balances`, `stock_count_lines`, `stock_adjustment_lines`
- Referenced by: `inventory_movements` (from/to), `picking_task_lines` (from), `putaway_task_lines` (source/destination)

**Kegunaan**:
- Tracking posisi exact barang dalam warehouse
- Manajemen kapasitas penyimpanan
- Optimasi putaway dan picking
- Slotting strategy implementation

---

## 3. Product Management

### 📐 unit_of_measures

**Tujuan**: Master data untuk standardisasi unit of measure yang digunakan di seluruh sistem.

**Kolom-kolom**:
- `id` (uint, PK): ID unik UOM
- `code` (string, unique, required): Kode UOM yang unik (misal: PCS, KG, L, BOX)
- `name` (string, required): Nama lengkap UOM (misal: Pieces, Kilogram, Liter, Box)
- `symbol` (string, nullable): Symbol untuk display (misal: pcs, kg, L)
- `category` (string, required): Kategori UOM (COUNT, WEIGHT, VOLUME, LENGTH, AREA, PACKAGING)
- `description` (text, nullable): Deskripsi detail UOM
- `is_active` (boolean, default: true): Status aktif/nonaktif UOM
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Referenced by: `products` (base_uom), `product_uoms` (uom), dan semua tabel transaksional yang menggunakan UOM

**Kegunaan**:
- **Standardisasi**: Mencegah typo dan inkonsistensi (misal: "PCS" vs "pcs" vs "Pcs")
- **Validasi**: Enforce hanya UOM yang valid yang bisa digunakan
- **Kategorisasi**: Grouping UOM berdasarkan tipe (berat, volume, count, dll)
- **Metadata**: Menyimpan informasi tambahan seperti symbol dan deskripsi
- **Reporting**: Memudahkan grouping dan filtering dalam report
- **Internationalization**: Support multi-language untuk nama UOM

**Kategori UOM**:
- **COUNT**: PCS, EA, UNIT, PAIR, SET, DOZEN - untuk item yang dihitung per piece
- **WEIGHT**: KG, G, MG, TON, LB, OZ - untuk produk yang dijual per berat
- **VOLUME**: L, ML, GAL, M3 - untuk produk liquid atau gas
- **LENGTH**: M, CM, MM, KM, FT, IN - untuk produk yang dijual per panjang
- **AREA**: M2, FT2 - untuk produk yang dijual per luas area
- **PACKAGING**: BOX, CARTON, PALLET, CASE, PACK, BAG, ROLL, BOTTLE, CAN, DRUM, CONTAINER - untuk unit packaging

**Best Practices**:
- Seed dengan UOM standar saat initial setup
- Gunakan kode yang konsisten dengan standar industri (ISO, dll)
- Maintain kategori untuk memudahkan validasi business rules
- Inactive UOM yang tidak digunakan lagi, jangan delete

---

### 📦 products

**Tujuan**: Menyimpan informasi master produk yang dikelola dalam warehouse.

**Kolom-kolom**:
- `id` (uint, PK): ID unik produk
- `sku` (string, unique, required): Stock Keeping Unit, identifier unik produk
- `name` (string, required): Nama produk
- `barcode` (string, nullable): Barcode produk untuk scanning
- `description` (text, nullable): Deskripsi detail produk
- `base_uom` (string, required): Unit of Measure dasar (misal: PCS, KG, LITER) - **Catatan**: Sebaiknya diubah menjadi FK ke `unit_of_measures`
- `weight` (decimal, nullable): Berat produk
- `volume` (decimal, nullable): Volume produk
- `is_batch_managed` (boolean, default: false): Apakah produk dikelola per batch
- `is_serialized` (boolean, default: false): Apakah produk memiliki serial number
- `is_active` (boolean, default: true): Status aktif/nonaktif produk
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Has Many: `product_uoms`, `product_batches`, `purchase_order_lines`, `goods_receipt_lines`, `sales_order_lines`, `picking_task_lines`, `putaway_task_lines`, `inventory_balances`, `inventory_movements`, `shipment_package_items`, `stock_count_lines`, `stock_adjustment_lines`, `customer_return_lines`, `supplier_return_lines`

**Kegunaan**:
- Master data untuk semua transaksi inventory
- Tracking karakteristik fisik produk
- Mendukung berbagai strategi inventory (batch, serial)
- Dasar untuk perhitungan kapasitas dan space utilization

---

### 📏 product_uoms

**Tujuan**: Mendefinisikan berbagai unit of measure untuk satu produk dengan konversi.

**Kolom-kolom**:
- `id` (uint, PK): ID unik UOM
- `product_id` (uint, FK, required): Referensi ke produk
- `uom` (string, required): Unit of measure (misal: BOX, CARTON, PALLET) - **Catatan**: Sebaiknya diubah menjadi FK ke `unit_of_measures`
- `conversion_to_base` (decimal, required): Faktor konversi ke base UOM
- `is_default_sales` (boolean, default: false): UOM default untuk penjualan
- `is_default_purchase` (boolean, default: false): UOM default untuk pembelian
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to `products` (CASCADE on delete)
- Should reference `unit_of_measures` untuk standardisasi

**Kegunaan**:
- Mendukung transaksi dalam berbagai UOM
- Konversi otomatis antar UOM
- Fleksibilitas dalam pembelian dan penjualan
- Contoh: 1 BOX = 12 PCS, 1 PALLET = 50 BOX

---

### 🏷️ product_batches

**Tujuan**: Tracking batch/lot produk untuk traceability dan manajemen expiry.

**Kolom-kolom**:
- `id` (uint, PK): ID unik batch
- `product_id` (uint, FK, required): Referensi ke produk
- `batch_number` (string, required): Nomor batch/lot dari supplier
- `expiry_date` (date, nullable): Tanggal kadaluarsa batch
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to `products` (CASCADE on delete)
- Has Many: `goods_receipt_lines`, `inventory_balances`, `inventory_movements`, `picking_task_lines`, `putaway_task_lines`, `stock_count_lines`, `stock_adjustment_lines`, `supplier_return_lines`

**Kegunaan**:
- Traceability produk dari supplier ke customer
- Manajemen FEFO (First Expired First Out)
- Quality control per batch
- Recall management jika ada masalah kualitas

---

## 4. Partner Management

### 🏢 suppliers

**Tujuan**: Menyimpan informasi supplier/vendor yang memasok produk.

**Kolom-kolom**:
- `id` (uint, PK): ID unik supplier
- `code` (string, unique, required): Kode supplier yang unik
- `name` (string, required): Nama supplier
- `address` (text, nullable): Alamat supplier
- `city` (string, nullable): Kota supplier
- `country` (string, nullable): Negara supplier
- `phone` (string, nullable): Nomor telepon supplier
- `email` (string, nullable): Email supplier
- `tax_id` (string, nullable): Nomor pajak/NPWP supplier
- `is_active` (boolean, default: true): Status aktif/nonaktif supplier
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Has Many: `purchase_orders`, `goods_receipts`, `supplier_returns`

**Kegunaan**:
- Master data supplier untuk procurement
- Tracking performance supplier
- Manajemen hubungan dengan vendor
- Dokumentasi untuk compliance dan audit

---

### 👥 customers

**Tujuan**: Menyimpan informasi customer yang membeli produk.

**Kolom-kolom**:
- `id` (uint, PK): ID unik customer
- `code` (string, unique, required): Kode customer yang unik
- `name` (string, required): Nama customer
- `address` (text, nullable): Alamat customer
- `city` (string, nullable): Kota customer
- `country` (string, nullable): Negara customer
- `phone` (string, nullable): Nomor telepon customer
- `email` (string, nullable): Email customer
- `tax_id` (string, nullable): Nomor pajak customer
- `is_active` (boolean, default: true): Status aktif/nonaktif customer
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Has Many: `sales_orders`, `customer_returns`

**Kegunaan**:
- Master data customer untuk sales
- Shipping address management
- Customer service dan tracking
- Sales analytics dan reporting

---

### 🚚 carriers

**Tujuan**: Menyimpan informasi perusahaan pengiriman/kurir.

**Kolom-kolom**:
- `id` (uint, PK): ID unik carrier
- `code` (string, unique, required): Kode carrier yang unik
- `name` (string, required): Nama carrier (misal: JNE, TIKI, DHL)
- `is_active` (boolean, default: true): Status aktif/nonaktif carrier
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Has Many: `shipments`

**Kegunaan**:
- Manajemen partner pengiriman
- Tracking performance carrier
- Shipping cost analysis
- Carrier selection untuk optimasi biaya dan waktu

---

# 📊 Transactional Data Models

Data transaksional adalah data yang sering berubah dan merekam aktivitas operasional harian warehouse.

## 1. Purchase Orders

### 🛒 purchase_orders

**Tujuan**: Merekam pesanan pembelian barang dari supplier.

**Kolom-kolom**:
- `id` (uint, PK): ID unik PO
- `po_number` (string, unique, required): Nomor PO yang unik
- `supplier_id` (uint, FK, required): Referensi ke supplier
- `warehouse_id` (uint, FK, required): Warehouse tujuan
- `status` (string, required): Status PO (DRAFT, CONFIRMED, PARTIAL, COMPLETED, CANCELLED)
- `order_date` (date, required): Tanggal order dibuat
- `expected_date` (date, nullable): Tanggal kedatangan yang diharapkan
- `currency` (string, nullable): Mata uang transaksi
- `total_amount` (decimal, nullable): Total nilai PO
- `created_by` (uint, FK, required): User yang membuat PO
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `suppliers` (RESTRICT), `warehouses` (RESTRICT), `users` as creator (RESTRICT)
- Has Many: `purchase_order_lines`, `goods_receipts`

**Kegunaan**:
- Procurement planning dan execution
- Tracking expected inbound
- Budget control dan approval workflow
- Supplier performance monitoring

---

### 📝 purchase_order_lines

**Tujuan**: Detail item-item dalam purchase order.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `purchase_order_id` (uint, FK, required): Referensi ke PO header
- `line_no` (int, required): Nomor urut line
- `product_id` (uint, FK, required): Produk yang dipesan
- `uom` (string, required): Unit of measure
- `ordered_qty` (decimal, required): Jumlah yang dipesan
- `received_qty` (decimal, default: 0): Jumlah yang sudah diterima
- `unit_price` (decimal, nullable): Harga per unit
- `tax_percent` (decimal, nullable): Persentase pajak
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `purchase_orders` (CASCADE), `products` (RESTRICT)
- Has Many: `goods_receipt_lines`

**Kegunaan**:
- Detail produk yang dipesan
- Tracking received vs ordered quantity
- Costing dan pricing information
- Matching dengan goods receipt

---

## 2. Goods Receipt (Inbound)

### 📥 goods_receipts

**Tujuan**: Merekam penerimaan barang masuk ke warehouse.

**Kolom-kolom**:
- `id` (uint, PK): ID unik GR
- `gr_number` (string, unique, required): Nomor goods receipt yang unik
- `purchase_order_id` (uint, FK, nullable): Referensi ke PO (jika ada)
- `warehouse_id` (uint, FK, required): Warehouse penerima
- `supplier_id` (uint, FK, nullable): Supplier pengirim
- `status` (string, required): Status GR (DRAFT, RECEIVED, QUALITY_CHECK, COMPLETED)
- `received_at` (timestamp, required): Waktu penerimaan barang
- `received_by` (uint, FK, required): User yang menerima
- `external_ref` (string, nullable): Referensi eksternal (delivery note, invoice)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `purchase_orders` (SET NULL), `warehouses` (RESTRICT), `suppliers` (SET NULL), `users` as receiver (RESTRICT)
- Has Many: `goods_receipt_lines`, `putaway_tasks`

**Kegunaan**:
- Dokumentasi penerimaan barang
- Quality control checkpoint
- Trigger untuk putaway process
- Inventory increase transaction

---

### 📋 goods_receipt_lines

**Tujuan**: Detail item-item yang diterima dalam goods receipt.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `goods_receipt_id` (uint, FK, required): Referensi ke GR header
- `purchase_order_line_id` (uint, FK, nullable): Referensi ke PO line
- `line_no` (int, required): Nomor urut line
- `product_id` (uint, FK, required): Produk yang diterima
- `uom` (string, required): Unit of measure
- `received_qty` (decimal, required): Jumlah yang diterima
- `batch_id` (uint, FK, nullable): Batch produk
- `serial_number` (string, nullable): Serial number (untuk serialized product)
- `qc_status` (string, nullable): Status quality check (PASS, FAIL, PENDING)
- `source_location_id` (uint, FK, nullable): Lokasi receiving area
- `note` (text, nullable): Catatan tambahan
- `created_at` (timestamp): Waktu pembuatan record

**Relasi**:
- Belongs to: `goods_receipts` (CASCADE), `purchase_order_lines` (SET NULL), `products` (RESTRICT), `product_batches` (SET NULL), `locations` as source (SET NULL)
- Has Many: `putaway_task_lines`

**Kegunaan**:
- Detail barang yang diterima
- Batch/serial tracking
- Quality control documentation
- Matching dengan PO untuk verification

---

## 3. Sales Orders

### 🛍️ sales_orders

**Tujuan**: Merekam pesanan penjualan dari customer.

**Kolom-kolom**:
- `id` (uint, PK): ID unik SO
- `so_number` (string, unique, required): Nomor sales order yang unik
- `external_ref` (string, nullable): Referensi eksternal (order number dari e-commerce)
- `customer_id` (uint, FK, required): Customer pemesan
- `warehouse_id` (uint, FK, required): Warehouse untuk fulfillment
- `status` (string, required): Status SO (PENDING, CONFIRMED, PICKING, PACKED, SHIPPED, COMPLETED, CANCELLED)
- `order_date` (timestamp, required): Tanggal order
- `requested_ship_date` (timestamp, nullable): Tanggal pengiriman yang diminta
- `priority` (string, nullable): Prioritas order (HIGH, MEDIUM, LOW)
- `shipping_address` (text, nullable): Alamat pengiriman
- `shipping_city` (string, nullable): Kota pengiriman
- `shipping_country` (string, nullable): Negara pengiriman
- `shipping_phone` (string, nullable): Nomor telepon penerima
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `customers` (RESTRICT), `warehouses` (RESTRICT)
- Has Many: `sales_order_lines`, `picking_tasks`, `shipment_orders`, `customer_returns`

**Kegunaan**:
- Order management dan fulfillment
- Shipping planning
- Priority-based picking
- Customer service tracking

---

### 📦 sales_order_lines

**Tujuan**: Detail item-item dalam sales order.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `sales_order_id` (uint, FK, required): Referensi ke SO header
- `line_no` (int, required): Nomor urut line
- `product_id` (uint, FK, required): Produk yang dipesan
- `uom` (string, required): Unit of measure
- `ordered_qty` (decimal, required): Jumlah yang dipesan
- `allocated_qty` (decimal, default: 0): Jumlah yang sudah dialokasikan
- `shipped_qty` (decimal, default: 0): Jumlah yang sudah dikirim
- `unit_price` (decimal, nullable): Harga per unit
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `sales_orders` (CASCADE), `products` (RESTRICT)
- Has Many: `picking_task_lines`, `shipment_package_items`, `customer_return_lines`

**Kegunaan**:
- Detail produk yang dipesan
- Tracking allocated vs shipped quantity
- Backorder management
- Revenue calculation

---

## 4. Picking Operations

### 🌊 picking_waves

**Tujuan**: Mengelompokkan multiple picking tasks untuk efisiensi.

**Kolom-kolom**:
- `id` (uint, PK): ID unik wave
- `wave_number` (string, unique, required): Nomor wave yang unik
- `warehouse_id` (uint, FK, required): Warehouse
- `status` (string, required): Status wave (CREATED, RELEASED, IN_PROGRESS, COMPLETED)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir
- `created_by` (uint, FK, required): User yang membuat wave

**Relasi**:
- Belongs to: `warehouses` (RESTRICT), `users` as creator (RESTRICT)
- Has Many: `picking_tasks`

**Kegunaan**:
- Batch picking untuk efisiensi
- Wave picking strategy
- Workload balancing
- Performance tracking per wave

---

### 📋 picking_tasks

**Tujuan**: Task picking untuk memenuhi sales order.

**Kolom-kolom**:
- `id` (uint, PK): ID unik task
- `picking_wave_id` (uint, FK, nullable): Referensi ke wave (jika part of wave)
- `sales_order_id` (uint, FK, required): Sales order yang dipick
- `warehouse_id` (uint, FK, required): Warehouse
- `assigned_to` (uint, FK, nullable): User yang ditugaskan
- `status` (string, required): Status task (PENDING, ASSIGNED, IN_PROGRESS, COMPLETED, CANCELLED)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir
- `started_at` (timestamp, nullable): Waktu mulai picking
- `completed_at` (timestamp, nullable): Waktu selesai picking

**Relasi**:
- Belongs to: `picking_waves` (SET NULL), `sales_orders` (RESTRICT), `warehouses` (RESTRICT), `users` as assignee (SET NULL)
- Has Many: `picking_task_lines`

**Kegunaan**:
- Assignment picking ke picker
- Tracking picking progress
- Performance measurement (time to pick)
- Workload management

---

### 📝 picking_task_lines

**Tujuan**: Detail item-item yang harus dipick dalam task.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `picking_task_id` (uint, FK, required): Referensi ke picking task
- `sales_order_line_id` (uint, FK, required): Referensi ke SO line
- `product_id` (uint, FK, required): Produk yang dipick
- `from_location_id` (uint, FK, required): Lokasi sumber picking
- `batch_id` (uint, FK, nullable): Batch yang dipick
- `uom` (string, required): Unit of measure
- `planned_qty` (decimal, required): Jumlah yang direncanakan
- `picked_qty` (decimal, default: 0): Jumlah yang sudah dipick
- `sequence_no` (int, required): Urutan picking untuk optimasi rute
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `picking_tasks` (CASCADE), `sales_order_lines` (RESTRICT), `products` (RESTRICT), `locations` as from (RESTRICT), `product_batches` (SET NULL)

**Kegunaan**:
- Instruksi picking detail
- Route optimization
- Batch/FEFO compliance
- Variance tracking (planned vs picked)

---

## 5. Putaway Operations

### 📤 putaway_tasks

**Tujuan**: Task untuk menyimpan barang yang diterima ke lokasi penyimpanan.

**Kolom-kolom**:
- `id` (uint, PK): ID unik task
- `warehouse_id` (uint, FK, required): Warehouse
- `goods_receipt_id` (uint, FK, required): Goods receipt yang di-putaway
- `assigned_to` (uint, FK, nullable): User yang ditugaskan
- `status` (string, required): Status task (PENDING, ASSIGNED, IN_PROGRESS, COMPLETED, CANCELLED)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir
- `started_at` (timestamp, nullable): Waktu mulai putaway
- `completed_at` (timestamp, nullable): Waktu selesai putaway

**Relasi**:
- Belongs to: `warehouses` (RESTRICT), `goods_receipts` (RESTRICT), `users` as assignee (SET NULL)
- Has Many: `putaway_task_lines`

**Kegunaan**:
- Assignment putaway ke warehouse staff
- Tracking putaway progress
- Performance measurement
- Completion of inbound process

---

### 📝 putaway_task_lines

**Tujuan**: Detail item-item yang harus di-putaway.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `putaway_task_id` (uint, FK, required): Referensi ke putaway task
- `goods_receipt_line_id` (uint, FK, required): Referensi ke GR line
- `product_id` (uint, FK, required): Produk yang di-putaway
- `source_location_id` (uint, FK, required): Lokasi sumber (receiving area)
- `destination_location_id` (uint, FK, required): Lokasi tujuan (storage)
- `batch_id` (uint, FK, nullable): Batch produk
- `uom` (string, required): Unit of measure
- `planned_qty` (decimal, required): Jumlah yang direncanakan
- `putaway_qty` (decimal, default: 0): Jumlah yang sudah di-putaway
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `putaway_tasks` (CASCADE), `goods_receipt_lines` (RESTRICT), `products` (RESTRICT), `locations` as source (RESTRICT), `locations` as destination (RESTRICT), `product_batches` (SET NULL)

**Kegunaan**:
- Instruksi putaway detail
- Slotting strategy execution
- Space optimization
- Inventory placement tracking

---

## 6. Inventory Management

### 📊 inventory_balances

**Tujuan**: Menyimpan current state inventory per lokasi, produk, dan batch.

**Kolom-kolom**:
- `id` (uint, PK): ID unik balance
- `warehouse_id` (uint, FK, required): Warehouse
- `location_id` (uint, FK, required): Lokasi penyimpanan
- `product_id` (uint, FK, required): Produk
- `batch_id` (uint, FK, nullable): Batch (jika batch-managed)
- `status` (string, required): Status inventory (AVAILABLE, RESERVED, QUARANTINE, DAMAGED)
- `on_hand_qty` (decimal, default: 0): Jumlah fisik di lokasi
- `reserved_qty` (decimal, default: 0): Jumlah yang sudah direserve untuk order
- `available_qty` (decimal, default: 0): Jumlah available (on_hand - reserved)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `warehouses` (CASCADE), `locations` (CASCADE), `products` (CASCADE), `product_batches` (SET NULL)

**Kegunaan**:
- Real-time inventory visibility
- Available-to-promise calculation
- Reservation management
- Multi-status inventory tracking

---

### 📈 inventory_movements

**Tujuan**: Log semua pergerakan inventory untuk audit trail dan traceability.

**Kolom-kolom**:
- `id` (uint, PK): ID unik movement
- `movement_type` (string, required): Tipe movement (RECEIPT, PUTAWAY, PICK, SHIPMENT, ADJUSTMENT, TRANSFER, RETURN)
- `warehouse_id` (uint, FK, required): Warehouse
- `product_id` (uint, FK, required): Produk
- `batch_id` (uint, FK, nullable): Batch
- `from_location_id` (uint, FK, nullable): Lokasi asal
- `to_location_id` (uint, FK, nullable): Lokasi tujuan
- `qty` (decimal, required): Jumlah yang bergerak
- `uom` (string, required): Unit of measure
- `status_before` (string, nullable): Status sebelum movement
- `status_after` (string, nullable): Status setelah movement
- `reference_type` (string, nullable): Tipe referensi (GOODS_RECEIPT, PICKING_TASK, etc)
- `reference_id` (uint, nullable): ID referensi
- `created_at` (timestamp, indexed): Waktu movement
- `created_by` (uint, FK, nullable): User yang melakukan
- `note` (text, nullable): Catatan tambahan

**Relasi**:
- Belongs to: `warehouses` (CASCADE), `products` (CASCADE), `product_batches` (SET NULL), `locations` as from (SET NULL), `locations` as to (SET NULL), `users` as creator (SET NULL)

**Kegunaan**:
- Complete audit trail
- Traceability untuk compliance
- Inventory reconciliation
- Analytics dan reporting
- Debugging inventory discrepancies

---

## 7. Shipments

### 🚚 shipments

**Tujuan**: Mengelola pengiriman barang ke customer.

**Kolom-kolom**:
- `id` (uint, PK): ID unik shipment
- `shipment_number` (string, unique, required): Nomor shipment yang unik
- `warehouse_id` (uint, FK, required): Warehouse asal
- `carrier_id` (uint, FK, nullable): Carrier yang digunakan
- `status` (string, required): Status shipment (PREPARING, DISPATCHED, IN_TRANSIT, DELIVERED, CANCELLED)
- `dispatch_time` (timestamp, nullable): Waktu dispatch
- `delivered_time` (timestamp, nullable): Waktu delivered
- `created_at` (timestamp): Waktu pembuatan record

**Relasi**:
- Belongs to: `warehouses` (RESTRICT), `carriers` (SET NULL)
- Has Many: `shipment_orders`, `shipment_packages`

**Kegunaan**:
- Consolidation multiple orders dalam satu shipment
- Carrier management
- Delivery tracking
- Proof of delivery

---

### 🔗 shipment_orders

**Tujuan**: Menghubungkan shipment dengan sales orders (many-to-many).

**Kolom-kolom**:
- `id` (uint, PK): ID unik
- `shipment_id` (uint, FK, required): Referensi ke shipment
- `sales_order_id` (uint, FK, required): Referensi ke sales order
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `shipments` (CASCADE), `sales_orders` (RESTRICT)

**Kegunaan**:
- Multiple orders dalam satu shipment
- Order consolidation
- Shipping cost optimization

---

### 📦 shipment_packages

**Tujuan**: Detail paket-paket dalam shipment.

**Kolom-kolom**:
- `id` (uint, PK): ID unik package
- `shipment_id` (uint, FK, required): Referensi ke shipment
- `package_number` (string, unique, required): Nomor package yang unik
- `tracking_number` (string, nullable): Tracking number dari carrier
- `weight` (decimal, nullable): Berat package
- `volume` (decimal, nullable): Volume package
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `shipments` (CASCADE)
- Has Many: `shipment_package_items`

**Kegunaan**:
- Package-level tracking
- Dimensional weight calculation
- Carrier integration
- Multi-package shipment support

---

### 📋 shipment_package_items

**Tujuan**: Detail item-item dalam setiap package.

**Kolom-kolom**:
- `id` (uint, PK): ID unik item
- `shipment_package_id` (uint, FK, required): Referensi ke package
- `sales_order_line_id` (uint, FK, required): Referensi ke SO line
- `product_id` (uint, FK, required): Produk
- `uom` (string, required): Unit of measure
- `qty` (decimal, required): Jumlah dalam package
- `created_at` (timestamp): Waktu pembuatan record

**Relasi**:
- Belongs to: `shipment_packages` (CASCADE), `sales_order_lines` (RESTRICT), `products` (RESTRICT)

**Kegunaan**:
- Packing list generation
- Content verification
- Customs documentation
- Delivery confirmation

---

## 8. Stock Management

### 🔢 stock_counts

**Tujuan**: Mengelola cycle counting dan physical inventory.

**Kolom-kolom**:
- `id` (uint, PK): ID unik count
- `count_number` (string, unique, required): Nomor stock count yang unik
- `warehouse_id` (uint, FK, required): Warehouse
- `status` (string, required): Status count (PLANNED, IN_PROGRESS, COMPLETED, CANCELLED)
- `count_type` (string, required): Tipe count (FULL, CYCLE, SPOT)
- `scheduled_at` (timestamp, nullable): Waktu dijadwalkan
- `completed_at` (timestamp, nullable): Waktu selesai
- `created_by` (uint, FK, required): User yang membuat
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `warehouses` (RESTRICT), `users` as creator (RESTRICT)
- Has Many: `stock_count_lines`

**Kegunaan**:
- Inventory accuracy improvement
- Cycle counting program
- Variance identification
- Audit compliance

---

### 📝 stock_count_lines

**Tujuan**: Detail hasil counting per lokasi dan produk.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `stock_count_id` (uint, FK, required): Referensi ke stock count
- `location_id` (uint, FK, required): Lokasi yang dicount
- `product_id` (uint, FK, required): Produk
- `batch_id` (uint, FK, nullable): Batch
- `system_qty` (decimal, required): Jumlah menurut sistem
- `counted_qty` (decimal, nullable): Jumlah hasil counting
- `variance_qty` (decimal, nullable): Selisih (counted - system)
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `stock_counts` (CASCADE), `locations` (RESTRICT), `products` (RESTRICT), `product_batches` (SET NULL)

**Kegunaan**:
- Recording count results
- Variance analysis
- Adjustment preparation
- Root cause investigation

---

### ⚙️ stock_adjustments

**Tujuan**: Menyesuaikan inventory berdasarkan hasil count atau kondisi lain.

**Kolom-kolom**:
- `id` (uint, PK): ID unik adjustment
- `adjustment_number` (string, unique, required): Nomor adjustment yang unik
- `warehouse_id` (uint, FK, required): Warehouse
- `reason_code` (string, required): Kode alasan (COUNT_VARIANCE, DAMAGE, EXPIRY, THEFT, etc)
- `status` (string, required): Status adjustment (DRAFT, POSTED, CANCELLED)
- `created_by` (uint, FK, required): User yang membuat
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir
- `posted_at` (timestamp, nullable): Waktu posting ke inventory

**Relasi**:
- Belongs to: `warehouses` (RESTRICT), `users` as creator (RESTRICT)
- Has Many: `stock_adjustment_lines`

**Kegunaan**:
- Inventory correction
- Shrinkage recording
- Damage/expiry write-off
- Audit trail untuk adjustments

---

### 📝 stock_adjustment_lines

**Tujuan**: Detail adjustment per produk dan lokasi.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `stock_adjustment_id` (uint, FK, required): Referensi ke adjustment
- `location_id` (uint, FK, required): Lokasi
- `product_id` (uint, FK, required): Produk
- `batch_id` (uint, FK, nullable): Batch
- `qty_delta` (decimal, required): Perubahan quantity (+ atau -)
- `uom` (string, required): Unit of measure
- `note` (text, nullable): Catatan
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `stock_adjustments` (CASCADE), `locations` (RESTRICT), `products` (RESTRICT), `product_batches` (SET NULL)

**Kegunaan**:
- Detail perubahan inventory
- Linking ke inventory movement
- Documentation untuk audit
- Cost impact calculation

---

## 9. Returns

### ↩️ customer_returns

**Tujuan**: Mengelola return barang dari customer.

**Kolom-kolom**:
- `id` (uint, PK): ID unik return
- `return_number` (string, unique, required): Nomor return yang unik
- `sales_order_id` (uint, FK, required): Referensi ke sales order original
- `customer_id` (uint, FK, required): Customer yang return
- `warehouse_id` (uint, FK, required): Warehouse penerima return
- `status` (string, required): Status return (REQUESTED, APPROVED, RECEIVED, COMPLETED, REJECTED)
- `reason` (text, nullable): Alasan return
- `requested_at` (timestamp, required): Waktu request return
- `received_at` (timestamp, nullable): Waktu barang diterima kembali
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `sales_orders` (RESTRICT), `customers` (RESTRICT), `warehouses` (RESTRICT)
- Has Many: `customer_return_lines`

**Kegunaan**:
- Return authorization (RMA)
- Customer satisfaction management
- Refund/replacement processing
- Return rate analytics

---

### 📝 customer_return_lines

**Tujuan**: Detail item-item yang di-return oleh customer.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `customer_return_id` (uint, FK, required): Referensi ke customer return
- `sales_order_line_id` (uint, FK, required): Referensi ke SO line original
- `product_id` (uint, FK, required): Produk yang di-return
- `uom` (string, required): Unit of measure
- `returned_qty` (decimal, required): Jumlah yang di-return
- `qc_status` (string, nullable): Status QC (PASS, FAIL, PENDING)
- `return_reason_code` (string, nullable): Kode alasan return
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `customer_returns` (CASCADE), `sales_order_lines` (RESTRICT), `products` (RESTRICT)

**Kegunaan**:
- Detail barang yang di-return
- Quality inspection
- Restocking decision
- Return reason analysis

---

### ↩️ supplier_returns

**Tujuan**: Mengelola return barang ke supplier.

**Kolom-kolom**:
- `id` (uint, PK): ID unik return
- `return_number` (string, unique, required): Nomor return yang unik
- `supplier_id` (uint, FK, required): Supplier tujuan return
- `warehouse_id` (uint, FK, required): Warehouse asal return
- `status` (string, required): Status return (DRAFT, APPROVED, SHIPPED, COMPLETED, CANCELLED)
- `reason` (text, nullable): Alasan return
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir
- `created_by` (uint, FK, required): User yang membuat

**Relasi**:
- Belongs to: `suppliers` (RESTRICT), `warehouses` (RESTRICT), `users` as creator (RESTRICT)
- Has Many: `supplier_return_lines`

**Kegunaan**:
- Defective goods return
- Overstocked items return
- Supplier claim management
- Quality issue escalation

---

### 📝 supplier_return_lines

**Tujuan**: Detail item-item yang di-return ke supplier.

**Kolom-kolom**:
- `id` (uint, PK): ID unik line
- `supplier_return_id` (uint, FK, required): Referensi ke supplier return
- `product_id` (uint, FK, required): Produk yang di-return
- `batch_id` (uint, FK, nullable): Batch yang di-return
- `uom` (string, required): Unit of measure
- `qty` (decimal, required): Jumlah yang di-return
- `reason_code` (string, nullable): Kode alasan return
- `created_at` (timestamp): Waktu pembuatan record
- `updated_at` (timestamp): Waktu update terakhir

**Relasi**:
- Belongs to: `supplier_returns` (CASCADE), `products` (RESTRICT), `product_batches` (SET NULL)

**Kegunaan**:
- Detail barang yang di-return
- Batch traceability
- Supplier performance tracking
- Credit note processing

---

## 🔄 Relasi Antar Model

### Diagram Relasi Utama

```
Warehouses (1) ──┬── (N) Warehouse Zones (1) ─── (N) Locations
                 │
                 ├── (N) Purchase Orders (1) ─── (N) Purchase Order Lines
                 │                                    │
                 ├── (N) Goods Receipts (1) ─────────┴─── (N) Goods Receipt Lines
                 │         │                                      │
                 │         └─── (N) Putaway Tasks (1) ─── (N) Putaway Task Lines
                 │
                 ├── (N) Sales Orders (1) ─── (N) Sales Order Lines
                 │         │                           │
                 │         └─── (N) Picking Tasks (1) ─┴─── (N) Picking Task Lines
                 │                    │
                 │                    └─── (N) Picking Waves
                 │
                 ├── (N) Inventory Balances
                 ├── (N) Inventory Movements
                 ├── (N) Shipments (1) ─── (N) Shipment Packages (1) ─── (N) Shipment Package Items
                 ├── (N) Stock Counts (1) ─── (N) Stock Count Lines
                 ├── (N) Stock Adjustments (1) ─── (N) Stock Adjustment Lines
                 ├── (N) Customer Returns (1) ─── (N) Customer Return Lines
                 └── (N) Supplier Returns (1) ─── (N) Supplier Return Lines

Products (1) ──┬── (N) Product UOMs
               ├── (N) Product Batches
               └── (N) [Referenced by all transactional lines]

Users (N) ─── (N) Roles (via user_roles junction table)

Suppliers (1) ─── (N) Purchase Orders, Goods Receipts, Supplier Returns
Customers (1) ─── (N) Sales Orders, Customer Returns
Carriers (1) ─── (N) Shipments
```

---

## 💡 Best Practices

### 1. **Status Management**
Setiap tabel transaksional memiliki kolom `status` untuk tracking lifecycle:
- Gunakan status yang konsisten dan meaningful
- Implement state machine untuk validasi transisi status
- Log perubahan status untuk audit

### 2. **Soft Delete vs Hard Delete**
- Gunakan `is_active` untuk master data (soft delete)
- Hard delete untuk transactional data hanya jika diperlukan
- Perhatikan constraint `OnDelete` pada foreign keys

### 3. **Indexing Strategy**
- Foreign keys sudah di-index otomatis
- Tambahkan index pada kolom yang sering di-query (status, dates)
- Composite index untuk query yang kompleks

### 4. **Data Integrity**
- Gunakan transactions untuk operasi multi-table
- Validate business rules di application layer
- Implement optimistic locking untuk concurrent updates

### 5. **Performance Optimization**
- Partition large tables (inventory_movements) by date
- Archive old transactional data
- Use materialized views untuk reporting

---

## 📈 Use Cases dan Flow

### Inbound Flow
1. Create **Purchase Order** → Purchase Order Lines
2. Receive goods → **Goods Receipt** → Goods Receipt Lines
3. Create **Putaway Task** → Putaway Task Lines
4. Execute putaway → Update **Inventory Balance** + **Inventory Movement**

### Outbound Flow
1. Create **Sales Order** → Sales Order Lines
2. Create **Picking Task** (optional: in **Picking Wave**)
3. Execute picking → Picking Task Lines
4. Create **Shipment** → Shipment Packages → Shipment Package Items
5. Ship → Update **Inventory Balance** + **Inventory Movement**

### Inventory Control Flow
1. Create **Stock Count** → Stock Count Lines
2. Perform counting → Update counted_qty
3. Analyze variance → Create **Stock Adjustment** → Stock Adjustment Lines
4. Post adjustment → Update **Inventory Balance** + **Inventory Movement**

---

## 🎯 Summary

Sistem database warehouse management ini dirancang dengan:
- **38 models** yang comprehensive (termasuk `unit_of_measures` untuk standardisasi UOM)
- **Master-Transactional separation** untuk clarity
- **Complete audit trail** melalui inventory_movements
- **Standardized UOM** dengan master table untuk data consistency
- **Flexible UOM** support dengan konversi otomatis
- **Batch/Serial tracking** capability
- **Multi-warehouse** support
- **Status-driven workflows**
- **Proper relationships** dengan cascade rules yang tepat

Semua model saling terintegrasi untuk mendukung end-to-end warehouse operations dari procurement hingga shipping, dengan full traceability dan inventory accuracy.
