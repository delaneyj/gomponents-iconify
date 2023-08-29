package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCircleTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3.5c.23 0 .417.187.417.417v1.666h1.666a.417.417 0 0 1 0 .834H6.417v1.666a.417.417 0 0 1-.834 0V6.417H3.917a.417.417 0 0 1 0-.834h1.666V3.917c0-.23.187-.417.417-.417ZM1 6a5 5 0 1 1 10 0A5 5 0 0 1 1 6Zm5-4.167a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334Z"/>`),
		g.Group(children),
	)
}