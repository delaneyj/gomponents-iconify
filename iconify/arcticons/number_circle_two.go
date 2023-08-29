package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.141 19.991a8.141 8.141 0 0 0-8.579-8.13c-4.402.231-7.703 4.182-7.703 8.59M32.141 36.15H15.86l12.282-8.49c1.735-1.2 3.162-2.891 3.703-4.93a8.888 8.888 0 0 0 .297-2.278"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}