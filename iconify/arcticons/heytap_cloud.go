package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeytapCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.455 30.588a7.305 7.305 0 0 0 14.045-2.81a7.303 7.303 0 0 0-7.303-7.303a7.232 7.232 0 0 0-5.611 2.63c-1.766 2.153-5.276 6.486-7.058 8.616c-1.384 1.655-4.83 3.36-7.947 3.36C9.46 35.08 4.5 30.12 4.5 24s4.96-11.08 11.08-11.08c4.766 0 8.829 3.008 10.394 7.229"/>`),
		g.Group(children),
	)
}