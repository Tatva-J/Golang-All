// model.go

package main

import (
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// func (u *user) getUser(db *gorm.DB) error {
// 	// statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
// 	return db.Raw("SELECT name, age FROM users WHERE id=?", u.ID).Find(&u.Name, &u.Age).Error
// }

// func (u *user) updateUser(db *gorm.DB) error {
// 	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)
// 	_, err := db.Raw(statement)
// 	return err
// }

// func (u *user) deleteUser(db *gorm.DB) error {
// 	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
// 	_, err := db.Raw(statement)
// 	return err
// }

// func (u *user) createUser(db *gorm.DB) error {
// 	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", u.Name, u.Age)
// 	_, err := db.Raw(statement)

// 	if err != nil {
// 		return err
// 	}

// 	err = db.Raw("SELECT LAST_INSERT_ID()").Scan(&u.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func getUsers(db *gorm.DB, start, count int) ([]user, error) {
	// statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	if rows := db.Raw("SELECT id, name, age FROM users LIMIT ? OFFSET ?", count, start); rows == nil {

	} else {

	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
