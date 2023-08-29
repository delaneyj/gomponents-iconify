package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSendToMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 18H7V6h10v1h2V1H5v22h14v-6h-2z"/><path fill="currentColor" d="m22 12l-4-4v3h-5v2h5v3z"/>`),
		g.Group(children),
	)
}