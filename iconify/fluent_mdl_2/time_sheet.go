package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeSheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 256v1664H128V256h256v128H256v1408h1536V384h-128V256h256zm-384 384H512V0h1024v640zm-128-512H640v384h768V128zm-384 1024h128v128H896v-256h128v128zm-448 64q0-93 35-174t96-143t142-96t175-35q93 0 174 35t143 96t96 142t35 175q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175zm768 0q0-66-25-124t-69-101t-102-69t-124-26q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25q66 0 124-25t101-68t69-102t26-125z"/>`),
		g.Group(children),
	)
}