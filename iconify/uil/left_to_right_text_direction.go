package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftToRightTextDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.42 17.62a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H3.5a1 1 0 0 0 0 2h14.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76ZM8.5 10v4a1 1 0 0 0 2 0V4h2v10a1 1 0 0 0 2 0V4h1a1 1 0 0 0 0-2h-7a4 4 0 0 0 0 8Zm0-6v4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}