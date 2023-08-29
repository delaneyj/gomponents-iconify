package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersZeroOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9a5 5 0 0 0-5-5h-1v.1A5.002 5.002 0 0 0 7 9v6a5 5 0 0 0 10 0V9Zm-5-3a3 3 0 0 1 3 3v6a3 3 0 1 1-6 0V9a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}