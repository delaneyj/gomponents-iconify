package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChequeredFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 14h4v-4h-4v4Zm0 4v4h4v-4h-4Zm-4 0v4h-4v-4h-4v4H6v-4h4v-4H6v-4h4v4h4v-4h4v4h4v4h-4Zm0 0v-4h-4v4h4Z"/><path fill="currentColor" d="M3 5a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H3Zm23 2v3h3v4h-3v4h3v4h-3v3h-4v-3h-4v3h-4v-3h-4v3H6v-3H3v-4h3v-4H3v-4h3V7h4v3h4V7h4v3h4V7h4Z"/>`),
		g.Group(children),
	)
}