package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineDashesTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.854 3.854a.5.5 0 0 0-.708-.708l-1 1a.5.5 0 0 0 .708.708l1-1Zm-3 2.292a.5.5 0 0 1 0 .708l-1 1a.5.5 0 0 1-.708-.708l1-1a.5.5 0 0 1 .708 0Zm-3 3a.5.5 0 0 1 0 .708l-1 1a.5.5 0 0 1-.708-.708l1-1a.5.5 0 0 1 .708 0Zm-3 3a.5.5 0 0 1 0 .708l-1 1a.5.5 0 0 1-.708-.708l1-1a.5.5 0 0 1 .708 0Zm-3 3.708a.5.5 0 0 0-.708-.708l-1 1a.5.5 0 0 0 .708.708l1-1Z"/>`),
		g.Group(children),
	)
}