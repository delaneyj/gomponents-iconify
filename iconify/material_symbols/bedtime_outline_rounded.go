package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedtimeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q1.35 0 2.613-.438t2.312-1.262q-1.45-.525-3.063-1.688T11.026 13.7Q9.8 11.95 9.237 9.638T9.276 4.5q-2.325.825-3.8 2.863T4 12q0 3.325 2.337 5.663T12 20Zm0 2q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-3.675 2.313-6.475t6.187-3.4q.725-.125 1.038.338t.037 1.212q-.75 2.1-.525 4.238t1.225 3.937q1 1.8 2.688 3.113t3.862 1.787q.8.175 1.012.675t-.262 1.05Q18.2 20.1 16.238 21.05T12 22Zm-1.525-9.75Z"/>`),
		g.Group(children),
	)
}