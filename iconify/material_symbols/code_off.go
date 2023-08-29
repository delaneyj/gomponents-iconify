package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L7 9.85l-2.175 2.175L9.4 16.6L8 18l-6-6l3.575-3.575l-4.2-4.2L2.8 2.8l18.4 18.4l-1.425 1.425Zm-1.35-7.05L17 14.15l2.175-2.175L14.6 7.4L16 6l6 6l-3.575 3.575Z"/>`),
		g.Group(children),
	)
}