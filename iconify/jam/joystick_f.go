package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoystickF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.915 13H9V7.874A4.002 4.002 0 0 1 10 0a4 4 0 0 1 1 7.874V13h2a3 3 0 0 1 3 3v4H0v-4a3 3 0 0 1 3-3h.085A1.5 1.5 0 0 1 4.5 12h2a1.5 1.5 0 0 1 1.415 1z"/>`),
		g.Group(children),
	)
}