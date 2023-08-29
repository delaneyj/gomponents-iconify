package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhysicalTherapyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 17.1ZM3 20v-4q0-2.075 1.463-3.538T8 11h10.725q.95 0 1.613.65t.662 1.6q0 .775-.475 1.388t-1.2.812L17 16.125V20q0 .525-.237.95t-.638.7q-.4.275-.875.338t-.975-.138L9.55 20H3Zm12-3H9.375q-.175 0-.263.1T9 17.325q-.025.125.038.238t.212.162L15 20v-3ZM5 18h2.1q-.05-.15-.075-.3T7 17.375q0-.975.7-1.675t1.675-.7h4.075l5.35-1.475q.125-.05.175-.125t.025-.175q-.025-.1-.088-.163T18.726 13H8q-1.25 0-2.125.875T5 16v2Zm5-8q-1.65 0-2.825-1.175T6 6q0-1.65 1.175-2.825T10 2q1.65 0 2.825 1.175T14 6q0 1.65-1.175 2.825T10 10Zm0-2q.825 0 1.413-.588T12 6q0-.825-.588-1.413T10 4q-.825 0-1.413.588T8 6q0 .825.588 1.413T10 8Zm2.025 9.1ZM10 6Z"/>`),
		g.Group(children),
	)
}