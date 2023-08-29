package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prefect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.29 4L12 8v8l7.71-4V4L12 0L4.29 4zM12 16l-7.71-4v8L12 24v-8z"/>`),
		g.Group(children),
	)
}