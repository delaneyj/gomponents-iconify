package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendClockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.815 12.197l-7.532 1.255a.5.5 0 0 0-.386.318L2.3 20.728c-.248.64.421 1.25 1.035.942l7.674-3.837a6.5 6.5 0 0 1 10.589-5.38a.752.752 0 0 0-.263-1.124l-18-9c-.614-.307-1.283.303-1.035.942l2.598 6.958a.5.5 0 0 0 .386.318l7.532 1.255a.2.2 0 0 1 0 .395ZM17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm2 5.5h-2V15a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}