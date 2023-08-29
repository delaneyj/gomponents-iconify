package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M24 5a1 1 0 0 1 .894.553l18 36A1 1 0 0 1 42 43H6a1 1 0 0 1-.894-1.447l18-36A1 1 0 0 1 24 5Z"/>`),
		g.Group(children),
	)
}