package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18v-2h6v2h-6Zm-3.95 0l-5.7-5.695l1.425-1.43L6.05 15.15l9.172-9.175L16.65 7.4L6.05 18ZM14 14v-2h6v2h-6Zm4-4V8h6v2h-6Z"/>`),
		g.Group(children),
	)
}