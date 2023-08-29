package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1408H731l-475 475v-475H0V128zm896 512H640v256h128v128q27 0 50-10t40-27t28-41t10-50V640zm512 0h-256v256h128v128q27 0 50-10t40-27t28-41t10-50V640z"/>`),
		g.Group(children),
	)
}