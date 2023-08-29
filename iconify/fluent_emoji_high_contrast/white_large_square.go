package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1 4a3 3 0 0 1 3-3h24a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4Zm3-1a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h24a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4Z"/>`),
		g.Group(children),
	)
}