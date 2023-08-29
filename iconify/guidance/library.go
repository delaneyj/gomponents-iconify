package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 7.5V17m0-9.5a3 3 0 0 0-3-3H2v1.085A62.99 62.99 0 0 1 .5 19.25v.25H9a3 3 0 0 1 3 3m0-15a3 3 0 0 1 3-3h7v1.085c0 4.596.503 9.178 1.5 13.665v.25H15a3 3 0 0 0-3 3m0 0v.5"/>`),
		g.Group(children),
	)
}