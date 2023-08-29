package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 14h6m-6 8h6m-8-4h10m-13.5-.618C2.188 16.707 1 15.388 1 13c0-4 3.333-5 5-5c0-2 0-6 6-6s6 4 6 6c1.667 0 5 1 5 5c0 2.388-1.188 3.707-2.5 4.382"/>`),
		g.Group(children),
	)
}