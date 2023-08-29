package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileSelected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1073 1664l-57 56l72 72H832v-128h241zm143 256l128 128H512q-27 0-50-10t-40-27t-28-41t-10-50V128q0-27 10-50t27-40t41-28t50-10h1024q27 0 50 10t40 27t28 41t10 50v1264l-128 128V128H512v1792h704zm813-482l-557 557l-269-269l90-91l179 179l467-467l90 91z"/>`),
		g.Group(children),
	)
}