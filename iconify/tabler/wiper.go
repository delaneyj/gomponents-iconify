package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wiper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 18a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 9l5.5 5.5a5 5 0 0 1 7 0L21 9A12 12 0 0 0 3 9m9 9L9.8 5.2"/>`),
		g.Group(children),
	)
}