package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UngroupObject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 896v896h128v256h-256v-128H896v128H640v-256h128v-512H256v128H0v-256h128V256H0V0h256v128h896V0h256v256h-128v512h512V640h256v256h-128zM768 1152V896H640V640h256v128h256V256H256v896h512zm128-256v256h256V896H896zm896 0h-512v256h128v256h-256v-128H896v512h896V896z"/>`),
		g.Group(children),
	)
}