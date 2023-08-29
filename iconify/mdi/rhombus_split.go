package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c-.5 0-1 .19-1.41.59l-3.3 3.29l4.71 4.7l4.71-4.7l-3.3-3.29C13 2.19 12.5 2 12 2M5.88 7.29l-3.29 3.3c-.79.78-.79 2.04 0 2.82l3.29 3.3l4.7-4.71l-4.7-4.71m12.24 0L13.42 12l4.7 4.71l3.29-3.3c.79-.78.79-2.04 0-2.82l-3.29-3.3M12 13.42l-4.71 4.7l3.3 3.29c.78.79 2.04.79 2.82 0l3.3-3.29l-4.71-4.7Z"/>`),
		g.Group(children),
	)
}