package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 7l-5-5H3.993A.993.993 0 0 0 3 2.992v18.016a1 1 0 0 0 .993.992h16.014a.993.993 0 0 0 .993-.992V7ZM11 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6.5 7.5H8l5.5-7l4 7Z"/>`),
		g.Group(children),
	)
}