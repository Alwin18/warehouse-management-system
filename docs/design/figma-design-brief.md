# Warehouse Hub - Figma Design Brief

## Project Overview
Design modern, intuitive UI/UX untuk **Warehouse Management System (WMS)** yang komprehensif. Sistem ini mengelola seluruh operasional gudang mulai dari penerimaan barang, penyimpanan, picking, packing, hingga pengiriman.

---

## Target Users & Personas

### 1. **Admin / Super User**
- **Needs**: Dashboard overview, master data management, reporting
- **Tech Savvy**: High
- **Primary Device**: Desktop/Laptop
- **Pain Points**: Butuh visibility menyeluruh, approval workflow yang clear

### 2. **Warehouse Manager**
- **Needs**: Operational monitoring, task assignment, inventory overview
- **Tech Savvy**: Medium-High
- **Primary Device**: Desktop/Laptop + Tablet
- **Pain Points**: Real-time updates, bottleneck identification

### 3. **Warehouse Staff (Receiver/Picker/Packer)**
- **Needs**: Simple task list, easy data entry, barcode scanning
- **Tech Savvy**: Low-Medium
- **Primary Device**: Mobile/Tablet (hands-free scanning device)
- **Pain Points**: Complex UI, banyak step, input yang lama

### 4. **Inventory Controller**
- **Needs**: Stock count, adjustment, location management
- **Tech Savvy**: Medium
- **Primary Device**: Desktop + Mobile
- **Pain Points**: Data accuracy, variance resolution

---

## Core Modules & Key Screens

### 📊 Module 1: Dashboard & Analytics
**Screens:**
1. **Main Dashboard** (Admin/Manager view)
   - KPI Cards: Total SKUs, On-hand Qty, Orders Today, Pending Tasks
   - Charts: Inventory trends, Order fulfillment rate, Warehouse utilization
   - Recent Activities feed
   - Quick Actions (Create PO, Create SO, Start Picking)

2. **Warehouse Overview**
   - Warehouse selection dropdown
   - Zone capacity visualization (heatmap/bar chart)
   - Location occupancy status
   - Active tasks by zone

### 📦 Module 2: Master Data Management

**2.1 Warehouse Management**
- **List View**: Table dengan filter (Active, City, Zone Count)
- **Create/Edit Form**: Multi-step wizard
  - Step 1: Basic Info (Code, Name, Address)
  - Step 2: Zones (Add zones dengan type selection)
  - Step 3: Locations (Bulk import atau manual add)
- **Detail View**: Tabs untuk Info, Zones, Locations, Statistics

**2.2 Product Management**
- **List View**: Grid/Table toggle, search bar dengan autocomplete
- **Filter Panel**: Category, Brand, Batch Managed, Active/Inactive
- **Create/Edit Modal**: 
  - Product Info section
  - UOM Conversion table (inline editing)
  - Image upload area
- **Detail View**: Product card dengan specs, UOMs, Current Stock, Movement History
- **Bulk Import**: Drag-drop CSV/Excel dengan preview & validation

**2.3 Partners (Supplier/Customer/Carrier)**
- **List View**: Cards atau Table view
- **Create/Edit Modal**: Form dengan validation
- **Detail View**: Partner info + Transaction history

### 🛒 Module 3: Purchase Order & Receiving

**3.1 Purchase Order**
- **List View**: Status badges (Draft, Submitted, Approved), Filter by date range, supplier
- **Create Wizard**:
  - Step 1: Select Supplier & Warehouse
  - Step 2: Add Items (product search dengan inline add)
  - Step 3: Review & Submit
- **Detail View**: 
  - Header: PO number, status, dates, totals
  - Line items table (dengan received qty progress bar)
  - Approval workflow timeline
  - Action buttons (Submit, Approve, Create GR)

**3.2 Goods Receipt**
- **List View**: Filter by PO, Date, Status
- **Create from PO**: Auto-populate dari PO lines
- **Receiving Screen** (Mobile-optimized):
  - Scan barcode untuk auto-fill product
  - Large qty input dengan +/- buttons
  - Batch/Expiry entry (conditional)
  - QC status selection (chips/radio)
  - Photo upload untuk damage documentation
- **Complete GR**: Summary dengan discrepancies highlighted

**3.3 Putaway Tasks**
- **Task List** (Mobile-first):
  - Card-based layout
  - Product image, name, qty, recommended location
  - Start/Complete buttons
- **Execution Screen**:
  - Source location (highlighted)
  - Destination location (dengan navigation hints)
  - Scan confirmation
  - Override destination option

### 🛍️ Module 4: Sales Order & Picking

**4.1 Sales Order**
- **List View**: Priority badges, Status chips, Customer filter
- **Create Flow**: Similar to PO (Wizard atau Single page)
- **Allocation Screen**:
  - Product availability indicator
  - Auto-allocation preview
  - Manual location override

**4.2 Picking Waves**
- **Wave Creation**: Multi-select SOs, Auto-optimize picking route
- **Wave Dashboard**: Tasks by picker, Completion %, ETA

**4.3 Picking Tasks** (Mobile-optimized)
- **Task List**: Sortable by priority/sequence
- **Execution Screen**:
  - Step-by-step pick list
  - Location highlighted dengan map/zone indicator
  - Qty confirmation dengan scan verification
  - Short-pick handling (reason codes)
  - Progress indicator

### 📊 Module 5: Inventory Management

**5.1 Inventory Balance**
- **Grid View**: Location layout (visual warehouse map)
- **Table View**: Filterable by product, location, status, batch
- **Detail Panel**: On-hand, Reserved, Available dengan color coding
- **Drill-down**: Click to see movement history

**5.2 Stock Count**
- **Count Creation**: Select location/zone/product range
- **Count Execution** (Mobile):
  - Location-by-location walk-through
  - Blind count atau system qty comparison
  - Variance alert (threshold-based)
- **Variance Resolution**: Approve/Reject dengan reason codes

**5.3 Stock Adjustment**
- **Adjustment Form**: Simple reason code dropdown + qty delta
- **Batch Adjustment**: Multi-line editor (spreadsheet-like)
- **Approval Workflow**: Untuk adjustment di atas threshold

### 🚚 Module 6: Shipment & Packing

**6.1 Packing**
- **Pack Station Screen**:
  - SO details panel
  - Items to pack checklist
  - Package creation (dimensions, weight)
  - Label printing preview

**6.2 Shipment**
- **Shipment Creation**: Group multiple SOs
- **Carrier Selection**: Rate shopping (future)
- **Tracking Dashboard**: Status timeline untuk shipments

---

## Design Requirements

### Visual Design Language

**Color Palette:**
```
Primary:      #2563EB (Blue) - Actions, Links
Success:      #10B981 (Green) - Completed, Available
Warning:      #F59E0B (Amber) - Pending, Reserved
Danger:       #EF4444 (Red) - Errors, Critical
Info:         #06B6D4 (Cyan) - Info, Tips
Neutral:      #6B7280 (Gray) - Text, Borders

Background:   #F9FAFB (Light Gray)
Surface:      #FFFFFF (White)
Text Primary: #111827 (Dark Gray)
Text Sidebar: #374151 (Medium Gray)
```

**Typography:**
```
Headings:     Inter / Poppins (Bold 600-700)
Body:         Inter / Open Sans (Regular 400, Medium 500)
Monospace:    JetBrains Mono (untuk SKU, codes, numbers)

Sizes:
H1: 32px
H2: 24px
H3: 20px
H4: 18px
Body: 14px
Small: 12px
```

**Components:**
- **Buttons**: Rounded corners (6px), Clear hierarchy (Primary/Secondary/Ghost)
- **Cards**: Shadow sm, Border subtle, Hover effects
- **Tables**: Zebra striping, Fixed header on scroll, Row actions dropdown
- **Forms**: Clear labels, Inline validation, Helpful placeholders
- **Status Badges**: Pills dengan color coding, Icon support
- **Modals**: Centered, Max-width responsive, Clear close action
- **Toasts**: Top-right position, Auto-dismiss, Action support

### Responsive Behavior

**Desktop (1920×1080 / 1366×768)**
- Sidebar navigation (collapsible)
- Multi-column layouts
- Data tables dengan banyak kolom
- Dashboard dengan multiple widgets

**Tablet (1024×768 / 768×1024)**
- Sidebar → Top navbar (hamburger menu)
- 2-column layouts
- Simplified tables (hide non-critical columns)
- Larger touch targets (min 44px)

**Mobile (375×667 / 414×896)**
- Bottom navigation untuk staff apps
- Single column layouts
- Card-based lists (table → cards)
- Full-screen modals
- Large buttons untuk scanning (min 56px)

### UX Principles

**1. Progressive Disclosure**
- Jangan tampilkan semua fields sekaligus
- Gunakan wizards/steps untuk complex forms
- Expand/collapse untuk advanced options

**2. Error Prevention**
- Validation real-time (inline errors)
- Confirmation dialogs untuk destructive actions
- Undo capability untuk accidental changes

**3. Efficiency for Power Users**
- Keyboard shortcuts (document di tooltips)
- Bulk actions (multi-select)
- Quick filters & search
- Recent/Favorites untuk quick access

**4. Mobile-First for Warehouse Staff**
- Barcode scan first, manual input fallback
- Large touch targets
- Minimal typing (dropdowns, chips, toggles)
- Offline support indicators

**5. Clear Status & Feedback**
- Loading states (skeletons, spinners)
- Success/error toasts
- Progress indicators untuk long operations
- Empty states dengan call-to-action

---

## Specific Screen Flows

### Critical User Flow 1: Receiving Goods
```
1. Warehouse Staff opens app → sees "Pending Receipts" list
2. Tap GR-001 → sees expected items
3. Scan barcode → auto-fills product
4. Enter qty (large numeric keypad) → tap Next
5. If batch-managed: Enter batch & expiry → tap Next
6. Select QC status (chips: Pass/Pending/Fail) → tap Complete
7. Success toast → returns to list
8. System auto-generates putaway task
```

**UI Elements Needed:**
- GR list cards dengan status badges
- Barcode scan button (prominent, camera icon)
- Large numeric input dengan +/- steppers
- Batch/Expiry date picker (calendar)
- QC status selection (color-coded chips)
- Photo attachment button (untuk damage)

### Critical User Flow 2: Picking Items
```
1. Picker sees assigned task list → sorted by location sequence
2. Tap "Start Picking" → sees first pick
3. Navigate to location (map/zone highlighted)
4. Scan location barcode → verifies correct location
5. Scan product barcode → verifies correct item
6. Enter picked qty → tap Confirm
7. If short pick: Select reason → tap Continue
8. Repeat until task complete → tap "Complete Task"
9. System updates inventory & SO status
```

**UI Elements Needed:**
- Task cards dengan sequence numbers
- Warehouse map (simplified, zone-based)
- Dual scan (location + product verification)
- Qty adjustment dengan reason codes dropdown
- Progress tracker (X of Y items)
- Short-pick flow (reason + note)

---

## Navigation Structure

### Desktop App (Admin/Manager)
**Sidebar Menu:**
```
🏠 Dashboard
📦 Master Data
   ├─ Warehouses
   ├─ Products
   ├─ Suppliers
   ├─ Customers
   └─ Carriers
🛒 Inbound
   ├─ Purchase Orders
   ├─ Goods Receipts
   └─ Putaway Tasks
🛍️ Outbound
   ├─ Sales Orders
   ├─ Picking Waves
   ├─ Picking Tasks
   └─ Shipments
📊 Inventory
   ├─ Balance
   ├─ Movements
   ├─ Stock Counts
   └─ Adjustments
📈 Reports
👤 Settings
   ├─ Users
   ├─ Roles
   └─ Preferences
```

### Mobile App (Warehouse Staff)
**Bottom Navigation:**
```
🏠 Home (Task dashboard)
📦 Receive
📋 Pick
🔢 Count
👤 Profile
```

---

## Deliverables

### Phase 1: Core Screens (Priority)
- [ ] Dashboard (Desktop)
- [ ] Warehouse List & Detail
- [ ] Product List & Form
- [ ] Purchase Order Create & Detail
- [ ] Goods Receipt (Mobile)
- [ ] Picking Task (Mobile)
- [ ] Inventory Balance View

### Phase 2: Complete Flows
- [ ] Complete Master Data screens
- [ ] Complete Inbound flow
- [ ] Complete Outbound flow
- [ ] Stock Management screens
- [ ] Mobile app complete flows

### Phase 3: Polish & Components
- [ ] Design System documentation (Colors, Typography, Components)
- [ ] Icon set (custom atau dari library)
- [ ] Empty states
- [ ] Error states
- [ ] Loading states
- [ ] Onboarding/Tutorial screens

### Export Requirements
- Organized Figma file dengan proper naming
- Component library (reusable components)
- Variants untuk different states
- Responsive frames (Desktop/Tablet/Mobile)
- Auto-layout untuk responsive behavior
- Developer handoff annotations
- Export assets (icons, illustrations)

---

## Design Inspiration & References

**Style References:**
- Modern SaaS dashboards: Linear, Notion, Airtable
- E-commerce admin: Shopify Admin, WooCommerce
- Logistics apps: ShipBob, Deliverr, Flexport
- Warehouse-specific: Fishbowl, Cin7, Katana MRP

**UI Patterns:**
- Tables: Retool, Airtable
- Forms: Typeform, Google Forms (multi-step)
- Mobile scanning: Amazon Seller App, Shopify POS
- Status workflows: Jira, Trello

**Component Libraries:**
- Tailwind UI
- Radix UI
- Shadcn/ui
- Material Design 3

---

## Technical Constraints & Considerations

### Backend Integration Points
- REST API dengan pagination
- Real-time updates (WebSocket untuk live inventory)
- Barcode scanning (Camera API atau dedicated scanner)
- File upload (CSV import, images)
- Print support (labels, packing slips)

### Accessibility Requirements
- WCAG 2.1 Level AA compliance
- Keyboard navigation support
- Screen reader friendly (semantic HTML)
- Sufficient color contrast (4.5:1 minimum)
- Focus indicators clear
- Alt text untuk images

### Performance Considerations
- Lazy loading untuk large lists
- Virtual scrolling untuk tables >100 rows
- Image optimization
- Debounced search inputs
- Optimistic UI updates

---

## Timeline & Milestones

**Week 1-2: Research & Wireframes**
- User flow mapping
- Low-fidelity wireframes
- Navigation structure finalization

**Week 3-4: High-Fidelity Designs (Phase 1)**
- Core screens visual design
- Component library start
- First round of reviews

**Week 5-6: High-Fidelity Designs (Phase 2)**
- Complete flow designs
- Responsive variants
- Interaction design (prototypes)

**Week 7-8: Polish & Handoff**
- Design system documentation
- Developer handoff preparation
- Final revisions

---

## Questions for Designer

Before starting, please clarify:
1. Company branding guidelines available? (Logo, colors, fonts)
2. Existing design system to follow or start from scratch?
3. Specific industry compliance requirements? (e.g., cold storage, pharma)
4. Target markets? (International = i18n considerations)
5. Hardware constraints? (Specific scanner devices, screen sizes)
6. Dark mode required?
7. Print requirements? (Label sizes, formats)

---

## Success Metrics

Post-launch, design should enable:
- **Task completion time**: 30% faster vs. current process
- **Error rate**: Less than 2% picking errors
- **User satisfaction**: >4.5/5 rating from warehouse staff
- **Training time**: New staff productive within 2 hours
- **Mobile adoption**: >80% of tasks done via mobile

---

**Contact for Questions:**
[Your contact info here]

**Target Completion Date:**
[Insert date]

**Budget:**
[If applicable]
