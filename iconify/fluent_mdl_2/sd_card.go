package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SDCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 352v1568q0 27-10 50t-27 40t-41 28t-50 10H384q-27 0-50-10t-40-27t-28-41t-10-50v-512l128-128V928L256 768V128q0-27 10-50t27-40t41-28t50-10h1056l352 352zm-128 53l-277-277H384v595l128 160v450q-31 33-63 64t-65 64v459h1280V405zM640 256h128v384H640V256zm256 0h128v384H896V256zm256 0h128v384h-128V256z"/>`),
		g.Group(children),
	)
}