# [SETUP-03] Role-Based Access Control (RBAC) Middleware

**Status**: TODO  
**Priority**: P0 (Critical)  
**Estimated Time**: 2-3 days  
**Assignee**: TBD

## Description
Implement RBAC middleware untuk authorization berdasarkan role user.

## Requirements
- [ ] Create RBAC middleware di `pkg/middleware/rbac.go`
- [ ] Support multiple roles per user
- [ ] Support route-level permission checking
- [ ] Return proper 403 Forbidden jika tidak authorized
- [ ] Log unauthorized access attempts

## Acceptance Criteria
- ✅ Middleware `RequireRole("ADMIN")` berfungsi
- ✅ Middleware `RequireAnyRole("ADMIN", "WH_MGR")` berfungsi
- ✅ User dengan multiple roles bisa akses endpoint sesuai role
- ✅ Proper error response untuk unauthorized access
- ✅ JWT claims include user roles

## Technical Notes
```go
// Example usage
router.Post("/warehouses", 
    middleware.RequireRole("ADMIN"),
    handler.CreateWarehouse,
)

router.Get("/picking-tasks", 
    middleware.RequireAnyRole("PICKER", "WH_MGR"),
    handler.GetPickingTasks,
)
```

## Files to Create/Modify
- `pkg/middleware/rbac.go` (new)
- `internal/dto/auth.go` (add roles to JWT claims)
- `internal/usecase/auth_usecase.go` (include roles in token)

## Dependencies
- Current JWT middleware already exists
- User roles already included in login response

## Notes
Roles sudah di-load saat login (lihat `auth_usecase.go`). Tinggal extract dari JWT dan validate.
