package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 13.106h5.326v21.728H4.5zm10.824 6.852l11.57 15.157h-4.286c-1.776 0-2.69-1.172-3.631-2.416l-6.725-8.89l3.072-3.851l3.653-4.657c.94-1.244 1.855-2.416 3.63-2.416h4.286L18.447 24.05m18.777-1.872V34.73m-6.276-6.276H43.5"/>`),
		g.Group(children),
	)
}