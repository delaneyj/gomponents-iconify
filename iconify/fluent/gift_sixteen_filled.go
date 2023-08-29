package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 5H7v3H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1.268A2 2 0 0 1 7.5 2.677A2 2 0 0 1 10.732 5H12a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H8V5ZM5 4a1 1 0 0 0 1 1h1V4a1 1 0 0 0-2 0Zm3 1h1a1 1 0 1 0-1-1v1Zm4 4H8v5h2a2 2 0 0 0 2-2V9Zm-5 5V9H3v3a2 2 0 0 0 2 2h2Z"/>`),
		g.Group(children),
	)
}