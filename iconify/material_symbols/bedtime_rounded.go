package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedtimeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.575 18.475Q18.2 20.1 16.238 21.05T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-3.675 2.313-6.475t6.187-3.4q.725-.125 1.038.338t.037 1.212q-.75 2.1-.525 4.238t1.225 3.937q1 1.8 2.688 3.113t3.862 1.787q.8.175 1.012.675t-.262 1.05Z"/>`),
		g.Group(children),
	)
}