package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.575L16 18.8l-2.6 2.6q-.3.3-.663.45T12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662L5.2 8L1.375 4.2q-.3-.3-.288-.713t.313-.712q.3-.3.713-.3t.712.3l18.375 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM14.6 17.4l-8-8L4 12l8 8l2.6-2.6Zm4.25-1.45l-1.4-1.4L20 12l-8-8l-2.55 2.55l-1.4-1.4L10.6 2.6q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-2.55 2.55Zm-5.4-5.4ZM10.6 13.4ZM8 14v-3q0-.425.288-.713T9 10h1.025l2 2H10v2q0 .425-.288.713T9 15q-.425 0-.713-.288T8 14Zm7.45-1.45l.85-.85q.3-.3.3-.7t-.3-.7l-1.95-1.95q-.25-.25-.55-.125t-.3.475V10h-.6l2.55 2.55Z"/>`),
		g.Group(children),
	)
}