package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-4.5 6a1.5 3 0 0 0 0 6a1.5 3 0 0 0 0-6zm9 0a1.5 3 0 0 0 0 6a1.5 3 0 0 0 0-6zM9 19s1.605 5 7 5s7-5 7-5H9z"/>`),
		g.Group(children),
	)
}