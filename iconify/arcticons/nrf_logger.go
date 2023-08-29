package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NrfLogger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.072" cy="20.471" r="14.572" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.601 14.914h18.68M5.5 21.338h19.781M7.455 27.762h17.826m6.781.99l9.531 9.184c2.46 2.37-.6 5.715-3.07 3.363l-9.665-9.202"/>`),
		g.Group(children),
	)
}