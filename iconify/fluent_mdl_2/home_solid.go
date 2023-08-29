package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 166l941 941l-90 90l-83-82v805h-512v-640H768v640H256v-805l-83 82l-90-90l941-941z"/>`),
		g.Group(children),
	)
}