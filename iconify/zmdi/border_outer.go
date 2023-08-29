package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 85v43h-42V85h42zm0 86v42h-42v-42h42zm86 0v42h-43v-42h43zM0 0h384v384H0V0zm341 341V43H43v298h298zm-128-85v43h-42v-43h42zm-85-85v42H85v-42h43z"/>`),
		g.Group(children),
	)
}