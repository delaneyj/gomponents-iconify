package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.82 3 3 8.82 3 16s5.82 13 13 13s13-5.82 13-13S23.18 3 16 3ZM1 16C1 7.716 7.716 1 16 1c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15c-8.284 0-15-6.716-15-15Z"/>`),
		g.Group(children),
	)
}