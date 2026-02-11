# Warehouse Management System - Business Flow Documentation

Dokumentasi ini menjelaskan alur bisnis end-to-end dari Warehouse Management System, mulai dari procurement hingga delivery ke customer.

---

## 📋 Table of Contents

1. [Overview - Complete End-to-End Flow](#overview-complete-end-to-end-flow)
2. [Flow 1: Inbound Process (Procurement to Storage)](#flow-1-inbound-process)
3. [Flow 2: Outbound Process (Order to Delivery)](#flow-2-outbound-process)
4. [Flow 3: Inventory Management](#flow-3-inventory-management)
5. [Flow 4: Returns Management](#flow-4-returns-management)
6. [System Architecture Flow](#system-architecture-flow)

---

## Overview: Complete End-to-End Flow

Diagram ini menunjukkan gambaran besar seluruh proses warehouse dari procurement hingga delivery.

```mermaid
graph TB
    subgraph "PROCUREMENT"
        A[Create Purchase Order] --> B[Submit for Approval]
        B --> C{Approved?}
        C -->|Yes| D[PO Approved]
        C -->|No| E[Revise PO]
        E --> B
    end
    
    subgraph "INBOUND OPERATIONS"
        D --> F[Goods Receipt]
        F --> G[Quality Check]
        G --> H{QC Pass?}
        H -->|Pass| I[Generate Putaway Task]
        H -->|Fail| J[Reject/Return to Supplier]
        H -->|Pending| K[Hold in QC Area]
        I --> L[Execute Putaway]
        L --> M[Update Inventory Balance]
    end
    
    subgraph "STORAGE"
        M --> N[(Inventory in Warehouse)]
        N --> O[Stock Count & Adjustment]
        O --> N
    end
    
    subgraph "OUTBOUND OPERATIONS"
        P[Receive Sales Order] --> Q[Allocate Inventory]
        Q --> R{Stock Available?}
        R -->|Yes| S[Create Picking Wave]
        R -->|No| T[Backorder/Purchase]
        T --> D
        S --> U[Generate Picking Tasks]
        U --> V[Execute Picking]
        V --> W[Packing]
        W --> X[Create Shipment]
    end
    
    subgraph "DELIVERY"
        X --> Y[Assign Carrier]
        Y --> Z[Ship to Customer]
        Z --> AA[Delivery Confirmation]
    end
    
    subgraph "RETURNS"
        AA --> AB{Return?}
        AB -->|Customer Return| AC[Process Customer Return]
        AB -->|No Return| AD[Complete]
        AC --> AE{Condition?}
        AE -->|Good| N
        AE -->|Damaged| AF[Dispose/Supplier Return]
    end
    
    N --> Q
    
    style A fill:#e3f2fd
    style F fill:#fff3e0
    style P fill:#f3e5f5
    style Z fill:#e8f5e9
    style N fill:#fce4ec
```

---

## Flow 1: Inbound Process

Proses lengkap dari pembuatan Purchase Order hingga barang tersimpan di warehouse.

### 1.1 Purchase Order Flow

```mermaid
sequenceDiagram
    participant Buyer
    participant System
    participant Manager
    participant Supplier
    participant DB
    
    Buyer->>System: Create Purchase Order
    System->>DB: Save PO (Status: DRAFT)
    Buyer->>System: Add PO Lines (Products)
    System->>DB: Save PO Lines
    Buyer->>System: Submit PO
    System->>DB: Update Status to SUBMITTED
    System->>Manager: Notification (Approval Required)
    
    alt Approved
        Manager->>System: Approve PO
        System->>DB: Update Status to APPROVED
        System->>Supplier: Send PO (Email/API)
        System->>Buyer: Notification (PO Approved)
    else Rejected
        Manager->>System: Reject PO
        System->>DB: Update Status to REJECTED
        System->>Buyer: Notification (PO Rejected)
    end
```

### 1.2 Goods Receipt & Putaway Flow

```mermaid
graph TB
    subgraph "Goods Receipt Process"
        A[Supplier Delivers Goods] --> B[Warehouse Staff Opens GR]
        B --> C[Scan/Select PO]
        C --> D[Scan Product Barcode]
        D --> E[Enter Received Quantity]
        E --> F{Batch Managed?}
        F -->|Yes| G[Enter Batch Number & Expiry]
        F -->|No| H[Quality Check]
        G --> H
        H --> I{QC Status}
        I -->|Pass| J[Save GR Line - Status: PASS]
        I -->|Fail| K[Save GR Line - Status: FAIL]
        I -->|Pending| L[Save GR Line - Status: PENDING]
        J --> M{More Items?}
        K --> M
        L --> M
        M -->|Yes| D
        M -->|No| N[Complete Goods Receipt]
    end
    
    subgraph "Putaway Process"
        N --> O{QC Pass Items Exist?}
        O -->|Yes| P[Auto Generate Putaway Tasks]
        O -->|No| Q[End - No Putaway Needed]
        P --> R[Assign to Warehouse Staff]
        R --> S[Staff Scans Source Location]
        S --> T[Staff Scans Product]
        T --> U[System Suggests Destination]
        U --> V{Accept Suggestion?}
        V -->|Yes| W[Scan Destination Location]
        V -->|No| X[Manual Select Location]
        X --> W
        W --> Y[Confirm Putaway]
        Y --> Z[Update Inventory Balance]
        Z --> AA{More Tasks?}
        AA -->|Yes| S
        AA -->|No| AB[Complete Putaway]
    end
    
    style J fill:#c8e6c9
    style K fill:#ffcdd2
    style L fill:#fff9c4
    style Z fill:#b2dfdb
```

### 1.3 Inventory Update Flow

```mermaid
graph LR
    A[Putaway Completed] --> B[Create Inventory Movement]
    B --> C{Location Has Stock?}
    C -->|Yes| D[Update Existing Balance]
    C -->|No| E[Create New Balance Record]
    D --> F[Calculate New On-Hand Qty]
    E --> F
    F --> G[Save Inventory Balance]
    G --> H[Log Movement History]
    H --> I{Reorder Point Reached?}
    I -->|Yes| J[Trigger Reorder Alert]
    I -->|No| K[End]
    
    style G fill:#81c784
    style J fill:#ffb74d
```

---

## Flow 2: Outbound Process

Proses lengkap dari penerimaan Sales Order hingga pengiriman ke customer.

### 2.1 Sales Order & Allocation Flow

```mermaid
graph TB
    subgraph "Sales Order Creation"
        A[Customer Places Order] --> B[Create Sales Order]
        B --> C[Add Order Lines]
        C --> D[Select Warehouse]
        D --> E[Submit Order]
    end
    
    subgraph "Inventory Allocation"
        E --> F[Check Stock Availability]
        F --> G{Stock Sufficient?}
        G -->|Yes - All Items| H[Auto Allocate Inventory]
        G -->|Partial| I[Partial Allocation]
        G -->|No Stock| J[Create Backorder]
        
        H --> K[Reserve Inventory]
        I --> K
        I --> J
        
        K --> L[Update SO Status: ALLOCATED]
        J --> M[Update SO Status: BACKORDER]
        M --> N[Trigger Purchase Process]
    end
    
    subgraph "Picking Wave Creation"
        L --> O{Multiple Orders?}
        O -->|Yes| P[Batch into Picking Wave]
        O -->|No| Q[Create Single Wave]
        P --> R[Optimize Picking Route]
        Q --> R
        R --> S[Generate Picking Tasks]
    end
    
    style K fill:#81c784
    style J fill:#ffb74d
    style S fill:#64b5f6
```

### 2.2 Picking & Packing Flow

```mermaid
sequenceDiagram
    participant Picker
    participant MobileApp
    participant System
    participant Packer
    participant DB
    
    Note over Picker,DB: PICKING PHASE
    Picker->>MobileApp: Login & View Tasks
    MobileApp->>System: Get Assigned Picking Tasks
    System->>DB: Fetch Tasks (Sorted by Location)
    DB-->>MobileApp: Return Task List
    
    loop For Each Pick
        Picker->>MobileApp: Start Pick
        MobileApp->>Picker: Show Location & Product
        Picker->>MobileApp: Scan Location Barcode
        MobileApp->>System: Validate Location
        Picker->>MobileApp: Scan Product Barcode
        MobileApp->>System: Validate Product
        Picker->>MobileApp: Enter Picked Quantity
        
        alt Full Pick
            MobileApp->>System: Confirm Pick
            System->>DB: Update Inventory (Reduce On-Hand)
            System->>DB: Update Task Status: COMPLETED
        else Short Pick
            Picker->>MobileApp: Select Reason Code
            MobileApp->>System: Record Short Pick
            System->>DB: Update with Actual Qty
            System->>System: Trigger Replenishment Alert
        end
    end
    
    Picker->>MobileApp: Complete All Picks
    MobileApp->>System: Mark Wave as PICKED
    System->>Packer: Notification (Ready for Packing)
    
    Note over Packer,DB: PACKING PHASE
    Packer->>System: Open Packing Station
    System->>DB: Get Picked Items for SO
    Packer->>System: Scan Items to Verify
    Packer->>System: Create Package (Enter Dimensions)
    System->>DB: Create Shipment Package
    Packer->>System: Print Packing Slip & Label
    Packer->>System: Complete Packing
    System->>DB: Update SO Status: PACKED
```

### 2.3 Shipment & Delivery Flow

```mermaid
graph TB
    A[Packing Completed] --> B[Create Shipment]
    B --> C[Group Multiple Orders?]
    C -->|Yes| D[Batch Shipment]
    C -->|No| E[Single Shipment]
    D --> F[Select Carrier]
    E --> F
    F --> G[Generate Shipping Label]
    G --> H[Assign Tracking Number]
    H --> I[Update Shipment Status: READY]
    I --> J[Carrier Pickup]
    J --> K[Update Status: IN_TRANSIT]
    K --> L[Track Shipment]
    L --> M{Delivered?}
    M -->|Yes| N[Update Status: DELIVERED]
    M -->|In Progress| L
    M -->|Failed| O[Update Status: FAILED]
    O --> P[Reschedule Delivery]
    P --> J
    N --> Q[Send Delivery Confirmation]
    Q --> R[Update SO Status: COMPLETED]
    
    style N fill:#66bb6a
    style O fill:#ef5350
    style R fill:#42a5f5
```

---

## Flow 3: Inventory Management

Proses pengelolaan stok, stock count, dan adjustment.

### 3.1 Stock Count (Cycle Count) Flow

```mermaid
graph TB
    subgraph "Planning Phase"
        A[Create Stock Count] --> B[Select Count Type]
        B --> C{Count Type}
        C -->|Full Count| D[Select Entire Warehouse]
        C -->|Cycle Count| E[Select Specific Zones/Locations]
        C -->|Product Count| F[Select Specific Products]
        D --> G[Generate Count Tasks]
        E --> G
        F --> G
        G --> H[Assign to Counter]
    end
    
    subgraph "Execution Phase"
        H --> I[Counter Opens Task]
        I --> J[Navigate to Location]
        J --> K{Count Mode}
        K -->|Blind Count| L[Count Without System Qty]
        K -->|Comparison| M[See System Qty]
        L --> N[Enter Counted Quantity]
        M --> N
        N --> O[Save Count Line]
        O --> P{More Locations?}
        P -->|Yes| J
        P -->|No| Q[Complete Count]
    end
    
    subgraph "Variance Resolution"
        Q --> R[Calculate Variances]
        R --> S{Variance Exists?}
        S -->|Yes| T[Review Variance]
        S -->|No| U[Close Count - No Adjustment]
        T --> V{Variance Threshold}
        V -->|Within Limit| W[Auto Approve]
        V -->|Exceeds Limit| X[Require Manager Approval]
        W --> Y[Create Stock Adjustment]
        X --> Z{Approved?}
        Z -->|Yes| Y
        Z -->|No| AA[Recount Required]
        AA --> I
        Y --> AB[Update Inventory Balance]
    end
    
    style AB fill:#4caf50
    style AA fill:#ff9800
    style U fill:#2196f3
```

### 3.2 Stock Adjustment Flow

```mermaid
sequenceDiagram
    participant User
    participant System
    participant Manager
    participant DB
    
    User->>System: Create Stock Adjustment
    System->>User: Show Adjustment Form
    User->>System: Select Product & Location
    User->>System: Enter Adjustment Qty (+/-)
    User->>System: Select Reason Code
    Note over User,System: Reasons: Damage, Expiry,<br/>Found, Lost, etc.
    
    System->>System: Calculate New Balance
    System->>User: Show Preview
    
    alt Small Adjustment (Within Threshold)
        User->>System: Submit Adjustment
        System->>DB: Create Adjustment Record
        System->>DB: Update Inventory Balance
        System->>DB: Log Inventory Movement
        System->>User: Adjustment Completed
    else Large Adjustment (Exceeds Threshold)
        User->>System: Submit for Approval
        System->>DB: Save as PENDING_APPROVAL
        System->>Manager: Approval Request
        
        alt Approved
            Manager->>System: Approve Adjustment
            System->>DB: Update Status to APPROVED
            System->>DB: Update Inventory Balance
            System->>User: Notification (Approved)
        else Rejected
            Manager->>System: Reject with Comment
            System->>DB: Update Status to REJECTED
            System->>User: Notification (Rejected)
        end
    end
```

---

## Flow 4: Returns Management

Proses pengelolaan return dari customer dan return ke supplier.

### 4.1 Customer Return Flow

```mermaid
graph TB
    subgraph "Return Request"
        A[Customer Requests Return] --> B[Create Return Order]
        B --> C[Link to Original SO]
        C --> D[Select Return Items]
        D --> E[Specify Return Reason]
        E --> F[Submit Return Request]
    end
    
    subgraph "Return Authorization"
        F --> G{Manager Review}
        G -->|Approve| H[Generate RMA Number]
        G -->|Reject| I[Notify Customer - Rejected]
        H --> J[Send RMA to Customer]
    end
    
    subgraph "Return Receipt"
        J --> K[Customer Ships Back]
        K --> L[Warehouse Receives Return]
        L --> M[Scan RMA Number]
        M --> N[Inspect Items]
        N --> O{Item Condition}
        O -->|Good - Resellable| P[Return to Stock]
        O -->|Damaged| Q[Mark as Damaged]
        O -->|Defective| R[Quarantine for Supplier Return]
        
        P --> S[Update Inventory Balance]
        Q --> T[Create Disposal Record]
        R --> U[Create Supplier Return]
    end
    
    subgraph "Refund Processing"
        S --> V[Process Refund/Exchange]
        T --> V
        U --> V
        V --> W[Update Return Status: COMPLETED]
    end
    
    style P fill:#81c784
    style Q fill:#ffb74d
    style R fill:#e57373
```

### 4.2 Supplier Return Flow

```mermaid
graph LR
    A[Identify Defective Items] --> B[Create Supplier Return]
    B --> C[Link to Original PO/GR]
    C --> D[Add Return Lines]
    D --> E[Specify Return Reason]
    E --> F[Submit to Supplier]
    F --> G{Supplier Approval}
    G -->|Approved| H[Generate Return Label]
    G -->|Rejected| I[Dispose Internally]
    H --> J[Pack Items]
    J --> K[Ship to Supplier]
    K --> L[Update Inventory - Remove Stock]
    L --> M{Refund/Credit?}
    M -->|Refund| N[Record Refund]
    M -->|Credit Note| O[Record Credit]
    N --> P[Close Return]
    O --> P
    
    style H fill:#64b5f6
    style I fill:#ef5350
    style P fill:#66bb6a
```

---

## System Architecture Flow

Diagram ini menunjukkan bagaimana data mengalir melalui sistem dari layer ke layer.

```mermaid
graph TB
    subgraph "Client Layer"
        A1[Web Dashboard<br/>Admin/Manager]
        A2[Mobile App<br/>Warehouse Staff]
        A3[API Clients<br/>External Systems]
    end
    
    subgraph "API Gateway Layer"
        B[Fiber HTTP Server]
        B1[Authentication Middleware]
        B2[RBAC Middleware]
        B3[Validation Middleware]
    end
    
    subgraph "Handler Layer"
        C1[Auth Handler]
        C2[Warehouse Handler]
        C3[Product Handler]
        C4[PO Handler]
        C5[SO Handler]
        C6[Inventory Handler]
    end
    
    subgraph "Service Layer - Business Logic"
        D1[Auth Service]
        D2[Warehouse Service]
        D3[Product Service]
        D4[Purchase Service]
        D5[Sales Service]
        D6[Inventory Service]
        D7[Picking Service]
        D8[Shipment Service]
    end
    
    subgraph "Repository Layer - Data Access"
        E1[User Repository]
        E2[Warehouse Repository]
        E3[Product Repository]
        E4[PO Repository]
        E5[SO Repository]
        E6[Inventory Repository]
    end
    
    subgraph "Database Layer"
        F1[(PostgreSQL<br/>Master Data)]
        F2[(PostgreSQL<br/>Transactional Data)]
        F3[(Redis<br/>Cache & Sessions)]
    end
    
    A1 --> B
    A2 --> B
    A3 --> B
    B --> B1
    B1 --> B2
    B2 --> B3
    B3 --> C1 & C2 & C3 & C4 & C5 & C6
    
    C1 --> D1
    C2 --> D2
    C3 --> D3
    C4 --> D4
    C5 --> D5
    C6 --> D6
    
    D4 --> D6
    D5 --> D6
    D5 --> D7
    D7 --> D8
    
    D1 --> E1
    D2 --> E2
    D3 --> E3
    D4 --> E4
    D5 --> E5
    D6 --> E6
    D7 --> E6
    
    E1 & E2 & E3 --> F1
    E4 & E5 & E6 --> F2
    D1 & D2 & D3 & D4 & D5 --> F3
    
    style B fill:#42a5f5
    style D1 fill:#66bb6a
    style D2 fill:#66bb6a
    style D3 fill:#66bb6a
    style D4 fill:#66bb6a
    style D5 fill:#66bb6a
    style D6 fill:#66bb6a
    style F1 fill:#ab47bc
    style F2 fill:#ab47bc
    style F3 fill:#ff7043
```

---

## Data Flow Examples

### Example 1: Complete Purchase Order to Stock Flow

```mermaid
graph LR
    A["1. Create PO<br/>(Draft)"] --> B["2. Submit PO<br/>(Submitted)"]
    B --> C["3. Approve PO<br/>(Approved)"]
    C --> D["4. Goods Receipt<br/>(Create GR)"]
    D --> E["5. QC Check<br/>(Pass/Fail)"]
    E --> F["6. Putaway Task<br/>(Generated)"]
    F --> G["7. Execute Putaway<br/>(Move to Location)"]
    G --> H["8. Update Inventory<br/>(Stock Available)"]
    
    style A fill:#e3f2fd
    style C fill:#c8e6c9
    style E fill:#fff9c4
    style H fill:#b2dfdb
```

### Example 2: Complete Sales Order to Delivery Flow

```mermaid
graph LR
    A["1. Create SO<br/>(New Order)"] --> B["2. Allocate Stock<br/>(Reserved)"]
    B --> C["3. Create Wave<br/>(Batched)"]
    C --> D["4. Picking Task<br/>(Assigned)"]
    D --> E["5. Execute Pick<br/>(Picked)"]
    E --> F["6. Packing<br/>(Packed)"]
    F --> G["7. Create Shipment<br/>(Ready)"]
    G --> H["8. Dispatch<br/>(In Transit)"]
    H --> I["9. Delivery<br/>(Completed)"]
    
    style A fill:#f3e5f5
    style B fill:#fff9c4
    style E fill:#c8e6c9
    style I fill:#b2dfdb
```

---

## Status Transitions

### Purchase Order Status Flow

```mermaid
stateDiagram-v2
    [*] --> DRAFT: Create PO
    DRAFT --> SUBMITTED: Submit
    SUBMITTED --> APPROVED: Approve
    SUBMITTED --> REJECTED: Reject
    APPROVED --> PARTIALLY_RECEIVED: Partial GR
    APPROVED --> RECEIVED: Full GR
    PARTIALLY_RECEIVED --> RECEIVED: Complete GR
    REJECTED --> [*]
    RECEIVED --> CLOSED: Close PO
    CLOSED --> [*]
    
    note right of DRAFT: Buyer can edit
    note right of SUBMITTED: Awaiting approval
    note right of APPROVED: Ready for receiving
    note right of RECEIVED: All items received
```

### Sales Order Status Flow

```mermaid
stateDiagram-v2
    [*] --> NEW: Create SO
    NEW --> ALLOCATED: Allocate Stock
    NEW --> BACKORDER: No Stock
    BACKORDER --> ALLOCATED: Stock Available
    ALLOCATED --> PICKING: Start Picking
    PICKING --> PICKED: Complete Picking
    PICKED --> PACKED: Complete Packing
    PACKED --> SHIPPED: Create Shipment
    SHIPPED --> DELIVERED: Delivery Confirmed
    DELIVERED --> COMPLETED: Close Order
    DELIVERED --> RETURNED: Customer Return
    RETURNED --> COMPLETED: Process Return
    COMPLETED --> [*]
    
    note right of NEW: Order received
    note right of ALLOCATED: Stock reserved
    note right of BACKORDER: Waiting for stock
    note right of DELIVERED: Customer received
```

### Inventory Movement Types

```mermaid
graph TB
    A[Inventory Movement Types] --> B[INBOUND]
    A --> C[OUTBOUND]
    A --> D[TRANSFER]
    A --> E[ADJUSTMENT]
    
    B --> B1[Purchase Receipt]
    B --> B2[Customer Return]
    B --> B3[Production Input]
    
    C --> C1[Sales Shipment]
    C --> C2[Supplier Return]
    C --> C3[Production Output]
    
    D --> D1[Location Transfer]
    D --> D2[Warehouse Transfer]
    
    E --> E1[Stock Count Adjustment]
    E --> E2[Damage Write-off]
    E --> E3[Expiry Write-off]
    E --> E4[Found Stock]
    
    style B fill:#c8e6c9
    style C fill:#ffcdd2
    style D fill:#bbdefb
    style E fill:#fff9c4
```

---

## Integration Points

```mermaid
graph TB
    subgraph "External Systems"
        A[E-commerce Platform]
        B[ERP System]
        C[Accounting System]
        D[Shipping Carriers API]
        E[Barcode Scanner Devices]
    end
    
    subgraph "WMS Core"
        F[API Gateway]
        G[Event Bus/Queue]
    end
    
    subgraph "WMS Modules"
        H[Sales Order Module]
        I[Purchase Order Module]
        J[Inventory Module]
        K[Shipment Module]
    end
    
    A -->|Create SO via API| F
    B -->|Sync Master Data| F
    C -->|Export Transactions| F
    D -->|Track Shipments| F
    E -->|Scan Events| F
    
    F --> G
    G --> H
    G --> I
    G --> J
    G --> K
    
    H -->|Update Order Status| A
    K -->|Tracking Updates| D
    J -->|Inventory Sync| B
    I -->|PO Confirmation| B
    
    style F fill:#42a5f5
    style G fill:#ff7043
```

---

## Summary

Dokumentasi ini mencakup:

✅ **End-to-End Business Flow** - Dari procurement hingga delivery  
✅ **Inbound Process** - PO, Goods Receipt, Putaway  
✅ **Outbound Process** - Sales Order, Picking, Packing, Shipment  
✅ **Inventory Management** - Stock Count, Adjustments  
✅ **Returns Management** - Customer & Supplier Returns  
✅ **System Architecture** - Layer-by-layer data flow  
✅ **Status Transitions** - State diagrams untuk setiap entity  
✅ **Integration Points** - External system connections

Semua diagram dapat di-render di Markdown viewer yang support Mermaid (GitHub, GitLab, VS Code, dll).
