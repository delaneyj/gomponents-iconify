package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.5 5.25a1.5 1.5 0 0 0-3 0v31.835L10.32 24.698a1.5 1.5 0 1 0-2.14 2.104l14.75 15l.031.03c.27.259.636.418 1.039.418a1.495 1.495 0 0 0 1.07-.448l14.75-15a1.5 1.5 0 1 0-2.14-2.104L25.5 37.085V5.25Z"/>`),
		g.Group(children),
	)
}