package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-2h14v2H5Zm.35-4L12 5l6.65 10H5.35ZM12 13Zm-2.95 0h5.9L12 8.6L9.05 13Z"/>`),
		g.Group(children),
	)
}