package blog

import (
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/frankmeza/blog-go-htmx/pkg/utils"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v4"
)

func bootstrapAppState(c echo.Context) error {
	if err := appState.bootstrap(); err != nil {
		return utils.ReturnError("bootstrapAppState:44 ", err)
	}

	return renderHome(c)
}

func convertBytesToHtmlTemplate(content []byte) HtmlTemplateData {
	parserExtensions := parser.CommonExtensions | parser.Attributes | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(parserExtensions)

	html := markdown.ToHTML(content, parser, nil)
	return HtmlTemplateData{Content: template.HTML(html)}
}

func createDocuments(path string) ([]MarkdownDocument, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		e := utils.ReturnError("createDocuments:23 ", err)
		return []MarkdownDocument{}, e
	}

	var documents []MarkdownDocument

	for _, f := range files {
		doc, err := createDocumentFromFilePath(path + "/" + f.Name())

		if err != nil {
			e := utils.ReturnError("createDocuments:32 ", err)
			return []MarkdownDocument{}, e
		}

		documents = append(documents, doc)
	}

	return documents, nil
}

func createDocumentFromFilePath(path string) (MarkdownDocument, error) {
	markdownAsBytes, err := ioutil.ReadFile(path)
	if err != nil {
		e := utils.ReturnError("createDocumentFromFilePath:42 ", err)
		return MarkdownDocument{}, e
	}

	content, postMetaData, err := getBlogMetadata(markdownAsBytes)
	if err != nil {
		e := utils.ReturnError("createDocumentFromFilePath:48 ", err)
		return MarkdownDocument{}, e
	}

	document := MarkdownDocument{
		Content:       content,
		ContentString: string(content),
		PostMetaData:  postMetaData,
	}

	return document, nil
}

func getBlogMetadata(markdownAsBytes []byte) ([]byte, PostMetaData, error) {
	var postMetaData PostMetaData
	yamlReader := strings.NewReader(string(markdownAsBytes))

	content, err := frontmatter.Parse(yamlReader, &postMetaData)
	if err != nil {
		e := utils.ReturnError("getBlogMetadata:27 ", err)
		return nil, postMetaData, e
	}

	return content, postMetaData, nil
}

func getMdPagePath(pageName string) string {
	return MARKDOWN_PAGES_PATH + pageName + DOT_MD
}

func getMdPostPath(postName string) string {
	return MARKDOWN_POSTS_PATH + postName + DOT_MD
}

func renderHome(c echo.Context) error {
	var publishedPosts []MarkdownDocument
	for _, post := range appState.Posts {
		if !post.IsDraft {
			publishedPosts = append(publishedPosts, post)
		}
	}

	blogData := BlogData
	blogData.Posts = publishedPosts

	return t.templates.ExecuteTemplate(
		c.Response().Writer,
		TEMPLATE_HOME,
		blogData,
	)
}

func renderMarkdownAsHtml(c echo.Context) error {
	pageName := c.Param("page")
	postName := c.Param("post")

	var path string
	if pageName != "" {
		path = getMdPagePath(pageName)
	}

	if postName != "" {
		path = getMdPostPath(postName)
	}

	markdownDocument, err := createDocumentFromFilePath(
		utils.UseRelativePath(path),
	)

	if err != nil {
		return utils.ReturnError("renderMarkdownAsHtml:73 ", err)
	}

	pageContent := convertBytesToHtmlTemplate(
		markdownDocument.Content,
	)

	blogData := BlogData
	blogData.CurrentPost = pageContent

	return t.templates.ExecuteTemplate(
		c.Response().Writer,
		TEMPLATE_MARKDOWN_PAGE,
		blogData,
	)
}
