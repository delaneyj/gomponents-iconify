package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideSourceOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.075 21.9L17.5 20.35q-1.225.8-2.613 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.5.425-2.888T3.65 6.5L2.075 4.925q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM12 20q1.075 0 2.087-.275t1.963-.825L5.1 7.95q-.55.95-.825 1.962T4 12q0 3.325 2.337 5.663T12 20Zm8.35-2.5l-1.45-1.45q.55-.95.825-1.962T20 12q0-3.325-2.337-5.662T12 4q-1.075 0-2.087.275T7.95 5.1L6.5 3.65q1.225-.8 2.613-1.225T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.5-.425 2.888T20.35 17.5Zm-6.925-6.925Zm-2.85 2.85Z"/>`),
		g.Group(children),
	)
}