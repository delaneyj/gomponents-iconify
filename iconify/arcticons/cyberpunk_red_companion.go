package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CyberpunkRedCompanion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.3v20.8l5 4.6h34V15.9l-5-4.6h-34zm16.3 18.2h5.9m-5.9-11.9h5.9m-5.9 6h3.8m-3.8-6v11.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.4 29.5V17.6h1.9c3.3 0 5.9 2.7 5.9 5.9h0c0 3.3-2.7 5.9-5.9 5.9h-1.9v.1Zm-20.6 0V17.6h3.8c2.2 0 4 1.8 4 4s-1.8 4-4 4H9.8m4-.1l3.8 4"/>`),
		g.Group(children),
	)
}