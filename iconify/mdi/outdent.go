package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M2 3h20v3H2V3zm7 5h13v3H9V8zm0 5h13v3H9v-3zm-7 5h20v3H2v-3zM6 8l-4 4l4 4h1V8H6z" fill="currentColor"/>`),
		g.Group(children),
	)
}