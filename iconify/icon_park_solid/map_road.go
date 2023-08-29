package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMapRoad0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" d="m33 12l4 24M16 12l-4 24m12-24v4m0 6v4m0 6v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMapRoad0)"/>`),
		g.Group(children),
	)
}