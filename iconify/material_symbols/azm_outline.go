package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzmOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23v-9H1L11 4h9v9L10 23Zm6-8.825l2-2V6h-6.175l-2 2H16v6.175Zm-4 4l2-2V10H7.825l-2 2H12v6.175Z"/>`),
		g.Group(children),
	)
}