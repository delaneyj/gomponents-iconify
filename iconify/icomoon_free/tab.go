package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 0h1v8h-1V0zM0 8h1v8H0V8zm5 3h11v2H5v2.5L1.5 12L5 8.5V11zm6-6H0V3h11V.5L14.5 4L11 7.5z"/>`),
		g.Group(children),
	)
}