package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 0-9.56 12.91a.29.29 0 0 0 0 .1a9.83 9.83 0 0 0 1.79 3.32l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 9.46-6.78v-.08A9.89 9.89 0 0 0 22 12A10 10 0 0 0 12 2Zm0 18H5.41l.3-.29l8.41-8.42a1 1 0 0 1 1.4 0l3.62 3.6l.23.22A8 8 0 0 1 12 20Zm-7.46-5.13l1.58-1.58a1 1 0 0 1 1.41 0l.87.87l-2.72 2.74a7.67 7.67 0 0 1-1.14-2.03Zm15.41-2l-3-3a3 3 0 0 0-4.24 0l-2.89 2.89l-.88-.87a3 3 0 0 0-4.23 0l-.71.67A5.25 5.25 0 0 1 4 12a8 8 0 0 1 16 0a8.27 8.27 0 0 1 0 .86Z"/>`),
		g.Group(children),
	)
}