package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoggyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.584 13.007a8 8 0 1 1 14.873-5.908a5.5 5.5 0 0 1 6.52 5.908H1.584ZM4 19h17v2H4v-2Zm-2-4h21v2H2v-2Z"/>`),
		g.Group(children),
	)
}