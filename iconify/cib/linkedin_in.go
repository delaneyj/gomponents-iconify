package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.268 28H2.463V9.306h5.805zM5.362 6.756C3.506 6.756 2 5.218 2 3.362a3.362 3.362 0 0 1 6.724 0c0 1.856-1.506 3.394-3.362 3.394zM29.994 28h-5.792v-9.1c0-2.169-.044-4.95-3.018-4.95c-3.018 0-3.481 2.356-3.481 4.794V28h-5.799V9.306h5.567v2.55h.081c.775-1.469 2.668-3.019 5.492-3.019c5.875 0 6.955 3.869 6.955 8.894V28z"/>`),
		g.Group(children),
	)
}