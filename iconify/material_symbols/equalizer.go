package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-8h4v8H4Zm6 0V4h4v16h-4Zm6 0V9h4v11h-4Z"/>`),
		g.Group(children),
	)
}