package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opentracks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.37 7.64L28.18 41.53a1.55 1.55 0 0 1-2 .86a1.58 1.58 0 0 1-.86-.86l-5-13a1.5 1.5 0 0 0-.88-.88l-12.95-5a1.51 1.51 0 0 1-1-1.44h0a1.57 1.57 0 0 1 1-1.43L40.35 5.62a1.54 1.54 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}