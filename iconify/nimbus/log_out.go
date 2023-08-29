package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.3 3.75h5.88V2.5H1.3A1.25 1.25 0 0 0 .05 3.75v8.5A1.25 1.25 0 0 0 1.3 13.5h5.88v-1.25H1.3z"/><path fill="currentColor" d="m15.4 7l-4-2.74l-.71 1l3.08 2.1H4.71v1.26h9.07l-3.08 2.11l.71 1L15.4 9a1.24 1.24 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}