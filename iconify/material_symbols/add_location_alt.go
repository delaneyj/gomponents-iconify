package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddLocationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q.825 0 1.413-.588T14 10q0-.825-.588-1.413T12 8q-.825 0-1.413.588T10 10q0 .825.588 1.413T12 12Zm0 10q-4.025-3.425-6.012-6.362T4 10.2q0-3.75 2.413-5.975T12 2q.5 0 1 .063t1 .187V6h3v3h2.925q.05.275.063.588T20 10.2q0 2.5-1.988 5.438T12 22Zm6-14V5h-3V3h3V0h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}