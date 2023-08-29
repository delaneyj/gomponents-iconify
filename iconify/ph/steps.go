package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 56a8 8 0 0 1-8 8h-48v40a8 8 0 0 1-8 8h-48v40a8 8 0 0 1-8 8H80v40a8 8 0 0 1-8 8H16a8 8 0 0 1 0-16h48v-40a8 8 0 0 1 8-8h48v-40a8 8 0 0 1 8-8h48V56a8 8 0 0 1 8-8h56a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}