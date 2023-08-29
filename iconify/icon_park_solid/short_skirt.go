package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortSkirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShortSkirt0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m12 9l5-5h14l5 5l7 26s-3 9-19 9s-19-9-19-9l7-26Z"/><path stroke="#000" d="m13 42l4-16"/><path stroke="#fff" d="M5 35s3 9 19 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShortSkirt0)"/>`),
		g.Group(children),
	)
}