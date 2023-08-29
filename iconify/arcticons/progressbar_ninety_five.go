package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgressbarNinetyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 20.051h41v7.897h-41zm5.965 0v7.898m5.965-7.898v7.898m5.965-7.898v7.898"/>`),
		g.Group(children),
	)
}