package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M14.5 2.315c3.514.904 6.28 3.67 7.185 7.185M22 12a10 10 0 0 1-18 6m-1.808-4.05A10 10 0 0 1 12 2"/>`),
		g.Group(children),
	)
}