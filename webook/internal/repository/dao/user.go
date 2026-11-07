package dao

import (
	"context"
	"errors"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱已被注册")
	ErrUserNotFound       = errors.New("用户不存在")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, ErrUserNotFound
	}
	return user, err
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err:=dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok {
		const uniqueConstraintErrorCode uint16 = 1062
		if mysqlErr.Number == uniqueConstraintErrorCode {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
}
