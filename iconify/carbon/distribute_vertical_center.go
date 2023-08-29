package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 21h-4v-1a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v1H2v2h4v1a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2v-1h4zm-6 3H8v-4l16-.001zm6-15h-8V8a2.002 2.002 0 0 0-2-2h-8a2.002 2.002 0 0 0-2 2v1H2v2h8v1a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2v-1h8zm-10 3h-8V8l8-.001z"/>`),
		g.Group(children),
	)
}