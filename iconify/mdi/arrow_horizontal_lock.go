package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowHorizontalLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.8 7V5.5C14.8 4.1 13.4 3 12 3S9.2 4.1 9.2 5.5V7C8.6 7 8 7.6 8 8.2v3.5c0 .7.6 1.3 1.2 1.3h5.5c.7 0 1.3-.6 1.3-1.2V8.3c0-.7-.6-1.3-1.2-1.3m-1.3 0h-3V5.5c0-.8.7-1.3 1.5-1.3s1.5.5 1.5 1.3V7M6 17v3l-4-4l4-4v3h12v-3l4 4l-4 4v-3H6Z"/>`),
		g.Group(children),
	)
}