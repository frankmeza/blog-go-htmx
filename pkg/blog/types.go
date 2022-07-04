package blog

import (
	"html/template"
	"sort"

	"github.com/frankmeza/blog-go-htmx/pkg/utils"
)

const (
	DOT_MD                 string = ".md"
	ERROR_IN_FUNC          string = "error in func "
	MARKDOWN_PAGES_PATH    string = "pkg/markdown/pages/"
	MARKDOWN_POSTS_PATH    string = "pkg/markdown/posts/"
	TEMPLATE_HOME          string = "home"
	TEMPLATE_MARKDOWN_PAGE string = "markdownPage"
)

type PostMetaData struct {
	IsDraft bool     `yaml:"is_draft" toml:"is_draft" json:"is_draft"`
	Key     string   `yaml:"key" toml:"key" json:"key"`
	Summary string   `yaml:"summary" toml:"summary" json:"summary"`
	Tags    []string `yaml:"tags" toml:"tags" json:"tags"`
	Title   string   `yaml:"title" toml:"title" json:"title"`
	Type    string   `yaml:"type" toml:"type" json:"type"`

	CreatedAt string `yaml:"created_at" toml:"created_at" json:"created_at"`
	UpdatedAt string `yaml:"updated_at" toml:"updated_at" json:"updated_at"`
}

type GoTemplateSet struct {
	templates *template.Template
}

type HtmlTemplateData struct {
	Content template.HTML
}

type MarkdownDocument struct {
	PostMetaData  `yaml:"post_metadata" toml:"post_metadata" json:"post_metadata"`
	Content       []byte `yaml:"-" toml:"-" json:"-"`
	ContentString string `yaml:"content" toml:"content" json:"content"`
}

type AppState struct {
	Posts []MarkdownDocument
	Pages []MarkdownDocument
	Tags  []string
}

func (a AppState) bootstrap() error {
	pages, err := createDocuments(utils.UseRelativePath(MARKDOWN_PAGES_PATH))
	if err != nil {
		return utils.ReturnError("bootstrap:44 ", err)
	}

	appState.Pages = pages

	posts, err := createDocuments(utils.UseRelativePath(MARKDOWN_POSTS_PATH))
	if err != nil {
		return utils.ReturnError("bootstrap:52 ", err)
	}

	appState.Posts = posts
	allContent := append(pages, posts...)

	var tags []string
	sort.Strings(appState.Tags)

	for _, doc := range allContent {
		for _, tag := range doc.Tags {
			if !utils.IsStringInSlice(tag, appState.Tags) {
				tags = append(tags, tag)
			}
		}
	}

	appState.Tags = tags
	return nil
}

func BackdoorBootstrap() *AppState {
	backdoorAppState := appState
	backdoorAppState.bootstrap()

	return &backdoorAppState
}
