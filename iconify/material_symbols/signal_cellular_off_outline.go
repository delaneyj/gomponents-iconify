package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.825 20h10.35L12 14.825L6.825 20Zm13.95 3.6l-1.6-1.6H2l8.6-8.6l-8.2-8.175L3.8 3.8l18.4 18.4l-1.425 1.4ZM22 19.175l-2-2V6.825L14.825 12L13.4 10.6L22 2v17.175Zm-4.575-4.6ZM14.6 17.4Z"/>`),
		g.Group(children),
	)
}