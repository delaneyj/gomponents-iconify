package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subtraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 0a7 7 0 1 1-3 13.326A7 7 0 1 1 10 .673A6.963 6.963 0 0 1 13 0ZM7 2a5 5 0 1 0 1.002 9.9A6.972 6.972 0 0 1 6 7c0-1.907.763-3.637 2-4.9A4.976 4.976 0 0 0 7 2Zm6 0a5 5 0 1 0 0 10a5 5 0 0 0 0-10Z"/>`),
		g.Group(children),
	)
}