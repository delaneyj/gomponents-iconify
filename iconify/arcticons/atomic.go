package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atomic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.06 5.96L12.7 42.04M35.61 5.96l-6.36 36.08M7.41 15.71H43.5m-39 16.5h36.09"/>`),
		g.Group(children),
	)
}