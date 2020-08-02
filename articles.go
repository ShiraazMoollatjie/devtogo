package devtogo

import (
	"fmt"
)

// PublishedArticle returns a published article with post content for the provided article id.
// See https://docs.dev.to/api/#operation/getArticleById.
func (c *Client) PublishedArticle(id int32) (*Article, error) {
	var res Article
	err := c.get(c.baseURL+fmt.Sprintf("/articles/%d", id), &res)

	return &res, err
}

// PublishedArticleByPath returns a published article with post content for the provided article id.
// See https://docs.dev.to/api/#operation/getArticleById.
func (c *Client) PublishedArticleByPath(username, slug string) (*Article, error) {
	var res Article
	err := c.get(c.baseURL+fmt.Sprintf("/articles/%s/%s", username, slug), &res)

	return &res, err
}

// Articles returns a slice of articles according to https://docs.dev.to/api/#operation/getArticles.
func (c *Client) Articles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles?"+qp, &res)

	return res, err
}

// VideoArticles returns articles that are videos according to https://docs.dev.to/api/#operation/getArticlesWithVideo.
func (c *Client) VideoArticles(args Arguments) (VideoArticles, error) {
	var res VideoArticles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/videos?"+qp, &res)

	return res, err
}

// MyArticles returns a slice of articles according to https://docs.dev.to/api/#operation/getUserArticles.
func (c *Client) MyArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me?"+qp, &res)

	return res, err
}

// MyPublishedArticles returns a slice of published articles according to https://docs.dev.to/api/#operation/getUserPublishedArticles.
func (c *Client) MyPublishedArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/published?"+qp, &res)

	return res, err
}

// MyUnpublishedArticles returns a slice of unpublished articles according to https://docs.dev.to/api/#operation/getUserUnpublishedArticles.
func (c *Client) MyUnpublishedArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/unpublished?"+qp, &res)

	return res, err
}

// AllMyArticles returns a slice of unpublished articles according to https://docs.dev.to/api/#operation/getUserAllArticles.
func (c *Client) AllMyArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/all?"+qp, &res)

	return res, err
}

// CreateArticle creates a post with the provided request content on the dev.to platform according to https://docs.dev.to/api/#operation/createArticle.
func (c *Client) CreateArticle(req CreateArticleReq) (Article, error) {
	var res Article
	err := c.post(c.baseURL+"/articles", ArticleReq{Article: req}, &res)

	return res, err
}

// UpdateArticle will update a dev.to post with the provided ID and request content according to https://docs.dev.to/api/#operation/updateArticle.
func (c *Client) UpdateArticle(id int, req CreateArticleReq) (Article, error) {
	var res Article
	err := c.put(c.baseURL+fmt.Sprintf("/articles/%d", id), ArticleReq{Article: req}, &res)

	return res, err
}

// The structs in this file was generated via https://mholt.github.io/json-to-go/.

// ArticleReq is a container type to create articles.
type ArticleReq struct {
	Article CreateArticleReq `json:"article"`
}

// CreateArticleReq is a request struct that creates an article.
type CreateArticleReq struct {
	Title          string   `json:"title"`
	Published      bool     `json:"published"`
	BodyMarkdown   string   `json:"body_markdown"`
	Tags           []string `json:"tags"`
	Series         string   `json:"series,omitempty"`
	CanonicalURL   string   `json:"canonical_url"`
	MainImageURL   string   `json:"main_image"`
	Description    string   `json:"description"`
	OrganizationID int      `json:"organization_id"`
}

// Articles represents an article from the dev.to api.
type Articles []struct {
	Article
	TagList []string `json:"tag_list"`
	Tags    string   `json:"tags"`
}

// Article represents a single article in the dev.to api. Also has more fields than Articles.
type Article struct {
	TypeOf               string       `json:"type_of"`
	ID                   int          `json:"id"`
	Title                string       `json:"title"`
	Description          string       `json:"description"`
	CoverImage           string       `json:"cover_image"`
	ReadablePublishDate  string       `json:"readable_publish_date"`
	SocialImage          string       `json:"social_image"`
	TagList              []string     `json:"tags"`
	Tags                 string       `json:"tag_list"`
	Slug                 string       `json:"slug"`
	Path                 string       `json:"path"`
	URL                  string       `json:"url"`
	CanonicalURL         string       `json:"canonical_url"`
	CommentsCount        int          `json:"comments_count"`
	PublicReactionsCount int          `json:"public_reactions_count"`
	CollectionID         int          `json:"collection_id"`
	CreatedAt            emptyTime    `json:"created_at"`
	EditedAt             emptyTime    `json:"edited_at"`
	CrosspostedAt        emptyTime    `json:"crossposted_at"`
	PublishedAt          emptyTime    `json:"published_at"`
	LastCommentAt        emptyTime    `json:"last_comment_at"`
	PublishedTimestamp   emptyTime    `json:"published_timestamp"`
	BodyHTML             string       `json:"body_html"`
	BodyMarkdown         string       `json:"body_markdown"`
	User                 User         `json:"user"`
	Organization         Organization `json:"organization"`
	FlareTag             FlareTag     `json:"flare_tag"`
}

// Organization represents an organization from the dev.to api.
type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}

// FlareTag is a special colorized tag for the article.
type FlareTag struct {
	Name string `json:"name"`

	// BackgroundColorHex is a hexadecimal string value of the background color.
	BackgroundColorHex string `json:"bg_color_hex"`

	// TextColorHex is a hexadecimal string value of the background color.
	TextColorHex string `json:"text_color_hex"`
}

type VideoArticles []struct {
	TypeOf                 string `json:"type_of"`
	ID                     int    `json:"id"`
	Path                   string `json:"path"`
	CloudinaryVideoURL     string `json:"cloudinary_video_url"`
	Title                  string `json:"title"`
	UserID                 int    `json:"user_id"`
	VideoDurationInMinutes string `json:"video_duration_in_minutes"`
	User                   struct {
		Name string `json:"name"`
	} `json:"user"`
}
