package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v3.062A8.004 8.004 0 0 1 19.938 11H23v2h-3.062A8.004 8.004 0 0 1 13 19.938V23h-2v-3.062A8.004 8.004 0 0 1 4.062 13H1v-2h3.062A8.004 8.004 0 0 1 11 4.062V1h2Zm-1 5a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm-2 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}