package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.5L2 9.8l3.8 11.7h12.4L22 9.8L12 2.5Z"/>`),
		g.Group(children),
	)
}