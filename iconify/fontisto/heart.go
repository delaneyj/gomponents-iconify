package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 6.857l-.001-.097A6.76 6.76 0 0 1 6.863.001h-.005c3.428 0 6 2.572 6.857 3.428c.857-.856 3.428-3.428 6.857-3.428c5.143 0 6.857 3.428 6.857 6.857c0 6.857-12 16.285-13.715 17.143C11.999 23.144-.001 13.716-.001 6.858z"/>`),
		g.Group(children),
	)
}