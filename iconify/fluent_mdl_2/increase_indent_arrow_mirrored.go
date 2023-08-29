package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseIndentArrowMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 896v128h-646l163 163l-90 90l-317-317l317-317l90 90l-163 163h646z"/>`),
		g.Group(children),
	)
}