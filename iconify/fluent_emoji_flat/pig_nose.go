package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FF8687" d="M30 21.79C29.89 14.17 23.66 9 16 9S2.11 14.17 2 21.79V22c0 4.42 3.58 8 8 8h12c4.42 0 8-3.58 8-8v-.21Z"/><path fill="#1C1C1C" d="M10 16a3 3 0 0 0-3 3v4a3 3 0 1 0 6 0v-4a3 3 0 0 0-3-3Zm12 0a3 3 0 0 0-3 3v4a3 3 0 1 0 6 0v-4a3 3 0 0 0-3-3Z"/></g>`),
		g.Group(children),
	)
}