package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPentagonOne0"><path fill="#555" stroke="#fff" stroke-width="4" d="m25.23 4.958l17.63 13.743a2 2 0 0 1 .654 2.25l-7.04 19.721A2 2 0 0 1 34.59 42H13.41a2 2 0 0 1-1.884-1.328l-7.04-19.721a2 2 0 0 1 .654-2.25L22.77 4.958a2 2 0 0 1 2.46 0Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPentagonOne0)"/>`),
		g.Group(children),
	)
}