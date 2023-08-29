package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStyleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M6.676 8.345h8.156v25.633H6.676zm9.322 9.322h8.156v16.312h-8.156zm9.321-11.652h8.156v27.964h-8.156z"/>`),
		g.Group(children),
	)
}