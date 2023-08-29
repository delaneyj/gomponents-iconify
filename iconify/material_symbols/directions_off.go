package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-2.6 2.6q-.3.3-.663.45T12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l4.025-4.025L12.25 12.2V15L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM8 15h2v-3h2.025l-2-2H9q-.425 0-.713.288T8 11v4Zm10.85.95l-3.4-3.4L17 11l-3.5-3.5V10h-.6L8.05 5.15L10.6 2.6q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-2.55 2.55Z"/>`),
		g.Group(children),
	)
}