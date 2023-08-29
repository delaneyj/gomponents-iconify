package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeVerify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1280 1483l-128 128v-331H896v640H256v-805l-83 82l-90-90l941-942l941 942l-90 90l-83-82v214l-128 128V987l-640-640l-640 640v805h384v-640h512v331zm659-120l90 90l-557 558l-269-270l90-90l179 178l467-466z"/>`),
		g.Group(children),
	)
}