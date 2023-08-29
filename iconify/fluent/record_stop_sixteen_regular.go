package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordStopSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6ZM1 8a7 7 0 1 1 14 0A7 7 0 0 1 1 8Zm7-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2Z"/>`),
		g.Group(children),
	)
}