package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path d="M624 0Q515 0 423 53.5T277.5 199T224 400q0 146 94 258L191 786L56 650q-10-10-23-10t-22 10q-18 17-5 39q2 3 5 6l135 137L10 969q-10 9-10 22.5t9.5 22.5t23 9t22.5-9l136-137l137 138q9 9 22 9t23-9q3-4 5.5-8.5t3-9.5t0-10t-3-9.5t-5.5-7.5L236 832l128-129q112 97 260 97q81 0 155-32t127.5-85.5T992 555t32-155.5T992 244t-85.5-127.5t-127.5-85T624 0zm-.5 736Q555 736 493 709.5T385.5 638t-72-107.5T287 400q0-92 45-169.5T454.5 108T624 63q45 0 89 12t81 34t68 53t52.5 68t33.5 80.5t12 89.5q0 68-26.5 130.5T862 638t-107.5 71.5t-131 26.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}