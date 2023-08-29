package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallEndOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 17L.4 13.475l1-1.025q2.125-2.275 4.938-3.362T12 8q2.85 0 5.638 1.1t4.962 3.35l1 1.025L19.975 17L16 14v-3.35q-.95-.3-1.95-.475T12 10q-1.05 0-2.05.175T8 10.65V14l-3.975 3ZM6 11.45q-.725.375-1.4.863T3.2 13.4l1 1L6 13v-1.55Zm12 .05V13l1.8 1.4l1-.95q-.725-.65-1.4-1.125T18 11.5Zm-12-.05Zm12 .05Z"/>`),
		g.Group(children),
	)
}