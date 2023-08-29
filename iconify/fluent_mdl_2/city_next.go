package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1280 37l640 640v1371h-128V731l-384-384v421h128v1280h-384v-384h-128v384H640V768h384V640H256v1408H128V512h1024v256h128V37zm128 1883V896H768v1024h128v-384h384v384h128zM384 896V768h128v128H384zm0 256v-128h128v128H384zm0 256v-128h128v128H384zm0 256v-128h128v128H384zm0 256v-128h128v128H384zm896-896v128h-128v-128h128zm0 256v128h-128v-128h128zm-256-256v128H896v-128h128zm0 256v128H896v-128h128z"/>`),
		g.Group(children),
	)
}