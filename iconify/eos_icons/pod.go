package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19.003A1.999 1.999 0 0 0 8 21h8a1.999 1.999 0 0 0 2-1.997V18H6Zm9-15.011A.995.995 0 0 0 14.002 3H9.998A.995.995 0 0 0 9 3.992V5h6ZM14.993 6H9L4 17h16Zm-2.983 7.01a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}