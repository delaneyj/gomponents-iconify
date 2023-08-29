package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestedArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNestedArrows0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 17V4H4v26h13"/><path fill="#555" d="M43 43V17H17v26h26Z"/><path d="M33 30H17m16 0l-5-5l5 5Zm0 0l-5 5l5-5ZM17 17v26"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNestedArrows0)"/>`),
		g.Group(children),
	)
}