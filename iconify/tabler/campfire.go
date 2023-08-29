package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campfire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 21l16-4m0 4L4 17m8-2a4 4 0 0 0 4-4c0-3-2-3-2-8c-4 2-6 5-6 8a4 4 0 0 0 4 4z"/>`),
		g.Group(children),
	)
}