package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallEndSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 17L.4 13.475l1-1.025q2.175-2.225 4.963-3.337T12 8q2.85 0 5.625 1.113T22.6 12.45l1 1.025L19.975 17L16 14v-3.35q-.95-.3-1.95-.475T12 10q-1.05 0-2.05.175T8 10.65V14l-3.975 3Z"/>`),
		g.Group(children),
	)
}