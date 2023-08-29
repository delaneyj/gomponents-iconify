package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.29 10.28l3.24-3.25l1.06 1.06l-3.24 3.25l-1.06-1.06m.41 4.33l5.66-5.66L15.42 10l-5.66 5.67l-1.06-1.06M14.17 3L18 6.83v10.34L14.17 21H9.83L6 17.17V6.83L9.83 3h4.34M15 1H9L4 6v12l5 5h6l5-5V6l-5-5Z"/>`),
		g.Group(children),
	)
}