package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ampersand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20L8.597 9.028a2.948 2.948 0 0 1 0-4.165a2.94 2.94 0 0 1 4.161 0a2.948 2.948 0 0 1 0 4.165l-4.68 4.687a3.685 3.685 0 0 0 0 5.207a3.675 3.675 0 0 0 5.2 0L19 13"/>`),
		g.Group(children),
	)
}