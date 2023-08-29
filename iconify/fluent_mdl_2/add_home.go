package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 1664v-640h512v768h-128v-640H896v640H256V987l-83 82l-90-90l941-942l941 942l-90 90l-83-82v293h-128V859l-640-640l-640 640v805h384zm1024 0h256v128h-256v256h-128v-256h-256v-128h256v-256h128v256z"/>`),
		g.Group(children),
	)
}