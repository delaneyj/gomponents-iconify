package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NazarAmulet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M21.32 16.12a5.12 5.12 0 1 1-10.24 0a5.12 5.12 0 0 1 10.24 0Zm-2.56 0a2.56 2.56 0 1 0-5.12 0a2.56 2.56 0 0 0 5.12 0Z"/><path d="M16 1C7.716 1 1 7.716 1 16c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15Zm7.88 15.12a7.68 7.68 0 1 1-15.36 0a7.68 7.68 0 0 1 15.36 0Z"/></g>`),
		g.Group(children),
	)
}