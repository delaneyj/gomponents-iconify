package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecreaseIndentArrowMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1731 643l317 317l-317 317l-90-90l163-163h-646V896h646l-163-163l90-90z"/>`),
		g.Group(children),
	)
}