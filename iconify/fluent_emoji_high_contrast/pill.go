package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.318 3.528a8.629 8.629 0 0 1 12.204 0a8.629 8.629 0 0 1 0 12.204l-7.097 7.097l-.01-.01l-5.683 5.683a8.629 8.629 0 0 1-12.204 0a8.629 8.629 0 0 1 0-12.204L10.625 9.2l.01.01l5.683-5.683ZM20 21.405l-9.376-9.376l-5.683 5.683a6.629 6.629 0 0 0 0 9.376a6.629 6.629 0 0 0 9.376 0L20 21.405ZM28 10a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z"/>`),
		g.Group(children),
	)
}