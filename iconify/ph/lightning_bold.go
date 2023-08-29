package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M219.68 117.26a12 12 0 0 0-7.47-8.5l-54.44-20.41l14-70a12 12 0 0 0-20.54-10.54l-112 120a12 12 0 0 0 4.56 19.43l54.44 20.41l-14 70a12 12 0 0 0 20.54 10.54l112-120a12 12 0 0 0 2.91-10.93Zm-103.63 83.67l7.72-38.58a12 12 0 0 0-7.56-13.59L69 131.07l70.93-76l-7.72 38.58a12 12 0 0 0 7.56 13.59L187 124.93Z"/>`),
		g.Group(children),
	)
}