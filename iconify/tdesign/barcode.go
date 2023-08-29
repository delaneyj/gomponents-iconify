package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h3v19H2V2Zm6.75 0v19h-2V2h2Zm1.75 0h3v19h-3V2Zm6.75 0v19h-2V2h2ZM19 2h3v19h-3V2Z"/>`),
		g.Group(children),
	)
}