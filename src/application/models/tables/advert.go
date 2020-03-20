package tables

type Advert struct {
	Id        int64  `xorm:"pk" xorm:"autoincr" json:"id" schema:"id"`
	Name      string `json:"name" schema:"name"`
	SpaceID   int64  `xorm:"space_id" json:"space_id" schema:"space_id"`
	SpaceName string `xorm:"-" json:"space_name"  schema:"-"`
	LinkUrl   string `json:"link_url" schema:"link_url"`
	Image     string `json:"image" schema:"image"`
	ListOrder uint   `xorm:"listorder" json:"listorder" schema:"listorder"`
	StartTime string `json:"start_time" schema:"-"`  // 展示时间周期开始
	EndTime   string `json:"end_time" schema:"-"`    // 展示时间周期结束
	Status    uint   `json:"status" schema:"status"` // 状态
}
