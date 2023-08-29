package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstateAgent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 8h1V7h-1v1Zm-2 0h1V7h-1v1ZM14 22.5l-7-1.95V11h1.95l6.2 2.3q.825.3 1.337 1.05T17 16h-2q-1.05 0-1.65-.075T12.3 15.7l-2-.65l-.3.95l1.575.575q.7.275 1.275.35T14.2 17H19q1.65 0 2.325.537T22 19v1l-8 2.5ZM1 22V11h4v11H1Zm17.525-8q-.35-.825-1-1.45t-1.6-.975L9 9H7V6.5l7-5l7 5V14h-2.475ZM14.5 10h1V9h-1v1Zm-2 0h1V9h-1v1Z"/>`),
		g.Group(children),
	)
}