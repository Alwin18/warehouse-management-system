package warehouse

type CreateWarehouseRequest struct {
	Code     string  `json:"code" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Address  *string `json:"address"`
	City     *string `json:"city"`
	Country  *string `json:"country"`
	TimeZone *string `json:"time_zone"`
	IsActive *bool   `json:"is_active"`
}

type ListWarehouseRequest struct {
	Page     int    `json:"page" query:"page"`
	PerPage  int    `json:"per_page" query:"per_page"`
	Search   string `json:"search" query:"search"`
	IsActive *bool  `json:"is_active" query:"is_active"`
}

type ListWarehouseResponse struct {
	ID            uint    `json:"id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Address       *string `json:"address"`
	City          *string `json:"city"`
	Country       *string `json:"country"`
	TimeZone      *string `json:"time_zone"`
	IsActive      *bool   `json:"is_active"`
	TotalZone     int64   `json:"total_zone"`
	TotalLocation int64   `json:"total_location"`
}

type CreateWarehouseZoneRequest struct {
	WarehouseID uint   `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	ZoneType    string `json:"zone_type" validate:"required"`
	IsActive    *bool  `json:"is_active"`
}

type CreateWarehouseLocationRequest struct {
	WarehouseID  uint     `json:"warehouse_id" validate:"required"`
	ZoneID       uint     `json:"zone_id" validate:"required"`
	Code         string   `json:"code" validate:"required"`
	IsActive     *bool    `json:"is_active"`
	Description  *string  `json:"description"`
	LocationType string   `json:"location_type" validate:"required"`
	MaxVolume    *float64 `json:"max_volume"`
	MaxWeight    *float64 `json:"max_weight"`
}
