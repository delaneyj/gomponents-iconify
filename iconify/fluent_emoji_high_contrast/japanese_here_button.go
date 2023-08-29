package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseHereButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M5 8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v15a1 1 0 1 1-2 0H6a1 1 0 1 1 0-2h7V9H6a1 1 0 0 1-1-1Zm12 0a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v15a1 1 0 1 1-2 0h-7a1 1 0 1 1 0-2h7V9h-7a1 1 0 0 1-1-1Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}