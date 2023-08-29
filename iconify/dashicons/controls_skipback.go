package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlsSkipback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.98 7.63l6-3.6v12l-6-3.6v3.6l-8-4.8v4.8h-2v-12h2v4.8l8-4.8v3.6z"/>`),
		g.Group(children),
	)
}