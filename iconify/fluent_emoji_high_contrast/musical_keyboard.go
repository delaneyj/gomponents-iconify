package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.036 2h26.928C30.536 2 31 2.448 31 3v25c0 1.105-.927 2-2.071 2H4.07C2.927 30 2 29.105 2 28V3c0-.552.464-1 1.036-1ZM30 9h-5v12a1 1 0 0 1-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V9Zm-8 0h-4v12a1 1 0 0 1-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 1-1-1V9Zm-7 0h-4v12a1 1 0 0 1-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 1-1-1V9ZM8 9H3v19a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 1-1-1V9Z"/>`),
		g.Group(children),
	)
}