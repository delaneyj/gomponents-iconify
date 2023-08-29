package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Step(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 256v1408H128v-640h512V640h512V256h768zm-128 128h-512v384H768v384H256v384h1536V384z"/>`),
		g.Group(children),
	)
}