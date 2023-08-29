package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1 20.24c0-.42.344-.76.767-.76h20.466c.423 0 .767.34.767.76s-.344.76-.767.76H1.767A.764.764 0 0 1 1 20.24Z" clip-rule="evenodd"/><path d="M3.69 3.89c-.899.89-.899 2.324-.899 5.19v5.067c0 1.91 0 2.866.6 3.46c.599.593 1.564.593 3.493.593h10.233c1.93 0 2.894 0 3.493-.594c.6-.593.6-1.549.6-3.46V9.08c0-2.866 0-4.3-.9-5.19C19.412 3 17.965 3 15.07 3H8.93c-2.894 0-4.34 0-5.24.89Z" opacity=".5"/><path d="M8.93 14.4a.764.764 0 0 0-.767.76c0 .42.344.76.768.76h6.14c.423 0 .767-.34.767-.76s-.344-.76-.768-.76H8.93Z"/></g>`),
		g.Group(children),
	)
}