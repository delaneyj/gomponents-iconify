package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M24.894 8.553a1 1 0 0 0-1.788 0l-15 30A1 1 0 0 0 9 40h30a1 1 0 0 0 .894-1.447l-15-30Z"/>`),
		g.Group(children),
	)
}