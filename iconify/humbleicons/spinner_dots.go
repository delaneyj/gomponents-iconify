package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M13 4a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM7.34 6.34a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm11.32 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 11.32a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-11.32 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM21 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-8 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-8-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}