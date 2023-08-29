package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.5A1.5 1.5 0 0 1 5.5 2h5A1.5 1.5 0 0 1 12 3.5V4h1a2 2 0 0 1 2 2v4.5a1.5 1.5 0 0 1-1.5 1.5H12v.5a1.5 1.5 0 0 1-1.5 1.5h-5A1.5 1.5 0 0 1 4 12.5V12H2.5A1.5 1.5 0 0 1 1 10.5V6a2 2 0 0 1 2-2h1v-.5ZM4 11v-.5A1.5 1.5 0 0 1 5.5 9h5a1.5 1.5 0 0 1 1.5 1.5v.5h1.5a.5.5 0 0 0 .5-.5V6a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4.5a.5.5 0 0 0 .5.5H4Zm1-7h6v-.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5V4Zm0 6.5v2a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}