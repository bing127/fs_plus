package models

import "time"

type SysUser struct {
	ID			int		`gorm:"primary_key" json:" - "` //编号
	CompanyId	string	`json:"company_id"` //归属公司
	OfficeId	string	`json:"office_id"` //归属部门
	LoginName	string	`json:"login_name"` //登录名
	Password	string	`json:"password"` //密码
	No			string	`json:"no"` //工号
	Name		string	`json:"name"` //姓名
	Email		string	`json:"email"` //邮箱
	Phone		string	`json:"phone"` //电话
	Mobile		string	`json:"mobile"` //手机
	Photo		string	`json:"photo"` //用户头像
	LoginIp		string	`json:"login_ip"` //最后登陆IP
	LoginDate	time.Time	`json:"login_date"` //最后登陆时间
	LoginFlag	string	`json:"login_flag"` //是否可登录
	CreateBy	string	`json:"create_by"` //创建者
	CreateDate	time.Time	`json:"create_date"` //创建时间
	UpdateBy	string	`json:"update_by"` //更新者
	UpdateDate	time.Time	`json:"update_date"` //更新时间
	Remarks		string	`json:"remarks"` //备注信息
	DelFlag		string	`json:"del_flag"` //删除标记
	Qrcode		string	`json:"qrcode"` //二维码
	Sign		string	`json:"sign"` //个性签名
}
