package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h2V8H9v9Zm4 0h2V8h-2v9Zm-8 4V6H4V4h5V3h6v1h5v2h-1v15H5Z"/>`),
		g.Group(children),
	)
}