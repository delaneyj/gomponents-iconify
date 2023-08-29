package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneCallMissedOutgoing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10.41V15h2V7h-8v2h4.59L12 14.59L4.41 7L3 8.41l9 9z"/>`),
		g.Group(children),
	)
}