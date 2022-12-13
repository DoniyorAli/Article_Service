package postgres

import (
	"UacademyGo/Blogpost/article_service/models"
	"errors"
)

func (stg Postgres) AddAuthor(id string, box models.CreateModelAuthor) error {
	_, err := stg.homeDB.Exec(`INSERT INTO author 
	(
		id,
		firstname,
		lastname,
		fullname,
		middlename
	) VALUES (
		$1,
		'',
		'',
		$2,
		$3
	)`,
		id,
		box.Fullname,
		box.Middlename,
	)
	if err != nil {
		return err
	}
	return nil
}

// *=========================================================================
func (stg Postgres) GetAuthorById(id string) (models.Author, error) {
	var author models.Author
	var tempMiddlename *string
	err := stg.homeDB.QueryRow(`SELECT 
		id,
		fullname,
		middlename,
		created_at,
		updated_at,
		deleted_at
    FROM author WHERE id = $1 AND deleted_at IS NULL`, id).Scan(
		&author.ID,
		&author.Fullname,
		&tempMiddlename,
		&author.CreateAt,
		&author.UpdateAt,
		&author.DeletedAt,
	)
	if err != nil {
		return author, err
	}

	if tempMiddlename != nil {
		author.Middlename = *tempMiddlename
	}

	return author, nil
}

// *=========================================================================
func (stg Postgres) GetAuthorList(offset, limit int, search string) ([]models.Author, error) {
	var res []models.Author
	var tempMiddlename *string
	rows, err := stg.homeDB.Queryx(`SELECT 
		id,
		fullname,
		middlename,
		created_at,
		updated_at,
		deleted_at 
		FROM author
		WHERE ((fullname ILIKE '%' || $1 || '%') or (middlename ILIKE '%' || $1 || '%')) AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3
	`,
		search,
		limit,
		offset,
	)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.ID,
			&author.Fullname,
			&tempMiddlename,
			&author.CreateAt,
			&author.UpdateAt,
			&author.DeletedAt,
		)
		if err != nil {
			return res, err
		}
		if tempMiddlename != nil {
			author.Middlename = *tempMiddlename
		}
		res = append(res, author)

		


	}
	return res, err
}

// *=========================================================================
func (stg Postgres) UpdateAuthor(box models.UpdateAuthorResponse) error {

	res, err := stg.homeDB.NamedExec("UPDATE author  SET fullname=:f, middlename=:m, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": box.ID,
		"f":  box.Fullname,
		"m":  box.Middlename,
	})
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("author not found")
}

// *=========================================================================
func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.homeDB.Exec("UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("author not found")
}
