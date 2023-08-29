package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="18.5" cy="4.5" r="2.5"/><path stroke-linecap="round" d="M14.4 22v-.96a7 7 0 0 0-2.837-5.554c-.04-.03-.06-.045-.075-.058a2 2 0 0 1-.136-3.022l.07-.064l1.04-.946c1.628-1.479 1.133-4.153-.916-4.95a2.962 2.962 0 0 0-2.271.05l-.522.23M5.44 8.61L4 9.636M9 17l-.26.311A7.473 7.473 0 0 1 3 20m13-8a8.246 8.246 0 0 0 4 0"/></g>`),
		g.Group(children),
	)
}