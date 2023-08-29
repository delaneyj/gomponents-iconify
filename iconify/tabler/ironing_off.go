package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IroningOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h6.459a3 3 0 0 1 2.959 2.507l.577 3.464l.804 4.821l.007.044M18 18H3a7 7 0 0 1 7-7h1m4 0h4.8M3 3l18 18"/>`),
		g.Group(children),
	)
}