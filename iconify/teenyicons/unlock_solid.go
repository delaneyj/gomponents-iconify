package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V6H2.5A1.5 1.5 0 0 0 1 7.5v6A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 12.5 6H5V3.5Z"/>`),
		g.Group(children),
	)
}