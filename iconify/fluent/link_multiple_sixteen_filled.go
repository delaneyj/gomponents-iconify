package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 6.5A3.5 3.5 0 0 1 4.5 3h4a3.5 3.5 0 0 1 3.469 3.031a3.537 3.537 0 0 1-.122 1.499A3.502 3.502 0 0 1 8.5 10h-.75a.75.75 0 0 1 0-1.5h.75a2 2 0 1 0 0-4h-4a2 2 0 0 0-1.262 3.552a4.494 4.494 0 0 0-.235 1.613A3.5 3.5 0 0 1 1 6.5Zm8 .25a.75.75 0 0 1-.75.75H7.5a2 2 0 1 0 0 4h4a2 2 0 0 0 1.263-3.551a4.495 4.495 0 0 0 .235-1.613A3.5 3.5 0 0 1 11.5 13h-4a3.5 3.5 0 1 1 0-7h.75a.75.75 0 0 1 .75.75Z"/>`),
		g.Group(children),
	)
}