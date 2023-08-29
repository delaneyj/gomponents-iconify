package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.55 9.704c3.46-3.55 9.14-3.61 12.69-.15c3.28 3.21 3.58 8.33.86 11.89a8.372 8.372 0 0 1-6.83 3.52H8.58c-2.11 0-3.98-1-5.17-2.56A6.47 6.47 0 0 1 2 18.374a6.457 6.457 0 0 1 6.589-6.458a4.04 4.04 0 0 1 5.624-1.849c.109-.123.22-.244.337-.363Z"/>`),
		g.Group(children),
	)
}