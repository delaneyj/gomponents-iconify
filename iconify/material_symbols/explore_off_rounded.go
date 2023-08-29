package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExploreOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.5.425-2.888T3.65 6.5L2.075 4.925q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L17.5 20.35q-1.225.8-2.613 1.225T12 22Zm.05-7.1l-3-3l-2.325 4.875q-.125.25.063.438t.437.062L12.05 14.9Zm8.3 2.6l-5.45-5.45l2.375-4.825q.125-.25-.063-.437t-.437-.063L11.95 9.1L6.5 3.65q1.225-.8 2.613-1.225T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.5-.425 2.888T20.35 17.5Z"/>`),
		g.Group(children),
	)
}