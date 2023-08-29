package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func N(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21.1 33.9c12.7-4.6 26.9-.7 35.5 9.6L320 359.6V64c0-17.7 14.3-32 32-32s32 14.3 32 32v384c0 13.5-8.4 25.5-21.1 30.1s-26.9.7-35.5-9.6L64 152.4V448c0 17.7-14.3 32-32 32S0 465.7 0 448V64c0-13.5 8.4-25.5 21.1-30.1z"/>`),
		g.Group(children),
	)
}