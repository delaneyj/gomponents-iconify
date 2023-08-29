package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jetpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6a3 3 0 1 0-6 0v7h6V6zm4 7h6V6a3 3 0 0 0-6 0v7zm-9 3c0 2.333.667 4 2 5c1.333-1 2-2.667 2-5m6 0c0 2.333.667 4 2 5c1.333-1 2-2.667 2-5m-9-8h4m-4 3h4"/>`),
		g.Group(children),
	)
}