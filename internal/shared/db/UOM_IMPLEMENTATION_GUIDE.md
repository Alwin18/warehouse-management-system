# Penambahan Tabel Unit of Measure (UOM)

## 📋 Yang Telah Dibuat

### 1. Model File
**File**: `internal/shared/db/models/unit_of_measure.go`

Model `UnitOfMeasure` dengan struktur:
- `id` - Primary key
- `code` - Kode UOM unik (PCS, KG, L, dll)
- `name` - Nama lengkap
- `symbol` - Symbol untuk display
- `category` - Kategori (COUNT, WEIGHT, VOLUME, LENGTH, AREA, PACKAGING)
- `description` - Deskripsi detail
- `is_active` - Status aktif/nonaktif
- Timestamps

### 2. Migration Update
**File**: `internal/shared/db/migrate.go`

Ditambahkan `&models.UnitOfMeasure{}` ke AutoMigrate di bagian Product Management.

### 3. Seed File
**File**: `internal/shared/db/seed_uom.go`

Fungsi `SeedUnitOfMeasures()` yang menyediakan **38 UOM standar** across 6 kategori:

#### COUNT (6 UOMs)
- PCS, EA, UNIT, PAIR, SET, DOZEN

#### WEIGHT (6 UOMs)
- KG, G, MG, TON, LB, OZ

#### VOLUME (4 UOMs)
- L, ML, GAL, M3

#### LENGTH (6 UOMs)
- M, CM, MM, KM, FT, IN

#### AREA (2 UOMs)
- M2, FT2

#### PACKAGING (14 UOMs)
- BOX, CARTON, PALLET, CASE, PACK, BAG, ROLL, BOTTLE, CAN, DRUM, CONTAINER

### 4. Dokumentasi Update
**File**: `internal/shared/db/MODELS_DOCUMENTATION.md`

Ditambahkan section lengkap untuk `unit_of_measures` dengan:
- Penjelasan tujuan dan kegunaan
- Detail semua kolom
- Kategori UOM
- Best practices
- Catatan pada `products.base_uom` dan `product_uoms.uom` untuk future migration ke FK

## 🎯 Keuntungan Tabel UOM Master

### 1. Data Quality
✅ **Standardisasi** - Mencegah typo dan inkonsistensi
✅ **Validasi** - Hanya UOM valid yang bisa digunakan
✅ **Kategorisasi** - Grouping berdasarkan tipe

### 2. Functionality
✅ **Metadata** - Symbol, deskripsi, kategori
✅ **Reporting** - Mudah grouping dan filtering
✅ **I18n Ready** - Support multi-language

### 3. Maintenance
✅ **Centralized** - Satu tempat untuk manage UOM
✅ **Soft Delete** - Inactive UOM yang tidak dipakai
✅ **Audit Trail** - Timestamps untuk tracking

## 📝 Cara Menggunakan

### 1. Run Migration
```bash
# Migration akan otomatis create tabel unit_of_measures
# saat aplikasi start
```

### 2. Seed Data
```go
// Di main.go atau setup script
import "path/to/db"

// Seed UOM data
err := db.SeedUnitOfMeasures(dbInstance, logger)
if err != nil {
    log.Fatal(err)
}
```

### 3. Query UOM
```go
// Get all active UOMs
var uoms []models.UnitOfMeasure
db.Where("is_active = ?", true).Find(&uoms)

// Get UOMs by category
var weightUOMs []models.UnitOfMeasure
db.Where("category = ? AND is_active = ?", "WEIGHT", true).Find(&weightUOMs)

// Get specific UOM by code
var uom models.UnitOfMeasure
db.Where("code = ?", "PCS").First(&uom)
```

## 🔄 Future Migration (Optional)

Untuk full integration, pertimbangkan untuk:

### 1. Update Product Model
```go
// Dari:
BaseUOM string `gorm:"type:varchar(20);not null"`

// Ke:
BaseUOMID uint `gorm:"not null;index:idx_product_base_uom"`
BaseUOM   *UnitOfMeasure `gorm:"foreignKey:BaseUOMID"`
```

### 2. Update ProductUOM Model
```go
// Dari:
UOM string `gorm:"type:varchar(20);not null"`

// Ke:
UOMID uint `gorm:"not null;index:idx_product_uom_uom"`
UOM   *UnitOfMeasure `gorm:"foreignKey:UOMID"`
```

### 3. Update Transactional Tables
Semua tabel yang punya kolom `uom` string bisa diubah jadi FK ke `unit_of_measures`.

## ✅ Checklist

- [x] Create `UnitOfMeasure` model
- [x] Add to migration
- [x] Create seed file with 38 standard UOMs
- [x] Update documentation
- [x] Add notes for future FK migration
- [ ] Run migration (user action)
- [ ] Run seed (user action)
- [ ] Optional: Migrate existing models to use FK

## 📊 Total Models

Sekarang sistem memiliki **38 models**:
- **12 Master Data Models** (termasuk `unit_of_measures`)
- **26 Transactional Data Models**
