package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerStat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M49.932 22.56a4.725 2.593 0 0 0-3.274.76L1.383 48.166a4.725 2.593 0 0 0 0 3.668L46.658 76.68a4.725 2.593 0 0 0 6.684 0l45.275-24.846a4.725 2.593 0 0 0 0-3.668L53.342 23.32a4.725 2.593 0 0 0-3.41-.76zM50 28.82L88.596 50L50 71.18L11.406 50L50 28.82zm9.73 8.68l-7.01 10.748l18.38.19A18.385 11.628 0 0 0 59.73 37.5zm-12.455.795a18.385 11.628 0 0 0-15.271 5.16a18.385 11.628 0 0 0 2.338 14.727l12.94-8.26l7.14-10.715a18.385 11.628 0 0 0-7.147-.912zm1.381 12.576l-13.181 8.106a18.385 11.628 0 0 0 20.097 2.67A18.385 11.628 0 0 0 67.041 50.87H48.656z" color="currentColor"/>`),
		g.Group(children),
	)
}