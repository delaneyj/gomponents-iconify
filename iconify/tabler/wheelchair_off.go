package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16a5 5 0 1 0 10 0a5 5 0 1 0-10 0m14.582 1.59a2 2 0 0 0 2.833 2.824M14 14h-1.4M6 6v5m0-3h2m4 0h5m-2 0v3M3 3l18 18"/>`),
		g.Group(children),
	)
}