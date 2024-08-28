package orders

import (
	"gorm.io/gorm"
)

func RepoCreateOrder(db *gorm.DB, order NewOrder) (int64, error) {
	var resultId int64

}

// {success:true,id:5}
