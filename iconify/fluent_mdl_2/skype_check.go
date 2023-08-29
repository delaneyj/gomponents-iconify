package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkypeCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1570 437q42 0 78 15t64 43t42 63t16 79q0 40-15 77t-43 65l-794 795q-28 28-65 43t-77 16q-40 0-77-15t-65-44l-362-362q-28-28-43-65t-16-77q0-42 16-78t43-64t63-42t79-16q40 0 77 15t65 43l220 220l652-653q28-28 65-43t77-15z"/>`),
		g.Group(children),
	)
}