package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 12.299v17.768h11.569a11.868 11.868 0 0 0 0-23.736h-5.601M5.5 41.669V30.067M31.758 6.331H42.5v10.742H31.758z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.339 9.912h3.581v3.581h-3.581z"/>`),
		g.Group(children),
	)
}