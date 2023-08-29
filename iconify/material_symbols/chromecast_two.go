package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromecastTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-2.925 0-4.963-2.038T2 14q0-2.65 1.713-4.6T8 7.075V5.7q0-1.55 1.075-2.625T11.7 2q.85 0 1.613.375T14.6 3.4l5.15 6.5l.575-.475L23.1 12.9l-2.75 2.175L17.6 11.6l.575-.475l-5.15-6.5q-.25-.3-.588-.462T11.7 4q-.725 0-1.212.488T10 5.7v1.375q2.575.375 4.288 2.325T16 14q0 2.925-2.038 4.963T9 21Zm0-2q2.075 0 3.538-1.463T14 14q0-2.075-1.463-3.538T9 9q-2.075 0-3.538 1.463T4 14q0 2.075 1.463 3.538T9 19Z"/>`),
		g.Group(children),
	)
}