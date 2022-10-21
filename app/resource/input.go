package resource

type NewUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"required"`
}

type NewPhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type NewSocialMedia struct {
	ID              uint   `json:"id"`
	Name            string `json:"name" binding:"required"`
	SocialMedialUrl string `json:"social_media_url" binding:"required"`
}

type EditSocialMedia struct {
	Name            string `json:"name" binding:"required"`
	SocialMedialUrl string `json:"social_media_url" binding:"required"`
}

type NewComment struct {
	ID      uint   `json:"id"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type EditComment struct {
	Message string `json:"message" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EditUser struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}
