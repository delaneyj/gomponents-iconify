package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseIndentTextMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1024V896h1024v128H0zm0 512v-128h1024v128H0zm256-256v-128h768v128H256zm0-512V640h768v128H256zM0 384h1024v128H0V384z"/>`),
		g.Group(children),
	)
}