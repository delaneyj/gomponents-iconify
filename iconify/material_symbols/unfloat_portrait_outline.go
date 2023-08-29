package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfloatPortraitOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v9h-2V4H6v16h7v2H6Zm8.5-10.925L11.425 8H14V6H8v6h2V9.4l3.075 3.1l1.425-1.425ZM15 22v-7h5v7h-5Zm-3-10Z"/>`),
		g.Group(children),
	)
}