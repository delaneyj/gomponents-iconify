package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrend0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" d="m13.44 29.835l5.657-5.657l4.388 4.377L34 18"/><path stroke="#000" stroke-linecap="round" d="M26 18h8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrend0)"/>`),
		g.Group(children),
	)
}