package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDamageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-7H3.3q-.35 0-.475-.325t.15-.55l8.35-7.525q.275-.275.675-.275t.675.275l8.35 7.525q.275.225.15.55T20.7 12H19v7q0 .425-.288.713T18 20H6q-.425 0-.713-.288T5 19Zm2-1h10v-7.8l-5-4.5l-5 4.5V18Zm5-2q.825 0 1.413-.588T14 14q0-.675-.375-1.438T12 10q-1.25 1.8-1.625 2.563T10 14q0 .825.588 1.413T12 16Zm0-4.15Z"/>`),
		g.Group(children),
	)
}