package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appsmonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8v16l8.32 7.268"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 23.946a21.497 21.497 0 1 1-.001-.173"/>`),
		g.Group(children),
	)
}