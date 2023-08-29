package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlenderSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22v-5l1.85-1.85l-.6-4.15H3V3h7V2h4v1h4l-1.85 12.15L18 17v5H6ZM5 9h1.95L6.3 5H5v4Zm7 10q.425 0 .713-.288T13 18q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18q0 .425.288.713T12 19Zm-2.3-5h4.6l1.35-9h-7.3l1.35 9Z"/>`),
		g.Group(children),
	)
}