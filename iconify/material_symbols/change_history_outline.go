package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChangeHistoryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20L12 4l10 16H2Zm3.6-2h12.8L12 7.75L5.6 18Zm6.4-5.125Z"/>`),
		g.Group(children),
	)
}