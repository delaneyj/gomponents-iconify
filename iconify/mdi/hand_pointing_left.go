package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9h10l-1.69-3.2l-.03-.22c0-.29.12-.55.32-.74l.77-.74l4.2 4.9c.27.26.43.61.43 1v6.5c0 .77-.7 1.5-1.5 1.5h-4.36c-.61 0-1.14-.35-1.34-.85l-2.2-4.94L7.47 11H3a1 1 0 0 1-1-1a1 1 0 0 1 1-1m16 9v-8h3v8h-3Z"/>`),
		g.Group(children),
	)
}