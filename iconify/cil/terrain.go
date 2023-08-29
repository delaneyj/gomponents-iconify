package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m40.841 312l103.652-112.88l71.904 71.904l76.29 76.289l22.626-22.626l-77.069-77.07l89.494-95.887L470.836 312H496v-19.864L328.262 104.27L215.603 224.976l-72.096-72.096L16 291.741V312h24.841zM16 392h480v32H16z"/>`),
		g.Group(children),
	)
}