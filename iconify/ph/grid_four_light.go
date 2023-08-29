package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridFourLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 42H56a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm2 14v66h-68V54h66a2 2 0 0 1 2 2ZM56 54h66v68H54V56a2 2 0 0 1 2-2Zm-2 146v-66h68v68H56a2 2 0 0 1-2-2Zm146 2h-66v-68h68v66a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}