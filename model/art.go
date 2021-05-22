package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Art 芸術情報
type Art struct {
	ArtName              string    `db:"art_name"`
	ArtNameEn            string    `db:"art_name_en"`
	YearOfProductionFrom time.Time `db:"year_of_production_from"`
	YearOfProductionTo   time.Time `db:"year_of_production_to"`
	Evaluation           float32   `db:"evaluation"`
	Height               float32   `db:"height"`
	Width                float32   `db:"width"`
	FileUrl              string    `db:"file_url"`
	AuthorName           string    `db:"author_name"`
	SuppliesName         string    `db:"supplies_name"`
	CollectionName       string    `db:"collection_name"`
	PaintName            string    `db:"paint_name"`
}

// GetArts 芸術情報を検索します
func GetArts(tx *gorp.Transaction) ([]Art, error) {

	// 企業情報を検索
	arts, err := selectToArt(tx)
	if err != nil {
		return nil, err
	}

	return arts, nil
}

// 芸術情報を検索します
func selectToArt(tx *gorp.Transaction) ([]Art, error) {
	var arts []Art
	_, err := tx.Select(&arts, `
select 
  art.name art_name, 
  art.name_en art_name_en, 
  art.year_of_production_from, 
  art.year_of_production_to, 
  art.evaluation, 
  art.height, 
  art.width, 
  art.file_url, 
  author.name author_name, 
  supplies.name supplies_name, 
  collection.name collection_name, 
  paint.name paint_name 
from 
  art 
  inner join author on art.author_id = author.author_id 
  inner join art_supplies supplies on art.art_supplies_id = supplies.art_supplies_id 
  inner join collection on art.collection_id = collection.collection_id 
  inner join paint on art.paint_id = paint.paint_id 
order by
  art.id
		`)
	if err != nil {
		return arts, err
	}

	return arts, nil
}
