package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 10C0 4.477 4.477 0 10 0s10 4.477 10 10h-2a8 8 0 1 0-16 0H0zm4 0a6 6 0 1 1 12 0h-2a4 4 0 1 0-8 0H4zm4 0a2 2 0 1 1 4 0H8z"/>`),
		g.Group(children),
	)
}