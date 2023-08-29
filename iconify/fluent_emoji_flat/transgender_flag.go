package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#83CBFF" fill-rule="evenodd" d="M3 6a1 1 0 0 0-1 1v3l2 1h24l2-1V7a1 1 0 0 0-1-1H3Zm26 20a1 1 0 0 0 1-1v-3l-2-1H4l-2 1v3a1 1 0 0 0 1 1h26Z" clip-rule="evenodd"/><path fill="#FFB2FF" fill-rule="evenodd" d="M30 10H2v4l2 1h24l2-1v-4ZM4 17l-2 1v4h28v-4l-2-1H4Z" clip-rule="evenodd"/><path fill="#F3EEF8" d="M2 14h28v4H2v-4Z"/></g>`),
		g.Group(children),
	)
}