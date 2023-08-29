package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.5 17l-5 5l-3.5-3.5l1.5-1.5l2 2l3.5-3.5l1.5 1.5M12 9a3 3 0 0 1 3 3a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3m0-4.5c5 0 9.27 3.11 11 7.5c-.25.65-.56 1.26-.92 1.85a5.77 5.77 0 0 0-1.9-.73l.64-1.12a9.821 9.821 0 0 0-17.64 0A9.821 9.821 0 0 0 12 17.5l1.21-.07c-.14.5-.21 1.03-.21 1.57v.46l-1 .04c-5 0-9.27-3.11-11-7.5c1.73-4.39 6-7.5 11-7.5Z"/>`),
		g.Group(children),
	)
}