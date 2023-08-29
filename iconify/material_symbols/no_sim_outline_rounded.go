package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSimOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.175l-2-2V4h-7.15l-2 2L7.4 4.6L10 2h8q.825 0 1.413.588T20 4v13.175ZM6 22q-.825 0-1.413-.588T4 20V8.825q0-.4.15-.762t.425-.638L4.6 7.4L1.4 4.2q-.3-.3-.287-.7t.312-.7q.3-.275.7-.287t.7.287L21.2 21.175q.275.275.275.688t-.275.712q-.3.3-.712.3t-.713-.3L6 8.8V20h12v-2.025l2 2V20q0 .825-.588 1.413T18 22H6Zm7.525-11.325Zm-1.875 3.8Z"/>`),
		g.Group(children),
	)
}