package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicWand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.414 24L12 6.586a2.048 2.048 0 0 0-2.828 0L6.586 9.172a2.002 2.002 0 0 0 0 2.828l17.413 17.414a2.002 2.002 0 0 0 2.828 0l2.587-2.586a2 2 0 0 0 0-2.828ZM8 10.586L10.586 8l5 5l-2.587 2.587l-5-5ZM25.413 28l-11-10.999L17 14.414l11 11ZM2 16l2-2l2 2l-2 2zM14 4l2-2l2 2l-2 2zM2 4l2-2l2 2l-2 2z"/>`),
		g.Group(children),
	)
}