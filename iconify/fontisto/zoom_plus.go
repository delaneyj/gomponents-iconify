package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.88 16.56a10.01 10.01 0 0 1-5.667 1.92H9.2l-.152.001A9.048 9.048 0 0 1 0 9.524v-.245a9.28 9.28 0 0 1 18.56 0a9.858 9.858 0 0 1-1.939 5.707l.019-.027l7.44 7.44l-1.76 1.6zM2.64 9.28a6.761 6.761 0 0 0 13.52-.04v-.055a6.513 6.513 0 0 0-1.998-4.703l-.002-.002A6.732 6.732 0 0 0 2.64 9.212v.072zm5.92 3.733V10.08H5.6v-1.6h2.96V5.574h1.6V8.48h2.88v1.6h-2.88v2.934z"/>`),
		g.Group(children),
	)
}