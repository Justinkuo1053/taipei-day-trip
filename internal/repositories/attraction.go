package repositories

import (
	"taipei-day-trip-go-go/internal/models"

	"gorm.io/gorm"
)

type AttractionRepository struct {
	// 通常會包含資料庫連線或其他依賴
	DB *gorm.DB
}

// NewAttractionRepository 用來實體化 AttractionRepository
func NewAttractionRepository(db *gorm.DB) *AttractionRepository {
	return &AttractionRepository{DB: db}
}

func (r *AttractionRepository) GetAllAttractions() ([]models.Attraction, error) {
	// TODO: implement
	var attractions []models.Attraction
	return attractions, nil
}

func (r *AttractionRepository) GetByID(id uint) (*models.Attraction, error) {
	var attraction models.Attraction
	if err := r.DB.First(&attraction, id).Error; err != nil {
		return nil, err
	}
	// 查詢 images
	var images []models.Image
	// 查詢景點時，會額外查詢 images 資料表，取得該景點所有圖片網址，組成 imageURLs 陣列，並賦值給 Attraction.Images
	r.DB.Where("attraction_id = ?", id).Find(&images)
	imageURLs := make([]string, 0, len(images))
	for _, img := range images {
		imageURLs = append(imageURLs, img.URL)
	}
	attraction.Images = imageURLs
	return &attraction, nil
}

func (r *AttractionRepository) CreateAttraction(attraction *models.Attraction) error {
	// TODO: implement
	return nil
}

func (r *AttractionRepository) GetAll(page, limit int) ([]models.Attraction, error) {
	var attractions []models.Attraction
	offset := (page - 1) * limit
	result := r.DB.Limit(limit).Offset(offset).Find(&attractions)
	if result.Error != nil {
		return nil, result.Error
	}
	// 針對每個景點查詢 images，組成 Images 陣列
	for i := range attractions {
		var images []models.Image
		r.DB.Where("attraction_id = ?", attractions[i].ID).Find(&images)
		imageURLs := make([]string, 0, len(images))
		for _, img := range images {
			imageURLs = append(imageURLs, img.URL)
		}
		attractions[i].Images = imageURLs
	}
	return attractions, nil
}

func (r *AttractionRepository) Search(keyword string) ([]models.Attraction, error) {
	var attractions []models.Attraction
	// ...查詢資料庫...
	return attractions, nil
}

// GetMRTNames 取得所有捷運站名稱
func (r *AttractionRepository) GetMRTNames() ([]string, error) {
	var results []string
	// group by mrt, count attractions, order by count desc
	err := r.DB.Model(&models.Attraction{}).
		Select("mrt").
		Where("mrt != ''"). // 避免空白捷運站
		Group("mrt").
		Order("COUNT(*) DESC").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// // GetMRTsWithAttractionCount 取得所有捷運站名稱及其景點數量，並依數量排序
// func (r *AttractionRepository) GetMRTsWithAttractionCount() ([]struct {
// 	MRT   string
// 	Count int
// }, error) {
// 	var results []struct {
// 		MRT   string
// 		Count int
// 	}
// 	// group by mrt, count attractions, order by count desc
// 	err := r.DB.Model(&models.Attraction{}).
// 		Select("mrt, COUNT(*) as count").
// 		Where("mrt != ''"). // 避免空白捷運站
// 		Group("mrt").
// 		Order("count DESC").
// 		Scan(&results).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }

// SearchAttractionsByKeyword 使用 LIKE 模糊查詢，解決中文全文搜尋問題
// keyword: 使用者輸入的關鍵字，可以是景點名稱、捷運站名稱、描述等
// 回傳符合條件的景點陣列
func (r *AttractionRepository) SearchAttractionsByKeyword(keyword string) ([]models.Attraction, error) {
	var attractions []models.Attraction
	// 改用 LIKE 查詢，支援中文子字串模糊搜尋
	likePattern := "%" + keyword + "%"
	result := r.DB.Raw(`
		SELECT * FROM attractions
		WHERE name LIKE ? OR description LIKE ? OR mrt LIKE ?
	`, likePattern, likePattern, likePattern).Scan(&attractions)
	if result.Error != nil {
		return nil, result.Error
	}
	// 查詢每個景點的圖片，組成 Images 陣列
	for i := range attractions {
		var images []models.Image
		r.DB.Where("attraction_id = ?", attractions[i].ID).Find(&images)
		imageURLs := make([]string, 0, len(images))
		for _, img := range images {
			imageURLs = append(imageURLs, img.URL)
		}
		attractions[i].Images = imageURLs
	}
	return attractions, nil
}
