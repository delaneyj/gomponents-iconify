package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InHomeModeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10.775ZM4 20V10H1l11-9l11 9h-3v1h-2V8.475l-6-4.9l-6 4.9V18h2v2H4Zm10.2 2L10 17.8l1.4-1.4l2.8 2.8l5.9-5.9l1.4 1.4l-7.3 7.3Z"/>`),
		g.Group(children),
	)
}