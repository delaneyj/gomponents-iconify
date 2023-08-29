package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stretching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23.0005 5.99951H8.00049C6.89592 5.99951 6.00049 6.89494 6.00049 7.99951V39.9999C6.00049 41.1044 6.89592 41.9998 8.00049 41.9998H40.0005C41.1051 41.9998 42.0005 41.1044 42.0005 39.9998V24.9998"/><path stroke-linecap="round" d="M24.001 15.9998V23.9998"/><path stroke-linecap="round" d="M42 5.99951V13.9995"/><path stroke-linecap="round" d="M32.001 23.9998H24.001"/><path d="M42 5.99951L24 23.9995"/><path stroke-linecap="round" d="M42.0005 5.99951H34.0005"/></g>`),
		g.Group(children),
	)
}