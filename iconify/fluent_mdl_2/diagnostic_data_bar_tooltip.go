package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagnosticDataBarTooltip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 0h1792v1792h-640l-256 256l-256-256H128V0zm1664 1664V128H256v1536h564l204 204l204-204h564zM768 640v896H512V640h256zm384 384v512H896v-512h256zm384-640v1152h-256V384h256z"/>`),
		g.Group(children),
	)
}