package main

type Service struct {
	db *DB
}

func NewService(db *DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetItem(itemID string) (*InventoryItem, error) {
	return s.db.GetItem(itemID)
}

func (s *Service) AddItem(item InventoryItem) (*InventoryItem, error) {
	_, newItem, err := s.db.AddItem(item)
	if err != nil {
		return nil, err
	}
	return newItem, nil
}

func (s *Service) UpdateItem(itemID string, updatedItem InventoryItem) error {
	return s.db.UpdateItem(itemID, updatedItem)
}

func (s *Service) DeleteItem(itemID string) error {
	return s.db.DeleteItem(itemID)
}

func (s *Service) GetInventoryByWarehouse(warehouseID string) ([]InventoryItem, error) {
	return s.db.GetInventoryByWarehouse(warehouseID)
}
