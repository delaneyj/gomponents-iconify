package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 85h214v299H0V0h213v85zM85 341v-42H43v42h42zm0-85v-43H43v43h42zm0-85v-43H43v43h42zm0-86V43H43v42h42zm86 256v-42h-43v42h43zm0-85v-43h-43v43h43zm0-85v-43h-43v43h43zm0-86V43h-43v42h43zm213 256V128H213v43h43v42h-43v43h43v43h-43v42h171zm-43-170v42h-42v-42h42zm0 85v43h-42v-43h42z"/>`),
		g.Group(children),
	)
}