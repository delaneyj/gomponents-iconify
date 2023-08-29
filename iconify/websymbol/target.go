package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 501q0 136-67 251T751 934t-251 67t-251-67T67 752T0 501t67-251T249 68T500 1t251 67t182 182t67 251zm-152 59H706V442h142q-19-110-99-190t-190-99v142H441V153q-110 19-190 99t-99 190h142v118H152q19 110 99 190t190 99V707h118v142q110-19 190-99t99-190z"/>`),
		g.Group(children),
	)
}