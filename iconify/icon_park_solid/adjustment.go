package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAdjustment0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M5 8a3 3 0 0 1 3-3h32a3 3 0 0 1 3 3v32a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V8Z"/><path stroke="#000" d="M36 12L12 36m0-20h8m7 17h8M16 12v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAdjustment0)"/>`),
		g.Group(children),
	)
}