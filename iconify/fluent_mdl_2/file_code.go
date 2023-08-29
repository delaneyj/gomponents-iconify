package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 1408V640h-512V128H256v1792h1408zM749 941l-211 211l211 211l-90 90l-301-301l301-301l90 90zm512-90l301 301l-301 301l-90-90l211-211l-211-211l90-90z"/>`),
		g.Group(children),
	)
}