package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 640q128 0 245 48t208 139q91 91 139 208t48 245q0 133-50 249t-137 204t-203 137t-250 50v-128q106 0 199-40t162-110t110-163t41-199q0-106-40-199t-110-162t-163-110t-199-41H475l402 403l-90 90l-557-557l557-557l90 90l-402 403h677z"/>`),
		g.Group(children),
	)
}