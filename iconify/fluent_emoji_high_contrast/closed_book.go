package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.5 1.5a2 2 0 0 0-2 2v25a2 2 0 0 0 2 2h19a2 2 0 0 0 1.886-1.333l.236-.667H7a.5.5 0 0 1 0-1h20.5V4A2.5 2.5 0 0 0 25 1.5H6.5Zm20 24h-18v-23H25A1.5 1.5 0 0 1 26.5 4v21.5Z"/>`),
		g.Group(children),
	)
}