package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0A7.498 7.498 0 0 0 0 7.495a7.498 7.498 0 0 0 7.5 7.496c1.306 0 2.91-.328 4.054-.947l2.79.922a.5.5 0 0 0 .63-.634l-.926-2.771c.672-1.173.952-2.706.952-4.066A7.498 7.498 0 0 0 7.5 0Z"/>`),
		g.Group(children),
	)
}