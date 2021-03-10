package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Battery   string `gorm:"size:255;not null;unique" json:"battery"`
	VoltageL1 string `gorm:"size:255;not null;" json:"voltage_l1"`
	VoltageL2 string `gorm:"size:255;not null;" json:"voltage_l2"`
	VoltageL3 string `gorm:"size:255;not null;" json:"voltage_l3"`
	CurrentL1 string `gorm:"size:255;not null;" json:"current_l1"`
	CurrentL2 string `gorm:"size:255;not null;" json:"current_l2"`
	CurrentL3 string `gorm:"size:255;not null;" json:"current_l3"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Battery = html.EscapeString(strings.TrimSpace(p.Battery))
	p.VoltageL1 = html.EscapeString(strings.TrimSpace(p.VoltageL1))
	p.VoltageL2 = html.EscapeString(strings.TrimSpace(p.VoltageL2))
	p.VoltageL3 = html.EscapeString(strings.TrimSpace(p.VoltageL3))
	p.CurrentL1 = html.EscapeString(strings.TrimSpace(p.CurrentL1))
	p.CurrentL2 = html.EscapeString(strings.TrimSpace(p.CurrentL2))
	p.CurrentL3 = html.EscapeString(strings.TrimSpace(p.CurrentL3))
}

func (p *Post) Validate() error {
	if p.Battery == "" {
		return errors.New("Required Battery")
	}
	if p.VoltageL1 == "" {
		return errors.New("Required VoltageL1")
	}
	if p.VoltageL2 == "" {
		return errors.New("Required VoltageL2")
	}
	if p.VoltageL3 == "" {
		return errors.New("Required VoltageL3")
	}
	if p.CurrentL1 == "" {
		return errors.New("Required CurrentL1")
	}
	if p.CurrentL2 == "" {
		return errors.New("Required CurrentL2")
	}
	if p.CurrentL3 == "" {
		return errors.New("Required CurrentL3")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&Data{}).Where("id = ?", p.AuthorID).Take(&p.AuthorID).Error
	// 	if err != nil {
	// 		return &Post{}, err
	// 	}
	// }
	return p, nil
}

func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	posts := []Post{}
	err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &posts, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {

	var err error

	err = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Updates(Post{Battery: p.Battery, VoltageL1: p.VoltageL1, VoltageL2: p.VoltageL2, VoltageL3: p.VoltageL3, CurrentL1: p.CurrentL1, CurrentL2: p.CurrentL2, CurrentL3: p.CurrentL3}).Error
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
