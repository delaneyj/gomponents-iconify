package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LexmarkPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.959 5.5h9.904m8.281 5.672H24.36L18.959 5.5H8.066v33.219l3.511 3.781h25.837l2.52-3.781V14.232l-2.79-3.06zm2.79 3.06H8.066m31.868 9.363H8.066"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.066 33.948h10.217l3.129 3.128h18.522m-20.03-20.188h10.083v3.601H19.904z"/>`),
		g.Group(children),
	)
}