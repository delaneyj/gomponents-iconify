package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModuleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.025 17V7q0-.825.588-1.413T5.024 5H19q.825 0 1.413.588T21 7v10q0 .825-.588 1.413T19 19H5.025q-.825 0-1.413-.588T3.026 17Zm12.65-6H19V7h-3.325v4Zm-5.325 0h3.325V7H10.35v4Zm-5.325 0H8.35V7H5.025v4Zm0 6H8.35v-4H5.025v4Zm5.325 0h3.325v-4H10.35v4Zm5.325 0H19v-4h-3.325v4Z"/>`),
		g.Group(children),
	)
}