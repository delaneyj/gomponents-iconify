package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm3 10a1 1 0 0 1-.5.866l-3 1.732a1 1 0 0 1-1.5-.866v-3.464a1 1 0 0 1 1.5-.866l3 1.732A1 1 0 0 1 15 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}