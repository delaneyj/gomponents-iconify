package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StampSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 5 3.163V9h3.5a2.5 2.5 0 0 1 2.5 2.5V13H0v-1.5A2.5 2.5 0 0 1 2.5 9H6V6.663A3.5 3.5 0 0 1 4 3.5ZM0 14v1h15v-1H0Z"/>`),
		g.Group(children),
	)
}