package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseHereButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M6 7a1 1 0 0 0 0 2h7v12H6a1 1 0 1 0 0 2h7a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1H6Zm12 0a1 1 0 1 0 0 2h7v12h-7a1 1 0 1 0 0 2h7a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1h-8Z"/></g>`),
		g.Group(children),
	)
}