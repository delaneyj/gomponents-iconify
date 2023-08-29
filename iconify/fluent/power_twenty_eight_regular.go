package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.69 5.267a.75.75 0 0 1-.397.983A9.503 9.503 0 0 0 4.5 15a9.5 9.5 0 0 0 19 0a9.503 9.503 0 0 0-5.793-8.75a.75.75 0 0 1 .586-1.38A11 11 0 0 1 25 15c0 6.075-4.925 11-11 11S3 21.075 3 15A11 11 0 0 1 9.707 4.87a.75.75 0 0 1 .983.397ZM14 2a.75.75 0 0 1 .75.75v9.5a.75.75 0 0 1-1.5 0v-9.5A.75.75 0 0 1 14 2Z"/>`),
		g.Group(children),
	)
}