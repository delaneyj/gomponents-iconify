package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v219l-768 768v805H768v-805L0 347V128zm1920 165v-37H128v37l768 768v731h256v-731l768-768zm37 987l91 91l-230 229l230 229l-91 91l-229-230l-229 230l-91-91l230-229l-230-229l91-91l229 230l229-230z"/>`),
		g.Group(children),
	)
}