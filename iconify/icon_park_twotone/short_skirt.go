package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortSkirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShortSkirt0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m12 9l5-5h14l5 5l7 26s-3 9-19 9s-19-9-19-9l7-26Z"/><path d="m13 42l4-16M5 35s3 9 19 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShortSkirt0)"/>`),
		g.Group(children),
	)
}