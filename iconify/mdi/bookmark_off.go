package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20.72L18.73 22l-1.95-1.95L12 18l-7 3V8.27l-3-3L3.28 4L20 20.72m-1-3.56V5a2 2 0 0 0-2-2H7c-.59 0-1.11.27-1.5.68L19 17.16Z"/>`),
		g.Group(children),
	)
}