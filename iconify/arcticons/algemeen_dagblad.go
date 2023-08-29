package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlgemeenDagblad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.523 36.813h5.766c6.192 0 11.211-5.02 11.211-11.211v-3.204c0-6.191-5.02-11.21-11.21-11.21h-9.414v17.136m0 0H8.148M4.5 36.813l11.012-25.625l11.011 25.625"/>`),
		g.Group(children),
	)
}