package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 16.5V3m0 0l3.5 3.5M18 3l-3.5 3.5m3.5 10a3.5 3.5 0 1 1-7 0v-9m0 0a3.5 3.5 0 1 0-7 0v12"/>`),
		g.Group(children),
	)
}