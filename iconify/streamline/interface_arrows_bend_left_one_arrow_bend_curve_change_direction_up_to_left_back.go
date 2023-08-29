package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsBendLeftOneArrowBendCurveChangeDirectionUpToLeftBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 2.75L.5 5.25L3 7.75"/><path d="M13.5 11.25v-5a1 1 0 0 0-1-1H.5"/></g>`),
		g.Group(children),
	)
}