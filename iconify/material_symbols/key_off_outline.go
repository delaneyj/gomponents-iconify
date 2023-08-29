package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.75 22.575l-7.55-7.55q-.8 1.35-2.175 2.163T7 18q-2.5 0-4.25-1.75T1 12q0-1.65.813-3.025T3.974 6.8l-2.55-2.55l1.425-1.4l18.3 18.325l-1.4 1.4ZM21 9l3 3l-4.575 4.575l-3.175-3.15l1.25-.925l1.8 1.35L21.175 12l-1-1h-6.35l-2-2H21ZM7 16q1.275 0 2.263-.688t1.437-1.787l-1.4-1.4l-1.213-1.213L6.876 9.7l-1.4-1.4q-1.1.45-1.788 1.438T3 12q0 1.65 1.175 2.825T7 16Zm0-2q-.825 0-1.413-.588T5 12q0-.825.588-1.413T7 10q.825 0 1.413.588T9 12q0 .825-.588 1.413T7 14Z"/>`),
		g.Group(children),
	)
}