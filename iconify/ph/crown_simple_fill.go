package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m247.37 89.15l-37 113.33a8 8 0 0 1-11.71 4.39c-.23-.14-25.8-14.87-70.68-14.87s-70.45 14.73-70.7 14.88a8 8 0 0 1-11.68-4.4L8.62 89.12a12.11 12.11 0 0 1 16.9-14.62l50.19 25l41.92-69.66a12.11 12.11 0 0 1 20.74 0l41.92 69.66l50.21-25a12.1 12.1 0 0 1 16.87 14.66Z"/>`),
		g.Group(children),
	)
}