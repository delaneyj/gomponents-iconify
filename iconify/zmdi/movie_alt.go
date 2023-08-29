package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 0h42v384h-42v-43h-43v43H85v-43H43v43H0V0h43v43h42V0h171v43h43V0zM85 299v-43H43v43h42zm0-86v-42H43v42h42zm0-85V85H43v43h42zm214 171v-43h-43v43h43zm0-86v-42h-43v42h43zm0-85V85h-43v43h43z"/>`),
		g.Group(children),
	)
}