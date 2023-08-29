package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefundLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.005 8V5h-16v3h16Zm0 2h-16v9h16v-9Zm-17-7h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm8 11h6v2h-10.5l4.5-4.5V14Z"/>`),
		g.Group(children),
	)
}