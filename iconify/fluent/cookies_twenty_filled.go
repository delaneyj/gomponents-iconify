package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookiesTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 0 1 9.874-7.78a.5.5 0 0 1 .232.846a2.001 2.001 0 0 0 .955 3.386A.562.562 0 0 1 13.5 7a2 2 0 0 0 3.503 1.32a.5.5 0 0 1 .868.24A8 8 0 1 1 2 10Zm5-2a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm3 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 3a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm-6 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}