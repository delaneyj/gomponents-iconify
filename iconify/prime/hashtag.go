package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7.25h-3l.77-3.07a.75.75 0 0 0-1.46-.36l-.86 3.43H10l.77-3.07a.75.75 0 0 0-1.46-.36l-.9 3.43H5a.75.75 0 0 0 0 1.5h3l-1.63 6.5H3a.75.75 0 0 0 0 1.5h3l-.77 3.07a.75.75 0 0 0 1.46.36l.86-3.43H14l-.77 3.07a.75.75 0 0 0 1.46.36l.86-3.43H19a.75.75 0 0 0 0-1.5h-3l1.63-6.5H21a.75.75 0 0 0 0-1.5Zm-5 1.5l-1.63 6.5H8l1.63-6.5Z"/>`),
		g.Group(children),
	)
}