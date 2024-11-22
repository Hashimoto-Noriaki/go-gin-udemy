package repositories

type IAuthRepository interface {
	CreateUser(user models.User) error
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB)  IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository)  CreateUser(user models.user) error {
	result :=  r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}