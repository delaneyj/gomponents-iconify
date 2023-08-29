package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckerboardRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h4v-4H8v4m4-4h4V8h-4v4M2 2v20h11.5c-.5-.6-.9-1.3-1.2-2H8v-4H4v-4h4V8H4V4h4v4h4V4h4v4h4v4.4c.7.3 1.4.7 2 1.2V2H2m18.1 12.5L18 16.6l-2.1-2.1l-1.4 1.4l2.1 2.1l-2.1 2.1l1.4 1.4l2.1-2.1l2.1 2.1l1.4-1.4l-2.1-2.1l2.1-2.1l-1.4-1.4Z"/>`),
		g.Group(children),
	)
}