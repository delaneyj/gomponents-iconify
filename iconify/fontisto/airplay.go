package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.076 24v-.007a1.201 1.201 0 0 1-.701-2.066l.001-.001l7.076-8.262a1.2 1.2 0 0 1 1.896-.003l.002.003l7.076 8.261a1.2 1.2 0 0 1-.695 2.067h-.005v.007zm14.35-6l-2.057-2.4h7.025V2.4H2.4v13.2h7.024L7.367 18H2.4A2.402 2.402 0 0 1 0 15.6V2.401a2.402 2.402 0 0 1 2.4-2.4h23.999a2.402 2.402 0 0 1 2.4 2.4V15.6a2.402 2.402 0 0 1-2.4 2.4z"/>`),
		g.Group(children),
	)
}