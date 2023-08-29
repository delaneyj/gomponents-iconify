package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m364.118 67.313l69.255 69.255H160v32h273.373l-69.255 69.255l22.627 22.627l107.883-107.882L386.745 44.687l-22.627 22.626zM147.882 267.882l-22.627-22.627L17.373 353.137L125.255 461.02l22.627-22.627l-69.255-69.256H352v-32H78.627l69.255-69.255z"/>`),
		g.Group(children),
	)
}