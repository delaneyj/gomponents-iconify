package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 19q-.825 0-1.413-.588T13 17V7q0-.825.588-1.413T15 5h2q.825 0 1.413.588T19 7v10q0 .825-.588 1.413T17 19h-2Zm-8 0q-.825 0-1.413-.588T5 17V7q0-.825.588-1.413T7 5h2q.825 0 1.413.588T11 7v10q0 .825-.588 1.413T9 19H7Zm8-2h2V7h-2v10Zm-8 0h2V7H7v10ZM7 7v10V7Zm8 0v10V7Z"/>`),
		g.Group(children),
	)
}