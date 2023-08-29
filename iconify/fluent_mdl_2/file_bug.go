package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 1408V640h-512V128H256v1792h1408zM896 768h128v512H896V768zm0 640h128v128H896v-128z"/>`),
		g.Group(children),
	)
}