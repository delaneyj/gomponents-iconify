package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2042 1728l-317 317l-90-90l163-163h-646v-128h646l-163-163l90-90l317 317zm-668 192l128 128H128V0h1115l549 549v731h-128V640h-512V128H256v1792h1118zm-94-1408h293l-293-293v293z"/>`),
		g.Group(children),
	)
}