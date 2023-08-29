package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandFields(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-4.7 2C14.5 5.1 14 6.5 14 8h-4c0-1.6.8-3.1 2-4h3.3M14 14h-4c0-1.5-.5-2.9-1.3-4H12c1.2.9 2 2.4 2 4M4 4h5.3C8.5 5.1 8 6.5 8 8H4V4m0 6h2c1.2.9 2 2.3 2 4H4v-4m0 10v-4h5.3C8.5 17.1 8 18.5 8 20H4m6 0c0-1.6.8-3.1 2-4h3.3c-.8 1.1-1.3 2.5-1.3 4h-4m10 0h-4c0-1.6.8-3.1 2-4h2v4m0-6h-4c0-1.5-.5-2.9-1.3-4H20v4m0-6h-4c0-1.6.8-3.1 2-4h2v4Z"/>`),
		g.Group(children),
	)
}