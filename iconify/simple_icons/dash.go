package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.21 9.967c-2.288 0-2.615 1.49-2.83 2.393A40.898 40.898 0 0 0 0 14.02h8.947c2.29 0 2.617-1.492 2.832-2.394c.285-1.178.379-1.66.379-1.66zM15.72 2.26H6.982L6.26 6.307l7.884.01c3.885 0 5.03 1.41 4.997 3.748c-.019 1.196-.537 3.225-.762 3.884c-.598 1.753-1.827 3.749-6.435 3.744l-7.666-.004l-.725 4.052h8.718c3.075 0 4.38-.36 5.767-.995c3.071-1.426 4.9-4.455 5.633-8.41C24.76 6.448 23.403 2.26 15.72 2.26z"/>`),
		g.Group(children),
	)
}