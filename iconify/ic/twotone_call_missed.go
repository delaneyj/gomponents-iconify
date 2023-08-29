package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneCallMissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 10.41l7 7l9-9L19.59 7L12 14.59L6.41 9H11V7H3v8h2z"/>`),
		g.Group(children),
	)
}