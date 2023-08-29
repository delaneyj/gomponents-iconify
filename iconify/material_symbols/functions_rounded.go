package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 12L6.325 6.3q-.15-.15-.238-.337T6 5.575V5q0-.425.288-.713T7 4h9.5q.625 0 1.063.438T18 5.5q0 .625-.438 1.063T16.5 7h-5.725l4.6 4.275q.325.3.325.725t-.325.725L10.775 17H16.5q.625 0 1.063.438T18 18.5q0 .625-.438 1.063T16.5 20H6.725q-.3 0-.512-.213T6 19.276v-.95q0-.15.05-.287t.175-.263L12.5 12Z"/>`),
		g.Group(children),
	)
}