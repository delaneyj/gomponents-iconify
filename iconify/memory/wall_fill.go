package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 3h6v4H5m8-4h5v4h-5m-8 8v4h6v-4m2 0h5v4h-5M8 9H3v4h5m2-4h10v4H10"/>`),
		g.Group(children),
	)
}