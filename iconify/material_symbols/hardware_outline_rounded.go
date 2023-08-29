package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.425 0-.713-.288T9 20V8H4q0-2.075 1.463-3.538T9 3h4q.825 0 1.413.588T15 5v1l2.575-2.575q.175-.175.475-.3T18.625 3q.575 0 .975.4t.4.975v5.25q0 .575-.413.975t-.987.4q-.25 0-.55-.125t-.475-.3L15 8v12q0 .425-.288.713T14 21h-4Zm3-9Zm-2 7h2v-6h-2v6Zm0-8h2V5H9q-.65 0-1.225.263t-1 .737H11v5Zm2 0V5v6Zm0 8v-6v6Z"/>`),
		g.Group(children),
	)
}