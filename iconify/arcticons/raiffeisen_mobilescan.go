package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaiffeisenMobilescan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.31 5.5H7.5a2 2 0 0 0-2 2v11.81m37 0V7.5a2 2 0 0 0-2-2H28.69M5.5 28.69V40.5a2 2 0 0 0 2 2h11.81m23.19-9.743V28.69M28.691 42.5h4.065"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 42.5v-8h2.619a2.687 2.687 0 0 1 0 5.373H35.85m2.619 0l2.619 2.625"/>`),
		g.Group(children),
	)
}