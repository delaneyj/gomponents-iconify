package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PharmacyEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6 4l2-2a.732.732 0 0 1 0-.21a.75.75 0 1 1 .75.75h-.11L7.5 4H6zm3.48 1.83L8.65 7.5l.83 1.67V10h-8v-.83l.84-1.67l-.84-1.67V5h8v.83zM7.5 7H6V5.5H5V7H3.5v1H5v1.5h1V8h1.5V7z" fill="currentColor"/>`),
		g.Group(children),
	)
}