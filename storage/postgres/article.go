package postgres

import (
	"UacademyGo/Article/models"
	"errors"
)

//*=========================================================================
func (stg Postgres) AddNewArticle(id string, box models.CreateModelArticle) error {
	var err error
	_, err = stg.GetAuthorById(box.AuthorID)
	if err != nil {
		return err
	}

	_, err = stg.homeDB.Exec(`INSERT INTO article 
	(
		id,
		title,
		body,
		author_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		id,
		box.Title,
		box.Body,
		box.AuthorID,
	)
	if err != nil {
		return err
	}
	return nil
}
//*=========================================================================
func (stg Postgres) GetArticleById(id string) (models.GetByIDArticleModel, error) {
	var a models.GetByIDArticleModel

	var tempMiddlename *string

	err := stg.homeDB.QueryRow(`SELECT 
		ar.id,
		ar.title,
		ar.body,
		ar.created_at,
		ar.updated_at,
		ar.deleted_at,
		au.id,
		au.fullname,
		au.middlename,
		au.created_at,
		au.updated_at,
		au.deleted_at
    FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
		&a.ID,
		&a.Title,
		&a.Body,
		&a.CreateAt,
		&a.UpdateAt,
		&a.DeletedAt,
		&a.Author.ID,
		&a.Author.Fullname,
		&tempMiddlename,
		&a.Author.CreateAt,
		&a.Author.UpdateAt,
		&a.Author.DeletedAt,
	)
	if err != nil {
		return a, err
	}

	if tempMiddlename == nil {
		a.Author.Middlename = ""
	} else {
		a.Author.Middlename = *tempMiddlename
	}

	// if tempMiddlename != nil {
	// 	a.Author.Middlename = *tempMiddlename
	// }

	return a, nil
}
//*=========================================================================
func (stg Postgres) GetArticleList(offset, limit int, search string) (dataset []models.Article, err error) {
	
	rows, err := stg.homeDB.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at,
	deleted_at 
	FROM article WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return dataset, err
	}

	for rows.Next() {
		var a models.Article

		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Body,
			&a.AuthorID,
			&a.CreateAt,
			&a.UpdateAt,
			&a.DeletedAt,
		)
		if err != nil {
			return dataset, err
		}
		dataset = append(dataset, a)
	}
	return dataset, err
}
//*=========================================================================
func (stg Postgres) UpdateArticle(box models.UpdateArticleResponse) error {
	res, err := stg.homeDB.NamedExec("UPDATE article  SET title=:t, body=:b, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": box.ID,
		"t":  box.Title,
		"b":  box.Body,
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
	return errors.New("article not found")
}
//*=========================================================================
func (stg Postgres) DeleteArticle(id string) error {
	res, err := stg.homeDB.Exec("UPDATE article  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
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
	return errors.New("article not found")
}
