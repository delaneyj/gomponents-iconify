package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gravethree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 1024H0q0-45 51.5-84T192 874V320q0-87 43-160.5T351.5 43T512 0t160.5 43T789 159.5T832 320v554q89 27 140.5 66t51.5 84z"/>`),
		g.Group(children),
	)
}