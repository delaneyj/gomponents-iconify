package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.23 23.934l-14.145-8.167l-1.2.693v14.95l1.2.692l8.956-5.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.23 23.934l10.37 5.987l.069.02C42.09 38.926 33.813 45.5 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5c9.769 0 18.016 6.515 20.634 15.438"/>`),
		g.Group(children),
	)
}