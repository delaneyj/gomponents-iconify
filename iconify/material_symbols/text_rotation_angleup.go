package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationAngleup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.725 21.7l-1.4-1.4l9.3-9.3h-1.6V9h5v5h-2v-1.6l-9.3 9.3Zm-1.5-5.6L3.375 5.4l1.4-1.4l10.7 4.9l-1.35 1.35l-2.65-1.3l-3.15 3.15l1.25 2.65l-1.35 1.35Zm-.65-5.55l2.35-2.3l-4.35-2.1l-.05.05l2.05 4.35Z"/>`),
		g.Group(children),
	)
}