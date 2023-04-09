package isql

import (
	"errors"
	"github.com/yin-zt/itsm-workflow/pkg/models/order"
	"github.com/yin-zt/itsm-workflow/pkg/utils/common"
	"gorm.io/gorm"
)

type OrderService struct{}

// List 获取数据列表
//func (s OrderService) List(req *request.GroupListReq) ([]*order.Order, error) {
//	var list []*order.Group
//	db := common.DB.Model(&order.Order{}).Order("created_at DESC")
//
//	groupName := strings.TrimSpace(req.GroupName)
//	if groupName != "" {
//		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
//	}
//	groupRemark := strings.TrimSpace(req.Remark)
//	if groupRemark != "" {
//		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
//	}
//
//	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
//	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Users").Find(&list).Error
//	return list, err
//}

// List 获取数据列表
//func (s OrderService) ListAll(req *request.GroupListAllReq) ([]*model.Group, error) {
//	var list []*model.Group
//	db := common.DB.Model(&model.Group{}).Order("created_at DESC")
//
//	groupName := strings.TrimSpace(req.GroupName)
//	if groupName != "" {
//		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
//	}
//	groupRemark := strings.TrimSpace(req.Remark)
//	if groupRemark != "" {
//		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
//	}
//	groupType := strings.TrimSpace(req.GroupType)
//	if groupType != "" {
//		db = db.Where("group_type = ?", groupType)
//	}
//	source := strings.TrimSpace(req.Source)
//	if source != "" {
//		db = db.Where("source = ?", source)
//	}
//	sourceDeptId := strings.TrimSpace(req.SourceDeptId)
//	if sourceDeptId != "" {
//		db = db.Where("source_dept_id = ?", sourceDeptId)
//	}
//	sourceDeptParentId := strings.TrimSpace(req.SourceDeptParentId)
//	if sourceDeptParentId != "" {
//		db = db.Where("source_dept_parent_id = ?", sourceDeptParentId)
//	}
//
//	err := db.Find(&list).Error
//	return list, err
//}

// Count 获取数据总数
func (s OrderService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&order.Order{}).Count(&count).Error
	return count, err
}

// Add 添加资源
func (s OrderService) Add(data *order.Order) error {
	return common.DB.Create(data).Error
}

// Update 更新资源
func (s OrderService) Update(dataObj *order.Order) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s OrderService) Find(filter map[string]interface{}, data *order.Order, args ...interface{}) error {
	return common.DB.Where(filter, args).Preload("Users").First(&data).Error
}

// Exist 判断资源是否存在
func (s OrderService) Exist(filter map[string]interface{}) bool {
	var dataObj order.Order
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s OrderService) Delete(groups []*order.Order) error {
	return common.DB.Debug().Select("Users").Unscoped().Delete(&groups).Error
}

// GetApisById 根据接口ID获取接口列表
func (s OrderService) GetGroupByIds(ids []uint) (datas []*order.Order, err error) {
	err = common.DB.Where("id IN (?)", ids).Preload("Users").Find(&datas).Error
	return datas, err
}
