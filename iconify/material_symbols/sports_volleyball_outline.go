package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsVolleyballOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.55 9.45q-.775-2.175-2.525-3.637T13 4.05v1.6l6.55 3.8ZM8 13.15l3-1.75V4.05q-.8.075-1.55.363T8 5.1v8.05ZM4.65 15.1L6 14.3V6.75q-.975 1.1-1.488 2.45T4 12q0 .8.163 1.588T4.65 15.1ZM8 18.9l7-4l-3-1.75l-6.35 3.7q.5.625 1.075 1.15T8 18.9Zm4 1.1q1.875 0 3.55-.85t2.8-2.35L17 16.05l-6.6 3.8q.4.075.8.113T12 20Zm7.35-4.9q.325-.725.488-1.512T20 12l-7-4.05v3.45l6.35 3.7ZM12 12Zm0 10q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}