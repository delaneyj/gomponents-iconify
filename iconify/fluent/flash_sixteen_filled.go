package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.873 1a1 1 0 0 0-.959.714L3.031 8.036A.75.75 0 0 0 3.75 9h1.583l-1.28 4.389c-.384 1.316 1.324 2.2 2.178 1.128l6.607-8.3A.75.75 0 0 0 12.25 5h-2.03l.994-2.649A1 1 0 0 0 10.279 1H5.873Z"/>`),
		g.Group(children),
	)
}