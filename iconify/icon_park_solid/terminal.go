package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTerminal0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" d="m12 18l7 6l-7 6m11 2h13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTerminal0)"/>`),
		g.Group(children),
	)
}