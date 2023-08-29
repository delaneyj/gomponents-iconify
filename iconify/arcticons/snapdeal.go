package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapdeal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.612 6.664H33.29l9.163 15.552l-10.576-.028l-3.84-6.517H20.92l-5.307-9.007Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.953 7.773l9.76 17.106H42.5L30.567 41.336H13.856L5.5 27.155l7.453-19.382Z"/>`),
		g.Group(children),
	)
}