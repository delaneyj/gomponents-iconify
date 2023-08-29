package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#597B91" d="M471.695 411.923v47.823a6.913 6.913 0 0 1-6.913 6.913H47.217a6.913 6.913 0 0 1-6.913-6.913v-47.823a6.913 6.913 0 0 1 6.913-6.913h417.566a6.913 6.913 0 0 1 6.912 6.913z"/>`),
		g.Group(children),
	)
}