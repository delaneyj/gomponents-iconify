package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbtackSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8.702 1.41L8.146.853l.708-.708l6 6l-.707.708l-.556-.556l-4.456 7.13l.719.718l-.708.708L5 10.707L.854 14.854l-.708-.707L4.293 10L.146 5.854l.708-.708l.718.72l7.13-4.457Z"/>`),
		g.Group(children),
	)
}