package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 10l-6 6v4h4l6-6m1.99-1.99l2.504-2.504a2.828 2.828 0 1 0-4-4l-2.5 2.5M13.5 6.5l4 4M3 3l18 18"/>`),
		g.Group(children),
	)
}