package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchImportVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 11h5L3.5 8.5l1.42-1.42L9.84 12l-4.92 4.92L3.5 15.5L6 13H1v-2M8 0h8l.83 5H17a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-.17L16 24H8l-.83-5H7c-.54 0-1-.21-1.38-.56L7.06 17H17V7H7.06L5.62 5.56C6 5.21 6.46 5 7 5h.17L8 0Z"/>`),
		g.Group(children),
	)
}