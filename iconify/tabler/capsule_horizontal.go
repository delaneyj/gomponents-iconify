package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a6 6 0 0 1 6-6h6a6 6 0 0 1 6 6v0a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6z"/>`),
		g.Group(children),
	)
}