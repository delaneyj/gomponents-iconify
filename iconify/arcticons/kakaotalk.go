package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kakaotalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 6C13.23 6 4.5 12.93 4.5 21.46c0 5.39 3.5 10.13 8.79 12.89l-1.76 7.06a.46.46 0 0 0 .68.51L21 36.7a24 24 0 0 0 3 .2c10.77 0 19.5-6.91 19.5-15.44S34.77 6 24 6Z"/><circle cx="23.91" cy="21.46" r="2.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.65" cy="21.46" r="2.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.18" cy="21.46" r="2.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}