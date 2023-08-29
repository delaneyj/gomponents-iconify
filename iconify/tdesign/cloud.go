package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5a5 5 0 0 0-4.916 5.919l.175.942l-.934.215A3.001 3.001 0 0 0 6 18h11a4 4 0 1 0-.067-8l-.86.014l-.142-.848A5.002 5.002 0 0 0 11 5Zm-7 5a7 7 0 0 1 13.723-1.957A6.001 6.001 0 0 1 17 20H6a5 5 0 0 1-1.988-9.589A7.092 7.092 0 0 1 4 10Z"/>`),
		g.Group(children),
	)
}