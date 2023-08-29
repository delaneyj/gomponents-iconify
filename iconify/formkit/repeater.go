package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 5h-11C1.67 5 1 4.33 1 3.5v-1C1 1.67 1.67 1 2.5 1h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5Zm-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5h-11Zm11 8h-11C1.67 10 1 9.33 1 8.5v-1C1 6.67 1.67 6 2.5 6h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5Zm-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5h-11Zm11 8h-11c-.83 0-1.5-.67-1.5-1.5v-1c0-.83.67-1.5 1.5-1.5h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5Zm-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5h-11Z"/>`),
		g.Group(children),
	)
}