package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.997 24L1.605 17.997v-12L12 0l10.396 5.997L16.5 9.402L12 6.8L7.496 9.4v5.2l4.502 2.6l4.5-2.6l5.895 3.397L12.003 24h-.006z"/>`),
		g.Group(children),
	)
}