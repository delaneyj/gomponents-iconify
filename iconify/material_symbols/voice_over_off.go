package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceOverOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.85 10.1L8 5.2q.225-.075.475-.138T9 5q1.65 0 2.825 1.175T13 9q0 .275-.038.55t-.112.55ZM1 21v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T9 14q.65 0 1.238.063t1.162.137L10 12.85q-.225.075-.475.113T9 13q-1.65 0-2.825-1.175T5 9q0-.275.038-.525T5.15 8L1.4 4.2l1.4-1.4l18.25 18.5l-1.35 1.45L17 20v1H1Zm18.95-5.05L18.4 14.4q1.1-1.025 1.725-2.425T20.75 9q0-1.575-.625-2.95T18.4 3.65l1.55-1.6q1.4 1.325 2.225 3.125T23 9q0 2.025-.825 3.825T19.95 15.95Zm-3.2-3.2l-1.6-1.6q.45-.425.725-.963T16.15 9q0-.65-.275-1.188t-.725-.962l1.6-1.6q.8.725 1.25 1.688T18.45 9q0 1.1-.45 2.063t-1.25 1.687Z"/>`),
		g.Group(children),
	)
}