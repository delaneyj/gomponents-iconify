package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M200.359 382.269L61.057 251.673l21.886-23.346l116.698 109.404l229.045-229.044l22.628 22.626l-250.955 250.956z"/>`),
		g.Group(children),
	)
}