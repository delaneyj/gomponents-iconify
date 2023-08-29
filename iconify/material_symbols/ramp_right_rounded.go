package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RampRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.425 0-.713-.288T11 20v-5.3q-.7.95-1.588 1.8T7.4 18.125q-.4.275-.863.25T5.726 18q-.275-.275-.187-.688T6 16.65q1.575-1.05 3.288-2.987T11 9V6.825l-.9.875q-.275.275-.7.275T8.7 7.7q-.3-.3-.287-.713T8.7 6.3l2.6-2.6q.15-.15.325-.212T12 3.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.275.7t-.275.7q-.3.3-.713.3t-.687-.275l-.9-.9V20q0 .425-.288.713T12 21Z"/>`),
		g.Group(children),
	)
}