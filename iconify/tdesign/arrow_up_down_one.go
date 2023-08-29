package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 20.414L11.586 15.5L13 14.086l2.5 2.5V4.5h2v12.086l2.5-2.5l1.414 1.414l-4.914 4.914Zm-10-.914V7.414L4 9.914L2.586 8.5L7.5 3.586L12.414 8.5L11 9.914l-2.5-2.5V19.5h-2Z"/>`),
		g.Group(children),
	)
}