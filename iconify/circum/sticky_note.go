package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.44 3.065H5.56a2.507 2.507 0 0 0-2.5 2.5v12.87a2.507 2.507 0 0 0 2.5 2.5h8.68A2.482 2.482 0 0 0 16 20.2l4.21-4.2a2.505 2.505 0 0 0 .73-1.77V5.565a2.5 2.5 0 0 0-2.5-2.5Zm-4.38 13.5v3.37h-8.5a1.5 1.5 0 0 1-1.5-1.5V5.565a1.5 1.5 0 0 1 1.5-1.5h12.88a1.5 1.5 0 0 1 1.5 1.5v8.5h-3.38a2.507 2.507 0 0 0-2.5 2.5Zm1 3.13v-3.13a1.5 1.5 0 0 1 1.5-1.5h3.13Z"/>`),
		g.Group(children),
	)
}