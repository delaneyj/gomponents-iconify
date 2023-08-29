package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0A5.5 5.5 0 0 0 2 5.5V10H1v1h13v-1h-1V5.5A5.5 5.5 0 0 0 7.5 0ZM5 12.5V12h5v.5a2.5 2.5 0 0 1-5 0Z"/>`),
		g.Group(children),
	)
}