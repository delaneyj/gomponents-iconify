package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m638 65l114 229c2 3 3 7 4 10c7 18 12 43 12 59v109c0 34-23 62-55 70v64c0 31-25 55-55 55c-31 0-55-24-55-55v-62H165v62c0 31-24 55-55 55c-30 0-55-24-55-55v-64c-32-8-55-36-55-70V363c0-16 5-41 12-59c1-3 2-7 4-10L130 65c9-17 33-33 53-33h402c20 0 44 16 53 33zm-461 69l-61 139c-3 5-3 11-3 15h542c0-4 0-10-3-15l-61-139c-8-18-32-34-51-34H228c-19 0-43 16-51 34zm-27 338c32 0 57-26 57-59c0-32-25-57-57-57s-59 25-59 57c0 33 27 59 59 59zm468 0c32 0 59-26 59-59c0-32-27-57-59-57s-57 25-57 57c0 33 25 59 57 59z"/>`),
		g.Group(children),
	)
}