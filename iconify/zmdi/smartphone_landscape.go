package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 85q0-17 12.5-29.5T43 43h384q17 0 29.5 12.5T469 85v214q0 17-12.5 29.5T427 341H43q-18 0-30.5-12.5T0 299V85zm384 0H85v214h299V85z"/>`),
		g.Group(children),
	)
}