package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RampLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20V6.825l-.9.9Q9.825 8 9.412 8T8.7 7.7q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.213T12 3.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.287.688T15.3 7.7q-.275.275-.7.275t-.7-.275l-.9-.875V9q0 2.725 1.713 4.663T18 16.65q.375.25.463.663t-.188.687q-.35.35-.812.375t-.863-.25q-1.125-.775-2.013-1.625T13 14.7V20q0 .425-.288.713T12 21q-.425 0-.713-.288T11 20Z"/>`),
		g.Group(children),
	)
}