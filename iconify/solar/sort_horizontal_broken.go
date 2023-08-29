package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHorizontalBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 8l4.125-4M6 8l4.125 4M6 8h7m5 0h-2m2 8l-4.125-4M18 16l-4.125 4M18 16h-7m-5 0h2"/>`),
		g.Group(children),
	)
}