package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deliveroo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.92 21.93l1.9-17.43l9.72 1.01l-2.45 25.7L33.5 43.5L9.25 38.7L6.46 26.74L20.65 24L17.3 9.19l9.38-1.9l3.24 14.64z"/><circle cx="30.78" cy="30.34" r="2.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.87" cy="29.19" r="2.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}