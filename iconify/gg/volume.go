package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M24 12a8 8 0 0 1-8 8v-2a6 6 0 0 0 0-12V4a8 8 0 0 1 8 8Z"/><path d="M20 12a4 4 0 0 1-4 4v-2a2 2 0 1 0 0-4V8a4 4 0 0 1 4 4Z"/><path fill-rule="evenodd" d="m9 16l6 4V4L9 8H5a4 4 0 1 0 0 8h4Zm-4-6h4l4-2.5v9L9 14H5a2 2 0 1 1 0-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}