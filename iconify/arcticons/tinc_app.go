package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TincApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.898 32.213l-4.066-16.426M15.926 31.265l16.149 4.664m-5.351-21.614L14.639 26.071"/><ellipse cx="37.183" cy="37.404" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.317" ry="5.356"/><ellipse cx="30.547" cy="10.596" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.317" ry="5.356"/><ellipse cx="10.817" cy="29.789" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.317" ry="5.356"/>`),
		g.Group(children),
	)
}