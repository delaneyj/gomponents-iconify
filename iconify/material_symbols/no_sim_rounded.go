package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSimRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17.175L7.4 4.6l2.025-2.025q.275-.275.638-.425t.762-.15H18q.825 0 1.413.588T20 4v13.175ZM6 22q-.825 0-1.413-.588T4 20V8.825q0-.4.15-.762t.425-.638L4.6 7.4L1.4 4.2q-.3-.3-.287-.7t.312-.7q.3-.275.7-.287t.7.287L21.2 21.175q.275.275.275.688t-.275.712q-.3.3-.712.3t-.713-.3L15.2 18l1.425-1.4L20 19.975V20q0 .825-.588 1.413T18 22H6Z"/>`),
		g.Group(children),
	)
}