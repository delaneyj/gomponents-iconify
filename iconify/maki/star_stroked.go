package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarStroked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 2.69l1.07 2.68l.25.63h3l-2 1.75l-.5.44l.23.63l1 3.13l-2.48-1.77l-.57-.4l-.57.4l-2.52 1.77l1-3.13l.21-.63l-.5-.44l-2-1.75h3l.25-.63zM7.5 0l-2 5h-5l4 3.5l-2 6l5-3.5l5 3.5l-2-6l4-3.5h-5z"/>`),
		g.Group(children),
	)
}