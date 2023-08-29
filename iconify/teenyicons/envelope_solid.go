package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5 2.5A1.5 1.5 0 0 1 2 1h12a1.5 1.5 0 0 1 1.5 1.5v1.208L8 7.926L.5 3.708V2.5Z"/><path fill="currentColor" d="M.5 4.855V12.5A1.5 1.5 0 0 0 2 14h12a1.5 1.5 0 0 0 1.5-1.5V4.855L8 9.074L.5 4.854Z"/>`),
		g.Group(children),
	)
}