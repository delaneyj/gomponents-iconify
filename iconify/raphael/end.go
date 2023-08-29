package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func End(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.167 5.5v8.18L6.684 5.32v20.364l14.483-8.364v8.18H25.5v-20z"/>`),
		g.Group(children),
	)
}