package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupObject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 256H0V0h256v256zM2048 0v256h-256V0h256zM0 1792h256v256H0v-256zm1792 0h256v256h-256v-256zm-1024 0h896v128H640v-512H128V384h128v896h384V640h640V256H384V128h1024v512h512v1024h-128V768h-384v640H768v384zm0-512h512V768H768v512z"/>`),
		g.Group(children),
	)
}