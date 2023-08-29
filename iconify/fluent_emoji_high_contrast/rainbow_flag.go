package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 5a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H3Zm0 4v3h26V9H3Zm26 7H3v4h26v-4Zm0 7H3v2h26v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}