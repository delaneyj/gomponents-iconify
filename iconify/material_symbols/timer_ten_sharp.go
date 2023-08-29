package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16h3V8h-3v8Zm-3 3V5h9v14h-9Zm-5 0V8H4V5h5v14H6Z"/>`),
		g.Group(children),
	)
}