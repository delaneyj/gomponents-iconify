package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistorySixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3a5 5 0 1 1-5 5a.5.5 0 0 0-1 0a6 6 0 1 0 2-4.472V2.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1H4a4.992 4.992 0 0 1 4-2Zm.001 2.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1h-1.5V5.5Z"/>`),
		g.Group(children),
	)
}