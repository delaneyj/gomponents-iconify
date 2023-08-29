package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutofpsSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.95 9.2l1-2.9h.1l1 2.9h-2.1ZM12 15q-2.5 0-4.25-1.75T6 9q0-2.5 1.75-4.25T12 3q2.5 0 4.25 1.75T18 9q0 2.5-1.75 4.25T12 15Zm-3.25-3h1.2l.65-1.8h2.8l.65 1.8h1.2L12.6 5h-1.25l-2.6 7ZM3 22v-5h2v5H3Zm4 0v-5h2v5H7Zm4 0v-5h2v5h-2Zm4 0v-5h6v5h-6Z"/>`),
		g.Group(children),
	)
}