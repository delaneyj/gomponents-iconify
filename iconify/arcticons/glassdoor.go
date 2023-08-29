package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glassdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 10.5v20h6v-20h21a6 6 0 0 0-6-6h-15a6 6 0 0 0-6 6Zm21 27h-21a6 6 0 0 0 6 6h15a6 6 0 0 0 6-6v-20h-6v20Z"/>`),
		g.Group(children),
	)
}