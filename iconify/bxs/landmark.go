package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17h-2V9h2V7a.998.998 0 0 0-.594-.914l-9.432-4.191l-8.421 4.21A1 1 0 0 0 2 7v2h2v8H2v4h19v-4zm-5-8v8h-3V9h3zM7 9h3v8H7V9z"/>`),
		g.Group(children),
	)
}