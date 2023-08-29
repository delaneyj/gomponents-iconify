package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManFourRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.325 22q-.575 0-1-.363t-.5-.937L8.3 9.275q-.125-.9.475-1.587t1.5-.688h3.45q.9 0 1.5.688t.475 1.587L14.175 20.7q-.075.575-.5.938t-1 .362h-1.35ZM12 6q-.825 0-1.412-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Z"/>`),
		g.Group(children),
	)
}