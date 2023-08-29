package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.026 4.843a.5.5 0 0 1 .632-.316L6 4.974l1.342-.447a.5.5 0 0 1 .316.948L6.5 5.861v1.14a.5.5 0 0 1-1 0V5.86l-1.158-.386a.5.5 0 0 1-.316-.632Zm1.547-2.777a1.5 1.5 0 0 1 .854 0l2.858.85a1 1 0 0 1 .715.958v4.254a1 1 0 0 1-.715.958l-2.858.85a1.5 1.5 0 0 1-.854 0l-2.858-.85A1 1 0 0 1 2 8.128V3.866a.99.99 0 0 1 .708-.949l2.865-.85Zm.57.959a.5.5 0 0 0-.285 0L3 3.874v4.254l2.858.849a.5.5 0 0 0 .284 0L9 8.127V3.875l-2.858-.85Z"/>`),
		g.Group(children),
	)
}