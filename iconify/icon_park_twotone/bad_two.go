package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBadTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4.18 26.834A2 2 0 0 0 6.175 29H10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7.84a2 2 0 0 0-1.993 1.834l-1.666 20ZM18 26.625c0 .836.52 1.584 1.275 1.94c1.649.778 4.458 2.341 5.725 4.454c1.633 2.724 1.941 7.645 1.991 8.772c.007.158.003.316.024.472c.271 1.953 4.04-.328 5.485-2.74c.785-1.308.885-3.027.803-4.37c-.089-1.435-.51-2.823-.923-4.201l-.88-2.937h10.857a2 2 0 0 0 1.925-2.543l-5.37-19.016A2 2 0 0 0 36.986 5H20a2 2 0 0 0-2 2v19.625Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBadTwo0)"/>`),
		g.Group(children),
	)
}