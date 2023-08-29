package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubbles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.5 19a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Zm4.475-9.025a3.5 3.5 0 1 1-4.95-4.95a3.5 3.5 0 0 1 4.95 4.95ZM6.5 19a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM8 15a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12 9.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-4.5.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}