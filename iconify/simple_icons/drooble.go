package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drooble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 11.986a7.599 7.599 0 0 0-7-7.559v7.574a5 5 0 0 1-10 0c0-3.604 3.707-6.022 7-4.583V.17C6.615-1.069 0 4.63 0 12c0 6.628 5.373 12 12 12c6.628 0 12-5.372 12-12v-.014m-14 .015a2 2 0 1 0 4 0a2 2 0 0 0-4 0m14-.015a7.599 7.599 0 0 0-7-7.559v7.574a5 5 0 0 1-10 0c0-3.604 3.707-6.022 7-4.583V.17C6.615-1.069 0 4.63 0 12c0 6.628 5.373 12 12 12c6.628 0 12-5.372 12-12v-.014m-14 .015a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		g.Group(children),
	)
}