package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1280 1024l695 401l-128 222l-695-401v802H896v-802l-695 401l-128-222l695-401L73 623l128-222l695 401V0h256v802l695-401l128 222l-695 401z"/>`),
		g.Group(children),
	)
}