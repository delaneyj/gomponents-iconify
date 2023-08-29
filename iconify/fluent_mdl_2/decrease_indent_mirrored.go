package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecreaseIndentMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v128H128V128h1792zM128 1792v-128h1792v128H128zM896 640v128H128V640h768zm0 512v128H128v-128h768zm835 125l-90-90l163-163h-646V896h646l-163-163l90-90l317 317l-317 317z"/>`),
		g.Group(children),
	)
}