package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 20.24c0-.42.344-.76.767-.76h20.466c.423 0 .767.34.767.76s-.344.76-.767.76H1.767A.764.764 0 0 1 1 20.24ZM3.69 3.89c-.9.89-.9 2.324-.9 5.19v5.067c0 1.91 0 2.866.6 3.46c.6.593 1.564.593 3.494.593h10.232c1.93 0 2.895 0 3.494-.594c.6-.593.6-1.549.6-3.46V9.08c0-2.866 0-4.3-.9-5.19C19.411 3 17.964 3 15.07 3H8.93c-2.894 0-4.341 0-5.24.89Zm4.473 11.27c0-.42.343-.76.767-.76h6.14c.424 0 .767.34.767.76s-.343.76-.767.76H8.93a.764.764 0 0 1-.767-.76Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}