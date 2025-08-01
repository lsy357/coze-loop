// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameItemSnapshot = "dataset_item_snapshot"

// ItemSnapshot NDB_SHARE_TABLE;数据集条目快照
type ItemSnapshot struct {
	ID             int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                                                                                                                               // ID
	AppID          int32          `gorm:"column:app_id;type:bigint(20) unsigned;not null;comment:应用 ID" json:"app_id"`                                                                                                                                                         // 应用 ID
	SpaceID        int64          `gorm:"column:space_id;type:bigint(20) unsigned;not null;comment:空间 ID" json:"space_id"`                                                                                                                                                     // 空间 ID
	DatasetID      int64          `gorm:"column:dataset_id;type:bigint(20) unsigned;not null;comment:数据集 ID" json:"dataset_id"`                                                                                                                                                // 数据集 ID
	SchemaID       int64          `gorm:"column:schema_id;type:bigint(20) unsigned;not null;comment:Schema ID" json:"schema_id"`                                                                                                                                               // Schema ID
	VersionID      int64          `gorm:"column:version_id;type:bigint(20) unsigned;not null;uniqueIndex:uk_version_item,priority:1;index:idx_version_item_created_at_item,priority:1;index:idx_version_item_updated_at_item,priority:1;comment:Version ID" json:"version_id"` // Version ID
	ItemPrimaryID  int64          `gorm:"column:item_primary_id;type:bigint(20) unsigned;not null;comment:条目主键 ID" json:"item_primary_id"`                                                                                                                                     // 条目主键 ID
	ItemID         int64          `gorm:"column:item_id;type:bigint(20) unsigned;not null;uniqueIndex:uk_version_item,priority:2;index:idx_version_item_created_at_item,priority:3;index:idx_version_item_updated_at_item,priority:3;comment:条目 ID" json:"item_id"`            // 条目 ID
	ItemKey        string         `gorm:"column:item_key;type:varchar(128);not null;comment:条目幂等 key" json:"item_key"`                                                                                                                                                         // 条目幂等 key
	Data           datatypes.JSON `gorm:"column:data;type:json;comment:数据内容" json:"data"`                                                                                                                                                                                      // 数据内容
	RepeatedData   datatypes.JSON `gorm:"column:repeated_data;type:json;comment:多轮数据内容" json:"repeated_data"`                                                                                                                                                                  // 多轮数据内容
	DataProperties datatypes.JSON `gorm:"column:data_properties;type:json;comment:内容属性" json:"data_properties"`                                                                                                                                                                // 内容属性
	AddVn          int64          `gorm:"column:add_vn;type:bigint(20) unsigned;not null;comment:添加版本号" json:"add_vn"`                                                                                                                                                         // 添加版本号
	DelVn          int64          `gorm:"column:del_vn;type:bigint(20) unsigned;not null;comment:删除版本号" json:"del_vn"`                                                                                                                                                         // 删除版本号
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:snapshot 创建时间" json:"created_at"`                                                                                                                         // snapshot 创建时间
	ItemCreatedBy  string         `gorm:"column:item_created_by;type:varchar(128);not null;comment:item 创建人" json:"item_created_by"`                                                                                                                                           // item 创建人
	ItemCreatedAt  time.Time      `gorm:"column:item_created_at;type:timestamp;not null;index:idx_version_item_created_at_item,priority:2;default:CURRENT_TIMESTAMP;comment:item 创建时间" json:"item_created_at"`                                                                 // item 创建时间
	ItemUpdatedBy  string         `gorm:"column:item_updated_by;type:varchar(128);not null;comment:item 修改人" json:"item_updated_by"`                                                                                                                                           // item 修改人
	ItemUpdatedAt  time.Time      `gorm:"column:item_updated_at;type:timestamp;not null;index:idx_version_item_updated_at_item,priority:2;default:CURRENT_TIMESTAMP;comment:item 修改时间" json:"item_updated_at"`                                                                 // item 修改时间
}

// TableName ItemSnapshot's table name
func (*ItemSnapshot) TableName() string {
	return TableNameItemSnapshot
}
