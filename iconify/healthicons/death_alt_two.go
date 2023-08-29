package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeathAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32.01 4.013L16.01 4L11 16.992l4.978 26.99l16 .018L37 17.013l-4.99-13ZM23 26v-8.003L20 18l-.002-2L23 15.998V13h2v2.996l2.998-.003l.002 2l-3 .003V26h-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}