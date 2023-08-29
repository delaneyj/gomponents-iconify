package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20v-3h2v3h12v-3h2v3q0 .825-.588 1.413T18 22H6ZM4 11V4q0-.825.588-1.413T6 2h8l6 6v3h-2V9h-5V4H6v7H4Zm-3 4v-2h22v2H1Zm11-4Zm0 6Z"/>`),
		g.Group(children),
	)
}