package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bankmillennium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.951 5.5c6.273.786 8.011 6.101 8.863 8.007L25.822 40.37l11.626-27.614l5.6 29.743"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel" d="M7.158 42.5s6.567-29.196 4.813-32.657"/>`),
		g.Group(children),
	)
}