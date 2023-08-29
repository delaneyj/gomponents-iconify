package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 5.5a4.5 4.5 0 0 1 9 0V7H21v16H3V7h4.5V5.5Zm0 3.5H5v12h14V9h-2.5v3h-2V9h-5v3h-2V9Zm7-2V5.5a2.5 2.5 0 0 0-5 0V7h5Z"/>`),
		g.Group(children),
	)
}