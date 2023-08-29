package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiquorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-2h2v-3.2q-.875-.3-1.438-1.063T3 14V6h6v8q0 .975-.563 1.738T7 16.8V20h2v2H3Zm2-11h2V8H5v3Zm1 4q.425 0 .713-.288T7 14v-1H5v1q0 .425.288.713T6 15Zm5 7V9.05l3-1.1V2h5v5.95l3 1.1V22H11Zm5-17h1V4h-1v1Zm-3 7h7v-1.55l-3-1.1V7h-1v2.35l-3 1.1V12Zm0 8h7v-2h-7v2Zm0-4h7v-2h-7v2Zm-7-1Zm7 1v-2v2Z"/>`),
		g.Group(children),
	)
}