package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 19.56l-.13.43H6.3l.16-.43a4.05 4.05 0 0 0 1.3-.17a1.6 1.6 0 0 0 .76-.55a7.22 7.22 0 0 0 .8-2l2.77-9.61a7.07 7.07 0 0 0 .35-1.81a.86.86 0 0 0-.15-.52a.94.94 0 0 0-.46-.32a4.28 4.28 0 0 0-1.22-.11l.14-.43h6.16l-.13.43a2.6 2.6 0 0 0-1.12.17a1.78 1.78 0 0 0-.81.67a9.08 9.08 0 0 0-.71 1.93l-2.74 9.63a8.76 8.76 0 0 0-.4 1.69a.83.83 0 0 0 .15.5a.92.92 0 0 0 .47.32a6.35 6.35 0 0 0 1.38.18Z"/>`),
		g.Group(children),
	)
}