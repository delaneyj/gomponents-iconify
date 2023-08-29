package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cdisplayex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.8 34.8a26.008 26.008 0 0 0 7.2.998c10.77 0 19.5-6.339 19.5-14.158S34.77 7.484 24 7.484S4.5 13.822 4.5 21.64c0 4.289 2.626 8.13 6.773 10.727c0 0-.058 3.873-3.657 8.15c6.154-1.043 9.183-5.716 9.183-5.716Z"/>`),
		g.Group(children),
	)
}