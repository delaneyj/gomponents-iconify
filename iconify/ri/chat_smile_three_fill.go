package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSmileThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10H2l2.929-2.929A9.969 9.969 0 0 1 2 12C2 6.477 6.477 2 12 2Zm4 11H8a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}