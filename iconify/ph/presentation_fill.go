package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40h-80V24a8 8 0 0 0-16 0v16H40a16 16 0 0 0-16 16v120a16 16 0 0 0 16 16h39.36l-21.61 27a8 8 0 0 0 12.5 10l29.59-37h56.32l29.59 37a8 8 0 1 0 12.5-10l-21.61-27H216a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}