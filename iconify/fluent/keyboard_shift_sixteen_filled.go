package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.184 1.571a1.511 1.511 0 0 0-2.367 0L2.218 7.373c-.52.656-.05 1.621.789 1.621h1.978V13c0 .553.45 1.001 1.005 1.001h4.02c.556 0 1.005-.448 1.005-1.001V8.994h1.978c.84 0 1.31-.964.789-1.62L9.184 1.57Z"/>`),
		g.Group(children),
	)
}