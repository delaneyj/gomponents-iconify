package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1276q1 3 1 48t1 113v146q0 79-1 149t0 123t-1 65H0v-65q0-52-1-122t0-150t-1-146v-113q0-45 2-48L383 128h1154l383 1148zm-128 132h-344l-128 256H600l-128-256H128v384h1664v-384zm-6-128L1445 256H475L134 1280h418l128 256h560l128-256h418z"/>`),
		g.Group(children),
	)
}