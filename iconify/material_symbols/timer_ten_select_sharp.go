package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenSelectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16h3V8h-3v8Zm-3 3V5h9v14H7Zm-4 0V8H1V5h5v14H3Zm14 0v-2h4v-1h-4v-5h6v2h-4v1h4v5h-6Z"/>`),
		g.Group(children),
	)
}