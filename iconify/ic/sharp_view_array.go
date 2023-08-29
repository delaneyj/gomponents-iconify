package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpViewArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5h-3v14h3V5zm-4 0H7v14h10V5zM6 5H3v14h3V5z"/>`),
		g.Group(children),
	)
}