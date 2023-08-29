package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joystick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm1 7h2a3 3 0 0 1 3 3v4H0v-4a3 3 0 0 1 3-3h6V7.874A4.002 4.002 0 0 1 10 0a4 4 0 0 1 1 7.874V13zm-9 5h12v-2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v2zm2.5-6h2a1.5 1.5 0 0 1 0 3h-2a1.5 1.5 0 0 1 0-3z"/>`),
		g.Group(children),
	)
}