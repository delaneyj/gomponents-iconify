package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9.293 10l4.853 4.854a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L3 3.707V13.5a.5.5 0 0 0 1 0V10h5.293ZM13 10h-.879l-8-8H13a.5.5 0 0 1 .407.79L11.114 6l2.293 3.21A.5.5 0 0 1 13 10Z"/>`),
		g.Group(children),
	)
}