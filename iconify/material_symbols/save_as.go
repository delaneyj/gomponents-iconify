package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h12l4 4v5.4L12.4 21H5Zm1-11h9V6H6v4Zm6 8q1.25 0 2.125-.875T15 15q0-1.25-.875-2.125T12 12q-1.25 0-2.125.875T9 15q0 1.25.875 2.125T12 18Zm3 5v-1.775l5-4.975l1.75 1.775L16.775 23H15Zm7.4-5.65l-1.775-1.75l.85-.85q.15-.15.363-.15t.362.15l1.05 1.05q.15.15.15.35t-.15.35l-.85.85Z"/>`),
		g.Group(children),
	)
}