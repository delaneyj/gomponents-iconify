package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 12h40v8l-1.398.84a7 7 0 0 1-7.203 0L34 20l-1.398.84a7 7 0 0 1-7.203 0L24 20l-1.398.84a7 7 0 0 1-7.204 0L14 20l-1.399.84a7 7 0 0 1-7.202 0L4 20v-8Z"/><path d="M8 22.489V44h32V22M8 11.822V4h32v8"/><path fill="#555" d="M19 32h10v12H19z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShop0)"/>`),
		g.Group(children),
	)
}