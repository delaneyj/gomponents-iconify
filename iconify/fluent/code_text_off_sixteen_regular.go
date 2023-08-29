package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTextOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.122 4l-1-1H7a.5.5 0 0 1 0 1h-.878ZM9.5 4h5a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1Zm.5 2a.5.5 0 0 0 0 1h3.5a.5.5 0 0 0 0-1H10Zm-1 6H1.5a.5.5 0 0 0 0 1H9a.5.5 0 0 0 0-1Zm5.854 2.145l-13-12.999a.5.5 0 0 0-.708.707L2.292 3H1.5a.5.5 0 0 0 0 1h1.792l2 2H4.5a.5.5 0 0 0 0 1h1.792l2 2H3.5a.5.5 0 0 0 0 1h5.792l4.854 4.854a.502.502 0 0 0 .708 0a.5.5 0 0 0 0-.708Z"/>`),
		g.Group(children),
	)
}