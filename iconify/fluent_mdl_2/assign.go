package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Assign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2011 1728l-318 317l-90-90l163-163h-614v-128h614l-163-163l90-90l318 317zm-624 192l128 128H256V256h512q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100h512v1073l-128-128V384h-128v256H512V384H384v1536h1003zM640 384v128h768V384h-256V256q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128H640z"/>`),
		g.Group(children),
	)
}