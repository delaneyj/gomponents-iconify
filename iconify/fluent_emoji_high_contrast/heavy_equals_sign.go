package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyEqualsSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.063 7C3.381 7 2 8.343 2 10c0 1.647 1.381 3 3.063 3h21.874C28.618 13 30 11.657 30 10c0-1.647-1.372-3-3.063-3H5.063Zm0 12C3.381 19 2 20.343 2 22c0 1.647 1.381 3 3.063 3h21.874C28.618 25 30 23.657 30 22c0-1.647-1.372-3-3.063-3H5.063Z"/>`),
		g.Group(children),
	)
}