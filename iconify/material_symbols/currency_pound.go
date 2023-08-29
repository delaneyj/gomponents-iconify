package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2l.413-.25q.412-.25.9-.738t.887-1.25Q8.6 16 8.6 15q0-.275-.038-.525T8.476 14H6v-2h1.5q-.525-.825-1.012-1.738T6 8q0-2.3 1.6-3.9t3.9-1.6q1.775 0 3.15.975T16.625 6l-1.85.775q-.375-1-1.263-1.638T11.5 4.5q-1.45 0-2.475 1.025T8 8q0 1.2.6 2t1.225 2H14v2h-3.475q.05.225.063.475T10.6 15q0 1.25-.438 2.25T9.1 19H14q1 0 1.525-.525t.725-1.35L18 18q-.275 1.375-1.413 2.188T14 21H6Z"/>`),
		g.Group(children),
	)
}