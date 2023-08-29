package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GgCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm2.545 4.486L14.69 13.34l3.965 3.967l1.254-1.254l-2.728-2.717l1.361-1.361l3.967 3.966l-3.967 3.965l-.57-.568l-1.239 1.252l1.809 1.808L25 15.945l-6.455-6.459zm-5.088.112L7 16.055l6.457 6.457l3.852-3.862l-3.965-3.966l-1.254 1.254l2.728 2.716l-1.367 1.366l-3.965-3.965l3.965-3.967l.57.57l1.245-1.244l-1.809-1.816z"/>`),
		g.Group(children),
	)
}