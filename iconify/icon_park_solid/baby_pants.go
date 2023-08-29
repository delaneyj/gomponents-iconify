package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyPants(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBabyPants0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M34 4H14c-2 1-5 5-5 15c0 12.5 5 25 8 25c2 0 2-2.5 2-2.5V24c0-5 8-4.5 8 0v11c0 8 7 11 9 9s-2-4-2-7c0-5 4-9 4-22c0-5-2.333-9-4-11Z"/><path stroke="#000" stroke-linecap="round" d="M10 12h27"/><path stroke="#fff" d="M10.803 8c-.732 1.696-1.338 3.981-1.624 7M38 15a16.91 16.91 0 0 0-1.125-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBabyPants0)"/>`),
		g.Group(children),
	)
}