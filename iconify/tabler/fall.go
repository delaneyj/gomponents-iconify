package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 21l1-5l-1-4l-3-4h4l3-3M6 16l-1-4l3-4M5 5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m8.5 7H16l4 2"/>`),
		g.Group(children),
	)
}