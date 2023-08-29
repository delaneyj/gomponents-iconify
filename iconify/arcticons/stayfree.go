package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stayfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 23.946a21.497 21.497 0 1 1-.001-.173"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.53 21.994l5.827 4.91l10.027-8.81"/>`),
		g.Group(children),
	)
}