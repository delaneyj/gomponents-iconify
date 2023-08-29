package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-2.6 2.6q-.3.3-.663.45T12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662L5.2 8L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM14.6 17.4l-8-8L4 12l8 8l2.6-2.6Zm4.25-1.45l-1.4-1.4L20 12l-8-8l-2.55 2.55l-1.4-1.4L10.6 2.6q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-2.55 2.55Zm-5.4-5.4ZM10.6 13.4ZM8 15v-4q0-.425.288-.713T9 10h1.025l2 2H10v3H8Zm7.45-2.45L17 11l-3.5-3.5V10h-.6l2.55 2.55Z"/>`),
		g.Group(children),
	)
}