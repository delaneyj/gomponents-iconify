package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSun0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m9.15 9.15l2.228 2.228M3 24h3.15m3 14.85l2.228-2.228M38.85 38.85l-2.228-2.228M45 24h-3.15m-3-14.85l-2.228 2.228M24 3v3.15"/><path fill="#000" d="M24 36c6.627 0 12-5.373 12-12s-5.373-12-12-12s-12 5.373-12 12s5.373 12 12 12Z"/><path stroke-linecap="round" d="M24 45v-3.15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSun0)"/>`),
		g.Group(children),
	)
}