package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableOffsetTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 3H12v4H3V5.5A2.5 2.5 0 0 1 5.5 3ZM8 8h9v4H8V8Zm-1 4V8H3v4h4Zm-4 1h9v4H5.5A2.5 2.5 0 0 1 3 14.5V13Zm10 0v4h1.5a2.5 2.5 0 0 0 2.5-2.5V13h-4Zm0-6h4V5.5A2.5 2.5 0 0 0 14.5 3H13v4Z"/>`),
		g.Group(children),
	)
}