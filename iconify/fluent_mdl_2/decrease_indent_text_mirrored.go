package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecreaseIndentTextMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1280v-128h768v128H256zM0 1536v-128h1024v128H0zM0 384h1024v128H0V384zm0 640V896h1024v128H0zm256-256V640h768v128H256z"/>`),
		g.Group(children),
	)
}