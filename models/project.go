package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Project struct {
	Id             int       `gorm:"column(id);auto"`
	UserId         uint      `gorm:"column(user_id)"`
	Name           string    `gorm:"column(name);size(100);null"`
	Level          int16     `gorm:"column(level)"`
	Status         int16     `gorm:"column(status)"`
	Version        string    `gorm:"column(version);size(32);null"`
	RepoUrl        string    `gorm:"column(repo_url);type(text);null"`
	RepoUsername   string    `gorm:"column(repo_username);size(50);null"`
	RepoPassword   string    `gorm:"column(repo_password);size(100);null"`
	RepoMode       string    `gorm:"column(repo_mode);size(50);null"`
	RepoType       string    `gorm:"column(repo_type);size(10);null"`
	DeployFrom     string    `gorm:"column(deploy_from);size(200)"`
	Excludes       string    `gorm:"column(excludes);type(text);null"`
	ReleaseUser    string    `gorm:"column(release_user);size(50)"`
	ReleaseTo      string    `gorm:"column(release_to);size(200)"`
	ReleaseLibrary string    `gorm:"column(release_library);type(text);size(200)"`
	Hosts          string    `gorm:"column(hosts);type(text);null"`
	PreDeploy      string    `gorm:"column(pre_deploy);type(text);null"`
	PostDeploy     string    `gorm:"column(post_deploy);type(text);null"`
	PreRelease     string    `gorm:"column(pre_release);type(text);null"`
	PostRelease    string    `gorm:"column(post_release);type(text);null"`
	LastDeploy     string    `gorm:"column(last_deploy);type(text);null"`
	Audit          int16     `gorm:"column(audit);null"`
	KeepVersionNum int       `gorm:"column(keep_version_num)"`
	CreatedAt      time.Time `gorm:"column(created_at);type(datetime);null"`
	UpdatedAt      time.Time `gorm:"column(updated_at);type(datetime);null"`
	P2p            int16     `gorm:"column(p2p)"`
	HostGroup      string    `gorm:"column(host_group)"`
	Gzip           int16     `gorm:"column(gzip)"`
	IsGroup        int16     `gorm:"column(is_group)"`
	PmsProName     string    `gorm:"column(pms_pro_name);size(200)"`
}

func (u *Project) TableName() string {
	return "project"
}

func (u *Project) Save() error {
	return nil
}

func (u *Project) Create() (*Project, error) {
	return nil, nil
}

func FindProjects(where string, start, length int) (*[]Project, error) {
	var projects Project
	globalDB.Raw("SELECT *, (SELECT realname FROM `user` WHERE `user`.id=project.user_id LIMIT 1) as realname FROM `project`  WHERE 1=1 "+where+" ORDER BY id LIMIT ?,?", start, length).Scan(&projects)
	return nil, nil
}
