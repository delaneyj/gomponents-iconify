package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.844 9l-3.6 12H0V2.5h7.362l3 2.5H20.5v4h3.344ZM18.5 9V7H9.638l-3-2.5H2v8.687L3.256 9H18.5ZM2.344 19h16.412l2.4-8H4.744l-2.4 8Z"/>`),
		g.Group(children),
	)
}