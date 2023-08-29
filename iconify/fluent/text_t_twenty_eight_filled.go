package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0V6h-6v16h1a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2h1V6H7v2a1 1 0 0 1-2 0V5Z"/>`),
		g.Group(children),
	)
}