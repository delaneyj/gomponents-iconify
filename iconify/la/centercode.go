package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Centercode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-2.8 4c-2.355 2.26-2.821 7.59-1.391 13.87c2.82.63 6.91-1.07 11.08-4.44c1.137-4.75-6.555-8.29-9.69-9.43zm2.8 5a2 2 0 1 1 .001 3.999A2 2 0 0 1 16 14z"/>`),
		g.Group(children),
	)
}