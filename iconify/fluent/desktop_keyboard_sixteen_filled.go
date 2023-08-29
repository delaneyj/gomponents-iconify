package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopKeyboardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v5.997a2 2 0 0 0 2 2h1v1.011h-.494a.5.5 0 0 0 0 1H5V9.75A2.75 2.75 0 0 1 7.75 7H13V3a2 2 0 0 0-2-2H3Zm3 8.5A1.5 1.5 0 0 1 7.5 8h6A1.5 1.5 0 0 1 15 9.5v4a1.5 1.5 0 0 1-1.5 1.5h-6A1.5 1.5 0 0 1 6 13.5v-4Zm2 4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0-.5.5Zm.5-3.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm1.5 1.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm.5-1.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm1.5 1.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm.5-1.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}