package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type CbnDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateCbns(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS cbns(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Age TEXT NOT NULL,
		Name TEXT NOT NULL,
		Verified INTEGER NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewCbnDao() (*CbnDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateCbns(sqlClient)
	if err != nil {
		return nil, err
	}
	return &CbnDao{
		sqlClient,
	}, nil
}

func (cbnDao *CbnDao) CreateCbn(m *models.Cbn) (*models.Cbn, error) {
	insertQuery := "INSERT INTO cbns(Age, Name, Verified)values(?, ?, ?)"
	res, err := cbnDao.sqlClient.DB.Exec(insertQuery, m.Age, m.Name, m.Verified)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("cbn created")
	return m, nil
}

func (cbnDao *CbnDao) UpdateCbn(id int64, m *models.Cbn) (*models.Cbn, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	cbn, err := cbnDao.GetCbn(id)
	if err != nil {
		return nil, err
	}
	if cbn == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE cbns SET Age = ?, Name = ?, Verified = ? WHERE Id = ?"
	res, err := cbnDao.sqlClient.DB.Exec(updateQuery, m.Age, m.Name, m.Verified, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("cbn updated")
	return m, nil
}

func (cbnDao *CbnDao) DeleteCbn(id int64) error {
	deleteQuery := "DELETE FROM cbns WHERE Id = ?"
	res, err := cbnDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("cbn deleted")
	return nil
}

func (cbnDao *CbnDao) ListCbns() ([]*models.Cbn, error) {
	selectQuery := "SELECT * FROM cbns"
	rows, err := cbnDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var cbns []*models.Cbn
	for rows.Next() {
		m := models.Cbn{}
		if err = rows.Scan(&m.Id, &m.Age, &m.Name, &m.Verified); err != nil {
			return nil, err
		}
		cbns = append(cbns, &m)
	}
	if cbns == nil {
		cbns = []*models.Cbn{}
	}

	log.Debugf("cbn listed")
	return cbns, nil
}

func (cbnDao *CbnDao) GetCbn(id int64) (*models.Cbn, error) {
	selectQuery := "SELECT * FROM cbns WHERE Id = ?"
	row := cbnDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Cbn{}
	if err := row.Scan(&m.Id, &m.Age, &m.Name, &m.Verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("cbn retrieved")
	return &m, nil
}
