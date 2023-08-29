package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGPlusMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 15v-2h-2v-2h2V9h2v2h2v2h-2v2h-2ZM5 17v-3H1V7h2v5h2V7h2v5h1v2H7v3H5Zm12-6v4q0 .825-.588 1.413T15 17h-4q-.825 0-1.413-.588T9 15V9q0-.825.588-1.413T11 7h4q.825 0 1.413.588T17 9h-6v6h4v-2h-2v-2h4Z"/>`),
		g.Group(children),
	)
}