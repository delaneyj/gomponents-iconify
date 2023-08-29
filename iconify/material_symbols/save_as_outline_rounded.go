package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h11.175q.4 0 .763.15t.637.425l2.85 2.85q.275.275.425.638t.15.762V12.4l-2 2V7.825L16.175 5H5v14h9.4l-2 2H5Zm7-3q1.25 0 2.125-.875T15 15q0-1.25-.875-2.125T12 12q-1.25 0-2.125.875T9 15q0 1.25.875 2.125T12 18Zm-5-8h7q.425 0 .713-.288T15 9V7q0-.425-.288-.713T14 6H7q-.425 0-.713.288T6 7v2q0 .425.288.713T7 10Zm14.75 8.025L17.075 22.7q-.15.15-.338.225T16.35 23h-.85q-.2 0-.35-.15T15 22.5v-.85q0-.2.075-.387t.225-.338L20 16.25l1.75 1.775Zm.65-.675l-1.775-1.75l.85-.85q.15-.15.363-.15t.362.15l1.05 1.05q.15.15.15.35t-.15.35l-.85.85ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}