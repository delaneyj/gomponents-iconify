package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.425 0-.713-.288T9 20v-7h6v7q0 .425-.288.713T14 21h-4Zm8.6-10q-.25 0-.55-.125t-.475-.3L15 8v3H9V8H4q0-2.075 1.463-3.538T9 3h4q.825 0 1.413.588T15 5v1l2.575-2.575q.175-.175.475-.3T18.625 3q.575 0 .975.4t.4.975v5.25q0 .575-.413.975t-.987.4Z"/>`),
		g.Group(children),
	)
}