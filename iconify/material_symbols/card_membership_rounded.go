package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMembershipRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16q.825 0 1.413.588T22 4v11q0 .825-.588 1.413T20 17h-4v3.375q0 .575-.475.863t-.975.037l-2.1-1.05q-.2-.125-.45-.125t-.45.125l-2.1 1.05q-.5.25-.975-.037T8 20.375V17H4q-.825 0-1.413-.588T2 15V4q0-.825.588-1.413T4 2Zm0 11h16v-3H4v3Z"/>`),
		g.Group(children),
	)
}