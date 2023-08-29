package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h8.586a1.5 1.5 0 0 1 1.06.44l3.415 3.414A1.5 1.5 0 0 1 15 4.914V13.5a1.5 1.5 0 0 1-1.5 1.5H11v-3.5A1.5 1.5 0 0 0 9.5 10h-4A1.5 1.5 0 0 0 4 11.5V15H1.5A1.5 1.5 0 0 1 0 13.5v-12Z"/><path fill="currentColor" d="M5 15h5v-3.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5V15Z"/>`),
		g.Group(children),
	)
}