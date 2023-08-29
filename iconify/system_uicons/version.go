package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Version(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 12.5l8 4l8.017-4M2.5 8.657l8.008 3.843l8.009-3.843L10.508 4.5z"/>`),
		g.Group(children),
	)
}