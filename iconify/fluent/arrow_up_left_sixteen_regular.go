package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 3a.5.5 0 0 0 0-1h-6a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 1 0V3.707l10.146 10.147a.5.5 0 0 0 .708-.708L3.707 3H8.5Z"/>`),
		g.Group(children),
	)
}