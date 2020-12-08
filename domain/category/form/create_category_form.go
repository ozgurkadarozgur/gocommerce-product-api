package form

type CreateCategoryForm struct {
	Title string `form:"title" json:"title" binding:"required"`
}