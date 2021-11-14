package model

import "my-blog-sevice/pkg/app"

type Article struct {
	*Model
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`

	//Tag           *model.Tag `json:"tag"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (t Article) TableName() string {
	return "blog_article"
}
