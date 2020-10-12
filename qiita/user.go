package qiita

type User struct {
	Description       *string `json:"description"`
	FacebookID        *string `json:"facebook_id"`
	FolloweesCount    int     `json:"followees_count"`
	FollowersCount    int     `json:"followers_count"`
	GithubLoginName   *string `json:"github_login_name"`
	ID                string  `json:"id"`
	ItemsCount        int     `json:"items_count"`
	LinkedinID        *string `json:"linkedin_id"`
	Location          *string `json:"location"`
	Name              *string `json:"name"`
	Organization      *string `json:"organization"`
	PermanetID        int     `json:"permanent_id"`
	ProfileImageURL   string  `json:"profile_image_url"`
	TeamOnly          bool    `json:"team_only"`
	TwitterScreenName *string `json:"twitter_screen_name"`
	WebsiteURL        *string `json:"website_url"`
}
