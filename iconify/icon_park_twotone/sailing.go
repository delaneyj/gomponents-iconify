package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSailing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M19 30h-8l8-9m20 9C39 17.008 28.994 4 19 4v26h20Z"/><path fill="#555" stroke-linejoin="round" d="M32.651 41.577L42 36H6l2 6h23.114a3 3 0 0 0 1.537-.423Z"/><path stroke-linejoin="round" d="M19 30v6"/><path d="M29 21h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSailing0)"/>`),
		g.Group(children),
	)
}