package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkerboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20h20V2H2m18 10h-4v4h4v4h-4v-4h-4v4H8v-4H4v-4h4V8H4V4h4v4h4V4h4v4h4v4m-4-4v4h-4V8h4m-4 4v4H8v-4h4Z"/>`),
		g.Group(children),
	)
}