package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingGibbous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 12c0-4.5 1.93-8.74 6-10a10 10 0 0 1 0 20c-4.07-1.26-6-5.5-6-10Z"/>`),
		g.Group(children),
	)
}