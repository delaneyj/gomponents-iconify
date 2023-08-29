package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4q-.825 0-1.413-.588T2 18Zm18-7V6H4v5h16Zm-10 7h10v-5H10v5Zm-6 0h4v-5H4v5Z"/>`),
		g.Group(children),
	)
}