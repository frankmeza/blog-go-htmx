package blog

type LinkData struct {
	Text string
	Url  string
}

type Metadata struct {
	Name        string
	CurrentPost HtmlTemplateData
	CurrentPage HtmlTemplateData
	FooterLinks []LinkData
	HeaderLinks []LinkData
	Posts       []MarkdownDocument
	Pages       []MarkdownDocument
	Summary     string
}

var BlogData = Metadata{
	Name: "your blog name",
	FooterLinks: []LinkData{
		{Text: "GITHUB", Url: ""},
		{Text: "LINKED IN", Url: ""},
		{Text: "TELEGRAM", Url: ""},
		{Text: "INSTAGRAM", Url: ""},
	},
	HeaderLinks: []LinkData{
		{Text: "ABOUT", Url: "/pages/about"},
		{Text: "CONTACT", Url: "/pages/contact"},
		{Text: "NOW", Url: "/pages/now"},
	},
	Summary: "another smart guy programmer blog 👨🏾‍💻 ",
}
