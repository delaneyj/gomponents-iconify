package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polymer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M411.826 77.913h-89.043L147.144 359.068L89.044 256L189.216 77.913h-89.043L0 256l100.174 178.087h89.043l175.639-281.155L422.956 256L322.784 434.087h89.043L512 256L411.826 77.913z"/>`),
		g.Group(children),
	)
}