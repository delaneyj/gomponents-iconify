package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#FF9800"><path d="m40.997 6.065l7 7l-7 6.999l-7-7z"/><path d="M36 8h10v10H36z"/></g><circle cx="41" cy="13" r="3" fill="#FFEB3B"/><path fill="#2E7D32" d="M16.5 18L0 42h33z"/><path fill="#4CAF50" d="M33.6 24L19.2 42H48z"/>`),
		g.Group(children),
	)
}