package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Improvementroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.252 31.262A12.973 12.973 0 0 1 24 11.027m10.756 5.719A12.973 12.973 0 0 1 24 36.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 17.906l6.703-6.703L24 4.5h0v13.406zm-.001 12.188l-6.703 6.703l6.703 6.703V30.094z"/>`),
		g.Group(children),
	)
}