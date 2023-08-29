package feather

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDrizzle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19v2m0-8v2m8 4v2m0-8v2m-4 6v2m0-8v2m8-.42A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25"/>`),
		g.Group(children),
	)
}