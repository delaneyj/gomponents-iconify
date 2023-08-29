package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackToTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.275 11.7L4 7.425V11H2V4h7v2H5.4l4.3 4.275L8.275 11.7ZM4 20q-.825 0-1.413-.588T2 18v-5h2v5h8v2H4Zm16-7V6h-9V4h9q.825 0 1.413.588T22 6v7h-2Zm2 2v5h-8v-5h8Z"/>`),
		g.Group(children),
	)
}