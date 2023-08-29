package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActionChains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="20.01" cy="20.01" r="3" fill="currentColor"/><path fill="currentColor" d="M18 3H6.83a3 3 0 1 0 0 2H18a3 3 0 0 1 0 6h-3.17a3 3 0 0 0-5.64 0H6a5 5 0 0 0 0 10h6v3l4-4l-4-4v3H6a3 3 0 1 1 0-6h3.2a3 3 0 0 0 5.63 0H18a5 5 0 0 0 0-10Z"/>`),
		g.Group(children),
	)
}