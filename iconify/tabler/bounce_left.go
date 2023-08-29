package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BounceLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 15.5c-3-1-5.5-.5-8 4.5c-.5-3-1.5-5.5-3-8M6 9a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}