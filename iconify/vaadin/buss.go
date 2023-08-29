package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.67 4H14V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v2h-.68a.32.32 0 0 0-.32.32v2.36c0 .177.143.32.32.32H2v6c0 .55 0 1 1 1v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14h4v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14c1 0 1-.45 1-1V7h.67a.33.33 0 1 0 0-.66a.33.33 0 0 0 0 .66a.33.33 0 0 0 .33-.33V4.33a.33.33 0 0 0-.33-.33zM6 1h4v1H6V1zM4 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2zM3 8V3h10v5H3zm9 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}