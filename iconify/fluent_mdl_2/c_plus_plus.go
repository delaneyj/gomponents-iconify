package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CPlusPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H128V128h1792zm-119 119H247v1554h1554V247zm-521 1033v-256h128v256h256v128h-256v256h-128v-256h-256v-128h256zm-640-256H384V896h256V640h128v256h256v128H768v256H640v-256z"/>`),
		g.Group(children),
	)
}