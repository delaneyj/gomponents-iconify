package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbSpotOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L15.11 17H9v-3c-2.5-1.43-5-3-5-8v-.11L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M22 4V2H5.2l2 2H22m-2 2H9.2l7.17 7.17C18.33 11.87 20 10.07 20 6m-7 16h2v-3h-2v3m-4 0h2v-3H9v3Z"/>`),
		g.Group(children),
	)
}