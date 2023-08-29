package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccessOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h2v2H3Zm0-8V7h2v2H3Zm4 12v-2h2v2H7ZM7 5V3h2v2H7Zm8 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 12v-2h2v2h-2Zm0-8V7h2v2h-2ZM9 17q-.825 0-1.413-.588T7 15V9q0-.825.588-1.413T9 7h6q.825 0 1.413.588T17 9v6q0 .825-.588 1.413T15 17H9Zm0-2h6V9H9v6Zm3-3Z"/>`),
		g.Group(children),
	)
}