package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 5h10v3l4-4l-4-4v3H0v6h2zm12 6H4V8l-4 4l4 4v-3h12V7h-2z"/>`),
		g.Group(children),
	)
}