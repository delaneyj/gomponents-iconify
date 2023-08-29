package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 7a2 2 0 1 0-2-2m2 6v4h4l4 5m-4-9h1m-9 .5a5 5 0 1 0 6 7.5M3 3l18 18"/>`),
		g.Group(children),
	)
}