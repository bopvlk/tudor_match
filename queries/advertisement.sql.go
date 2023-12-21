// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: advertisement.sql

package queries

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAdvertisement = `-- name: CreateAdvertisement :one
WITH category_id AS (SELECT id FROM categories WHERE name = $6)
INSERT INTO advertisements (
  title,
  provider,
  provider_id,
  attachment,
  experience,
  category_id,
  time,
  price,
  format,
  language,
  description,
  mobile_phone,
  email,
  telegram,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at
`

type CreateAdvertisementParams struct {
	Title       string      `json:"title"`
	Provider    string      `json:"provider"`
	ProviderID  int64       `json:"provider_id"`
	Attachment  pgtype.Text `json:"attachment"`
	Experience  int32       `json:"experience"`
	CategoryID  int64       `json:"category_id"`
	Time        int32       `json:"time"`
	Price       int32       `json:"price"`
	Format      string      `json:"format"`
	Language    string      `json:"language"`
	Description string      `json:"description"`
	MobilePhone string      `json:"mobile_phone"`
	Email       string      `json:"email"`
	Telegram    string      `json:"telegram"`
	CreatedAt   time.Time   `json:"created_at"`
}

func (q *Queries) CreateAdvertisement(ctx context.Context, arg CreateAdvertisementParams) (Advertisement, error) {
	row := q.db.QueryRow(ctx, createAdvertisement,
		arg.Title,
		arg.Provider,
		arg.ProviderID,
		arg.Attachment,
		arg.Experience,
		arg.CategoryID,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Email,
		arg.Telegram,
		arg.CreatedAt,
	)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.CategoryID,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAdvertisementByID = `-- name: DeleteAdvertisementByID :exec
DELETE FROM advertisements
WHERE id = $1
`

func (q *Queries) DeleteAdvertisementByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByID, id)
	return err
}

const deleteAdvertisementByUserID = `-- name: DeleteAdvertisementByUserID :exec
DELETE FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) DeleteAdvertisementByUserID(ctx context.Context, providerID int64) error {
	_, err := q.db.Exec(ctx, deleteAdvertisementByUserID, providerID)
	return err
}

const filterAdvertisements = `-- name: FilterAdvertisements :many
WITH filtered_ads AS (
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
  WHERE
        (NULLIF($5::text, '')::text IS NULL OR category = $5::text)
        AND (NULLIF($6::int, 0) IS NULL OR time <= $6::int)
        AND (NULLIF($7::text, '') IS NULL OR format = $8::text)
        AND ((NULLIF($9::int, 0) IS NULL AND NULLIF($10::int, 0) IS NULL) OR (experience >= $9::int AND experience <= $10::int))
        AND ((NULLIF($11::int, 0) IS NULL AND NULLIF($12::int, 0) IS NULL) OR (price >= $11::int AND price <= $12::int))
        AND (NULLIF($13::text, '') IS NULL OR language = $13::text)
        AND (NULLIF($14::text, '') IS NULL OR title ILIKE '%' || $14::text || '%')
)
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at,
    COUNT(*) OVER () AS total_items
FROM filtered_ads
ORDER BY
  ( CASE
    WHEN $1::text = 'price' AND $2::text = 'desc' THEN CAST(price AS TEXT)
    WHEN $1::text = 'experience' AND $2::text = 'desc' THEN CAST(experience AS TEXT)
    WHEN $1::text = 'date' AND $2::text = 'desc' THEN CAST(created_at AS TEXT) END) DESC,
  ( CASE
    WHEN $1::text = 'price' THEN CAST(price AS TEXT)
    WHEN $1::text = 'experience' THEN CAST(experience AS TEXT)  
    ELSE CAST(created_at AS TEXT) END) ASC                                     
LIMIT $4::integer    
OFFSET $3::integer
`

type FilterAdvertisementsParams struct {
	Orderby      string `json:"orderby"`
	Sortorder    string `json:"sortorder"`
	Offsetadv    int32  `json:"offsetadv"`
	Limitadv     int32  `json:"limitadv"`
	Advcategory  string `json:"advcategory"`
	Timelength   int32  `json:"timelength"`
	Advfformat   string `json:"advfformat"`
	Advformat    string `json:"advformat"`
	Minexp       int32  `json:"minexp"`
	Maxexp       int32  `json:"maxexp"`
	Minprice     int32  `json:"minprice"`
	Maxprice     int32  `json:"maxprice"`
	Advlanguage  string `json:"advlanguage"`
	Titlekeyword string `json:"titlekeyword"`
}

type FilterAdvertisementsRow struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Provider    string      `json:"provider"`
	ProviderID  int64       `json:"provider_id"`
	Attachment  pgtype.Text `json:"attachment"`
	Experience  int32       `json:"experience"`
	CategoryID  int64       `json:"category_id"`
	Time        int32       `json:"time"`
	Price       int32       `json:"price"`
	Format      string      `json:"format"`
	Language    string      `json:"language"`
	Description string      `json:"description"`
	MobilePhone string      `json:"mobile_phone"`
	Email       string      `json:"email"`
	Telegram    string      `json:"telegram"`
	CreatedAt   time.Time   `json:"created_at"`
	TotalItems  int64       `json:"total_items"`
}

func (q *Queries) FilterAdvertisements(ctx context.Context, arg FilterAdvertisementsParams) ([]FilterAdvertisementsRow, error) {
	rows, err := q.db.Query(ctx, filterAdvertisements,
		arg.Orderby,
		arg.Sortorder,
		arg.Offsetadv,
		arg.Limitadv,
		arg.Advcategory,
		arg.Timelength,
		arg.Advfformat,
		arg.Advformat,
		arg.Minexp,
		arg.Maxexp,
		arg.Minprice,
		arg.Maxprice,
		arg.Advlanguage,
		arg.Titlekeyword,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterAdvertisementsRow
	for rows.Next() {
		var i FilterAdvertisementsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
			&i.TotalItems,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementAll = `-- name: GetAdvertisementAll :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
ORDER BY created_at DESC
LIMIT 10
`

func (q *Queries) GetAdvertisementAll(ctx context.Context) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByCategory = `-- name: GetAdvertisementByCategory :many
WITH id AS (SELECT id FROM categories WHERE name = $1)
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE category_id = id
`

func (q *Queries) GetAdvertisementByCategory(ctx context.Context, name string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByCategory, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByExperience = `-- name: GetAdvertisementByExperience :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE experience >= $1
AND experience <= $2
`

type GetAdvertisementByExperienceParams struct {
	Experience   int32 `json:"experience"`
	Experience_2 int32 `json:"experience_2"`
}

func (q *Queries) GetAdvertisementByExperience(ctx context.Context, arg GetAdvertisementByExperienceParams) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByExperience, arg.Experience, arg.Experience_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByFormat = `-- name: GetAdvertisementByFormat :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE format = $1
`

func (q *Queries) GetAdvertisementByFormat(ctx context.Context, format string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByFormat, format)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByID = `-- name: GetAdvertisementByID :one
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE id = $1
`

func (q *Queries) GetAdvertisementByID(ctx context.Context, id int64) (Advertisement, error) {
	row := q.db.QueryRow(ctx, getAdvertisementByID, id)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.CategoryID,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}

const getAdvertisementByLanguage = `-- name: GetAdvertisementByLanguage :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE language = $1
`

func (q *Queries) GetAdvertisementByLanguage(ctx context.Context, language string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByLanguage, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByTime = `-- name: GetAdvertisementByTime :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE time <= $1
`

func (q *Queries) GetAdvertisementByTime(ctx context.Context, argTime int32) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByTime, argTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByUserID = `-- name: GetAdvertisementByUserID :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE provider_id = $1
`

func (q *Queries) GetAdvertisementByUserID(ctx context.Context, providerID int64) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByUserID, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementByUsername = `-- name: GetAdvertisementByUsername :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements
WHERE provider = $1
`

func (q *Queries) GetAdvertisementByUsername(ctx context.Context, provider string) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementByUsername, provider)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAdvertisementMy = `-- name: GetAdvertisementMy :many
SELECT id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at FROM advertisements 
WHERE provider_id = $1
`

func (q *Queries) GetAdvertisementMy(ctx context.Context, providerID int64) ([]Advertisement, error) {
	rows, err := q.db.Query(ctx, getAdvertisementMy, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertisement
	for rows.Next() {
		var i Advertisement
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Provider,
			&i.ProviderID,
			&i.Attachment,
			&i.Experience,
			&i.CategoryID,
			&i.Time,
			&i.Price,
			&i.Format,
			&i.Language,
			&i.Description,
			&i.MobilePhone,
			&i.Email,
			&i.Telegram,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdvertisement = `-- name: UpdateAdvertisement :one
WITH c_id AS (SELECT id FROM categories WHERE name = $6)
UPDATE advertisements
SET
  title = COALESCE($2, title),
  created_at = COALESCE($3, created_at),
  attachment = COALESCE($4, attachment),
  experience = COALESCE($5, experience),
  category_id = COALESCE($6, c_id),
  time = COALESCE($7, time),
  price = COALESCE($8, price),
  format = COALESCE($9, format),
  language = COALESCE($10, language),
  description = COALESCE($11, description),
  mobile_phone = COALESCE($12, mobile_phone),
  telegram = COALESCE($13, telegram)
WHERE advertisements.id = $1
RETURNING id, title, provider, provider_id, attachment, experience, category_id, time, price, format, language, description, mobile_phone, email, telegram, created_at
`

type UpdateAdvertisementParams struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	CreatedAt   time.Time   `json:"created_at"`
	Attachment  pgtype.Text `json:"attachment"`
	Experience  int32       `json:"experience"`
	CategoryID  int64       `json:"category_id"`
	Time        int32       `json:"time"`
	Price       int32       `json:"price"`
	Format      string      `json:"format"`
	Language    string      `json:"language"`
	Description string      `json:"description"`
	MobilePhone string      `json:"mobile_phone"`
	Telegram    string      `json:"telegram"`
}

func (q *Queries) UpdateAdvertisement(ctx context.Context, arg UpdateAdvertisementParams) (Advertisement, error) {
	row := q.db.QueryRow(ctx, updateAdvertisement,
		arg.ID,
		arg.Title,
		arg.CreatedAt,
		arg.Attachment,
		arg.Experience,
		arg.CategoryID,
		arg.Time,
		arg.Price,
		arg.Format,
		arg.Language,
		arg.Description,
		arg.MobilePhone,
		arg.Telegram,
	)
	var i Advertisement
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Provider,
		&i.ProviderID,
		&i.Attachment,
		&i.Experience,
		&i.CategoryID,
		&i.Time,
		&i.Price,
		&i.Format,
		&i.Language,
		&i.Description,
		&i.MobilePhone,
		&i.Email,
		&i.Telegram,
		&i.CreatedAt,
	)
	return i, err
}
