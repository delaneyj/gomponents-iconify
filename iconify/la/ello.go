package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zM9.08 17c.488 3.387 3.401 6 6.92 6s6.432-2.613 6.92-6h-2.022A5.008 5.008 0 0 1 16 21a5.008 5.008 0 0 1-4.898-4H9.08z"/>`),
		g.Group(children),
	)
}