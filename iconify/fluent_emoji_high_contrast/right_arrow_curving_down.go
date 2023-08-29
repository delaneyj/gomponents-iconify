package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M10 7a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1.986c.008 0 .014.006.014.014s.007.014.015.013c.156-.018.318-.027.485-.027c2.197 0 3.5 1.599 3.5 3v5.75a.25.25 0 0 1-.25.25h-1.988a.75.75 0 0 0-.528 1.283l4.238 4.195a.75.75 0 0 0 1.056 0l4.238-4.195A.75.75 0 0 0 22.238 20H20.25a.25.25 0 0 1-.25-.25V14c0-3.866-3.358-7-7.5-7c-.165 0-.33.005-.492.015A.007.007 0 0 1 12 7.007A.007.007 0 0 0 11.993 7H10Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}