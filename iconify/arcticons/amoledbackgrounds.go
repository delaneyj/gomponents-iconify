package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amoledbackgrounds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.982 16.87l-6.575 13.056m11.204-16.111L20.574 36.5m11.019-25L26.87 34.833"/>`),
		g.Group(children),
	)
}