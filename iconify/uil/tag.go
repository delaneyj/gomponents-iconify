package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.12 10.71l-8.41-8.42A1 1 0 0 0 12 2H3a1 1 0 0 0-1 1v9a1 1 0 0 0 .29.71l8.42 8.41a3 3 0 0 0 4.24 0L21.12 15a3 3 0 0 0 0-4.24Zm-1.41 2.82l-6.18 6.17a1 1 0 0 1-1.41 0L4 11.59V4h7.59l8.12 8.12a1 1 0 0 1 .29.71a1 1 0 0 1-.29.7Z"/>`),
		g.Group(children),
	)
}