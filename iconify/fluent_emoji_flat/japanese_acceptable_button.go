package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseAcceptableButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FF822D" d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Z"/><path fill="#fff" d="M9 8a1 1 0 1 0 0 2h11v12h-3a1 1 0 1 0 0 2h4a1 1 0 0 0 1-1V10h1a1 1 0 1 0 0-2H9Z"/><path fill="#fff" d="M11 12a1 1 0 0 0-1 1v7a1 1 0 1 0 2 0h5a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6Zm5 2v4h-4v-4h4Z"/></g>`),
		g.Group(children),
	)
}