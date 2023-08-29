package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZipOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 12v-2h2v2h-2Zm0 2h-2v-2h2v2Zm0 2v-2h2v2h-2Zm-4.825-8l-2-2H4v12h10v-2h2v2h4V8h-4v2h-2V8h-2.825ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm0-2V6v12Z"/>`),
		g.Group(children),
	)
}