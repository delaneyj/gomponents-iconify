package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.32 4.34L9 4l6 3l6-3v13m-2.67 1.335L15 20l-6-3l-6 3V7l2.665-1.333M9 4v1m0 4v8m6-10v4m0 4v5M3 3l18 18"/>`),
		g.Group(children),
	)
}