package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBetweenUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 18a.5.5 0 0 1-1 0v-.5A1.5 1.5 0 0 1 4.5 16h10a1.5 1.5 0 0 1 1.5 1.5v.5a.5.5 0 0 1-1 0v-.5a.5.5 0 0 0-.5-.5h-10a.5.5 0 0 0-.5.5v.5Zm5.5-3a.5.5 0 0 0 .5-.5V6.707l3.646 3.647a.5.5 0 0 0 .708-.708l-4.5-4.5a.5.5 0 0 0-.708 0l-4.5 4.5a.5.5 0 0 0 .708.708L9 6.707V14.5a.5.5 0 0 0 .5.5ZM3 2.5A1.5 1.5 0 0 0 4.5 4h10A1.5 1.5 0 0 0 16 2.5V2a.5.5 0 0 0-1 0v.5a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5V2a.5.5 0 0 0-1 0v.5Z"/>`),
		g.Group(children),
	)
}