package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transmission(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="10.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.25" ry="6.21"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.73 33.02h10.49V43.5H18.73zM24 33V16.92"/>`),
		g.Group(children),
	)
}