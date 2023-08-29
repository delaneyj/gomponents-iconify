package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14q-.825 0-1.413-.588T2 12q0-.825.588-1.413T4 10q.825 0 1.413.588T6 12q0 .825-.588 1.413T4 14Zm1.65 5.7l-1.4-1.4l4.35-4.35l1.4 1.4l-4.35 4.35Zm3-9.7L4.3 5.65l1.4-1.4l4.35 4.35l-1.4 1.4ZM12 22q-.825 0-1.413-.588T10 20q0-.825.588-1.413T12 18q.825 0 1.413.588T14 20q0 .825-.588 1.413T12 22Zm0-16q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm3.35 4.05l-1.4-1.45l4.4-4.35l1.4 1.4l-4.4 4.4Zm3 9.65L14 15.35l1.4-1.4l4.35 4.35l-1.4 1.4ZM20 14q-.825 0-1.413-.587T18 12q0-.825.588-1.413T20 10q.825 0 1.413.588T22 12q0 .825-.588 1.413T20 14Z"/>`),
		g.Group(children),
	)
}