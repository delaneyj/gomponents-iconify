package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 12h3.5v1.3q0 .35.3.475t.55-.125l1.95-1.95q.3-.3.3-.7t-.3-.7l-1.95-1.95q-.25-.25-.55-.125t-.3.475V10H9q-.425 0-.713.288T8 11v3q0 .425.288.713T9 15q.425 0 .713-.288T10 14v-2Zm2 10q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l8-8q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-8 8q-.3.3-.663.45T12 22Z"/>`),
		g.Group(children),
	)
}