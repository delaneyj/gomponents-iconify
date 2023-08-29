package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMoreDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 23.975L7.45 19.4l1.4-1.4l3.175 3.15l3.175-3.175l1.4 1.425l-4.575 4.575Zm0-5L7.45 14.4l1.4-1.4l3.175 3.15l3.175-3.175l1.4 1.425l-4.575 4.575ZM8.85 11L7.425 9.575l4.6-4.6l4.575 4.6L15.175 11l-3.15-3.175L8.85 11Zm0-5L7.425 4.575l4.6-4.6l4.575 4.6L15.175 6l-3.15-3.175L8.85 6Z"/>`),
		g.Group(children),
	)
}