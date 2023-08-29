package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unplug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 5l3-3M2 22l3-3m1.3 1.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6l-2.3 2.3a2.4 2.4 0 0 0 0 3.4Zm1.2-6.8L10 11m.5 5.5L13 14m-1-8l6 6l2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"/>`),
		g.Group(children),
	)
}