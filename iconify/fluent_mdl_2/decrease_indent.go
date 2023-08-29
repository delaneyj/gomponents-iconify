package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecreaseIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128h1792v128H128V128zm0 1664v-128h1792v128H128zM1152 768V640h768v128h-768zm0 512v-128h768v128h-768zM317 643l90 90l-163 163h646v128H244l163 163l-90 90L0 960l317-317z"/>`),
		g.Group(children),
	)
}