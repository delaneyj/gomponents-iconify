package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.127 4.5h7.57a13.101 13.101 0 0 1 0 26.202h-7.57m7.57 0L29.467 43.5m-12.77-39h22.176M9.127 17.601h29.746"/>`),
		g.Group(children),
	)
}