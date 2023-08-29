package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFrames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V6q0-.825.588-1.413T4 4h4l4-4l4 4h4q.825 0 1.413.588T22 6v14q0 .825-.588 1.413T20 22H4Zm0-2h16V6H4v14Zm2-2V8h12v10H6Z"/>`),
		g.Group(children),
	)
}