package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0Zm9 11C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11Zm-4.182-6.819V8.524h2v4.243l5.364-5.364l1.414 1.414l-5.364 5.364h4.243v2H7.818Z"/>`),
		g.Group(children),
	)
}