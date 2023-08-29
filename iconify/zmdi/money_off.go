package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M180 83q-18 0-32 6l-32-31q15-8 32-12V0h64v47q32 8 49.5 30t19.5 51h-48q-2-45-53-45zM27 23l312 312l-27 27l-48-48q-19 18-52 24v46h-64v-46q-33-7-55-28t-23-54h46q5 45 64 45q38 0 52-20l-75-74q-84-25-84-84L0 50z"/>`),
		g.Group(children),
	)
}