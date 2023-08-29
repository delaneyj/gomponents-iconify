package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 195h128v213H0V109h128V67l64-64l64 64v128zM85 365v-42H43v42h42zm0-85v-43H43v43h42zm0-85v-43H43v43h42zm128 170v-42h-42v42h42zm0-85v-43h-42v43h42zm0-85v-43h-42v43h42zm0-86V67h-42v42h42zm128 256v-42h-42v42h42zm0-85v-43h-42v43h42z"/>`),
		g.Group(children),
	)
}