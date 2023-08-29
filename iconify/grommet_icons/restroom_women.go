package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestroomWomen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13.5L11 8l1 13m5-7.5L13 8l-1 13m0-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2l1.5 8.5h-5L11 8Z"/>`),
		g.Group(children),
	)
}