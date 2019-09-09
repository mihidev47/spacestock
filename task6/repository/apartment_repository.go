package repository

import (
	"github.com/jmoiron/sqlx"
	"../model"
) 


type AptRepository interface {
	GetAll() ([]model.Apartment, error)
	FindById(id string) (*model.Apartment, error)
	Insert(newApt *model.Apartment) (int64, error)
	Update(Id string, upApt *model.Apartment) error
	Delete(Id string) error
	
}

type aptRepository struct {
	findByIdStmt        *sqlx.Stmt
	getAllStmt        	*sqlx.Stmt
	deleteByIdStmt        *sqlx.Stmt
}

func initAptRepository() AptRepository {
	// Init repository
	var r aptRepository
	// Init statements
	r.getAllStmt = MustPrepare(db, `SELECT * FROM apartment`)
	r.findByIdStmt = MustPrepare(db, `SELECT * FROM apartment WHERE id = ?`)
	r.deleteByIdStmt = MustPrepare(db, "DELETE FROM apartment WHERE id = ? ")
	// Return instance
	return &r
}

func (f *aptRepository) FindById(id string) (*model.Apartment, error) {
	var p model.Apartment
	err := f.findByIdStmt.Get(&p, id)
	return &p, err
}

func (f *aptRepository) GetAll() ([]model.Apartment, error) {
	var p []model.Apartment
	err := f.getAllStmt.Select(&p)
	return p, err
}

func (f *aptRepository) Delete(id string) (err error) {
	// Delete 
	_, err = f.deleteByIdStmt.Exec(id)
	return
}

func (f *aptRepository) Insert(newApt *model.Apartment) (int64, error) {
	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	// Release transaction on return func
	defer ReleaseTx(tx, &err)
	
	// Insert into foody table
	data, err := tx.Exec("INSERT INTO `apartment` (`name`, `address`) VALUES (?, ?)",
		newApt.Name, newApt.Address)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	idRow, _ := data.LastInsertId()

	return idRow, nil
}

func (f *aptRepository) Update(Id string, upApt *model.Apartment) error {
	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		log.Error(err)
		return err
	}
	// Release transaction on return func
	defer ReleaseTx(tx, &err)
	
	// Insert into foody table
	_, err = tx.Exec("UPDATE `apartment` SET `name` = ?, `address` = ?  WHERE `id` = ?",
		upApt.Name, upApt.Address, Id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}