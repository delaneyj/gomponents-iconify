package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lectern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21v1H7v-1h2V11h6v10h2m.5-15c0-1.61-1.09-2.95-2.57-3.36c-.15-.37-.5-.64-.93-.64c-.55 0-1 .45-1 1s.45 1 1 1c.31 0 .58-.15.76-.37c1.01.32 1.74 1.26 1.74 2.37H4l1 4h14l1-4h-2.5Z"/>`),
		g.Group(children),
	)
}