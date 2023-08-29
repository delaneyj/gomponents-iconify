package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 5a1 1 0 0 1 1-1h12.5a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0V6H13v12h1a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2h1V6H6.75v1a1 1 0 0 1-2 0V5Z"/>`),
		g.Group(children),
	)
}