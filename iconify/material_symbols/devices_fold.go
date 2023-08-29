package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesFold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.825 0-1.413-.588T10 19V4.275q0-.6.325-1.088t.875-.737l3-1.3q1-.425 1.9.163T17 3h3q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21h-8Zm2.675-2H20V5h-3v11.675q0 .6-.325 1.113t-.875.737L14.675 19ZM2 5V3h2v2H2Zm0 16v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V7h2v2H2Zm4-4V3h2v2H6Zm0 16v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}