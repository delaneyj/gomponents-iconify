package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jelbi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.027 9.66a8.485 8.485 0 0 0 16.97 0M10.124 41.766A8.485 8.485 0 0 0 4.5 27.29M37.876 8.238A8.485 8.485 0 0 0 43.5 22.714m-26.382 8.262a6.91 6.91 0 0 0 12.883-3.474V13.303M18.688 6.234h12.618M4.5 22.713h12.618m2.549 16.807h10.334m6.907-15.284V34.57"/>`),
		g.Group(children),
	)
}