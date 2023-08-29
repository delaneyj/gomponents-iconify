package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1536q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20v-128q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10h-293l162 163l-90 90l-317-317l317-317l90 90l-162 163h293zm-445 384l128 128H128V0h1115l549 549v729l-128-128V640h-512V128H256v1792h1091zm-67-1408h293l-293-293v293z"/>`),
		g.Group(children),
	)
}