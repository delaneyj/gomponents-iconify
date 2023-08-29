package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCreditCardCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M19 7H7a2.5 2.5 0 0 0-2.5 2.5v7A2.5 2.5 0 0 0 7 19h12a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 19 7ZM5.5 9.5A1.5 1.5 0 0 1 7 8h12a1.5 1.5 0 0 1 1.5 1.5v7A1.5 1.5 0 0 1 19 18H7a1.5 1.5 0 0 1-1.5-1.5v-7Z"/><path d="M5.5 9.5h15a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-15A.5.5 0 0 1 5 11v-1a.5.5 0 0 1 .5-.5Zm3.5 3H8a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1Zm-1 2v-1h1v1H8Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCreditCardCircleFilled0)"/></g>`),
		g.Group(children),
	)
}