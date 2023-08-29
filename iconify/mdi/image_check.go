package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 21c-.32-.77-.5-1.61-.5-2.5c0-.17 0-.33.03-.5H5l3.5-4.5l2.5 3l3.5-4.5l.69.92c.97-.58 2.1-.92 3.31-.92c.89 0 1.73.18 2.5.5V5c0-.53-.21-1.04-.59-1.41C20.04 3.21 19.53 3 19 3H5c-1.1 0-2 .9-2 2v14c0 .53.21 1.04.59 1.41c.37.38.88.59 1.41.59h7.5m5.25 1L15 19l1.16-1.16l1.59 1.59l3.59-3.59l1.16 1.41L17.75 22Z"/>`),
		g.Group(children),
	)
}