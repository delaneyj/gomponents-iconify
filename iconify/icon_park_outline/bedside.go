package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bedside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 18h40v12H4zm0 12h40v12H4zm18-6h4m-4 12h4M8 42v4m32-4v4M24 18v-8m8 0H16"/>`),
		g.Group(children),
	)
}