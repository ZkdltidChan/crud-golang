repo, 數據庫處理, 查詢插入等, crud, 無業務邏輯
controllers, 業務邏輯, 提供數據
delivery, 如何呈現數據


controllers 需要 repo層, 透過interface 通信
ex: repo層內新增一個 interface
type aRepo interface {
    GetByID
    CreateA
    DeleteA
}

type aController interface {
    Get
}