package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 13.5h1.45l3.9-3.925l-1.425-1.425l-3.925 3.9v1.45Zm6.075-4.65l.7-.7q.125-.125.125-.263t-.125-.262l-.9-.9q-.125-.125-.263-.125t-.262.125l-.7.7l1.425 1.425ZM12 22q-4.025-3.425-6.012-6.362T4 10.2q0-3.75 2.413-5.975T12 2q3.175 0 5.588 2.225T20 10.2q0 2.5-1.988 5.438T12 22Z"/>`),
		g.Group(children),
	)
}