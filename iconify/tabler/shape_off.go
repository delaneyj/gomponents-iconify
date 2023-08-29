package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.575 3.597a2 2 0 0 0 2.849 2.808M17 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m14.574-1.402a2 2 0 0 0 2.826 2.83M5 7v10M9 5h8M7 19h10m2-12v8M3 3l18 18"/>`),
		g.Group(children),
	)
}