package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuDortmund(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.341 7.978v27.433a4.61 4.61 0 0 0 4.61 4.61h3.384"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.586h15.558a3 3 0 0 1 3 3V30.8a9.221 9.221 0 0 0 9.22 9.22h0A9.221 9.221 0 0 0 42.5 30.8V15.586m0 15.215v9.221"/>`),
		g.Group(children),
	)
}