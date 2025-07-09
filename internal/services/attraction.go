package services

import (
	"fmt"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
)

type attractionService struct {
	repo interfaces.AttractionRepository // 修正為 interfaces.AttractionRepository
}

func NewAttractionService(repo interfaces.AttractionRepository) interfaces.AttractionService {
	return &attractionService{
		repo: repo,
	}
}

func (s *attractionService) GetAttractionByID(id int) (*models.Attraction, error) {
	return s.repo.GetByID(uint(id)) // 確保傳入的 ID 類型正確
}

func (s *attractionService) ListAttractions(page, limit int) ([]models.Attraction, error) {
	attractions, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, fmt.Errorf("取得景點列表失敗: %w", err)
	}
	return attractions, nil
}

func (s *attractionService) SearchAttractions(keyword string) ([]models.Attraction, error) {
	attractions, err := s.repo.Search(keyword)
	if err != nil {
		return nil, fmt.Errorf("搜尋景點失敗: %w", err)
	}
	return attractions, nil
}
func (s *attractionService) GetMRTNames() ([]string, error) {
	return s.repo.GetMRTNames()
}

// func (s *attractionService) GetMRTsWithAttractionCount() ([]struct {
// 	MRT   string
// 	Count int
// }, error) {
// 	return s.repo.GetMRTsWithAttractionCount()
// }

// SearchAttractionsByKeyword 依關鍵字搜尋景點
func (s *attractionService) SearchAttractionsByKeyword(keyword string) ([]models.Attraction, error) {
	// 呼叫 repository 層的全文搜尋方法
	return s.repo.SearchAttractionsByKeyword(keyword)
}
