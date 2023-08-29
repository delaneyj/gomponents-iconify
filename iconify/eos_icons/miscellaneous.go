package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miscellaneous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 13l.989 7.875A1 1 0 0 1 17.997 22H6.003a1 1 0 0 1-.992-1.125L6 13Zm3-3H3v2h18Zm-6.286-4.196A6.303 6.303 0 0 0 15 3.835C15 2.27 14.552 1 14 1s-1 1.27-1 2.835a7.115 7.115 0 0 0 .115 1.301a4.626 4.626 0 0 0-2.234.001A7.094 7.094 0 0 0 11 3.835C11 2.27 10.552 1 10 1S9 2.27 9 3.835a6.31 6.31 0 0 0 .283 1.971A5.11 5.11 0 0 0 7 9h2a1 1 0 0 1 2 0h2a1 1 0 0 1 2 0h2a5.11 5.11 0 0 0-2.286-3.196Z"/>`),
		g.Group(children),
	)
}