package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.575L16 18.8l-2.6 2.6q-.3.3-.663.45T12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l4.025-4.025L12.25 12.2l-.025 2.85L1.375 4.2q-.3-.3-.288-.712t.313-.713q.3-.3.713-.3t.712.3l18.375 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM9 15q.425 0 .713-.288T10 14v-2h2.025l-2-2H9q-.425 0-.713.288T8 11v3q0 .425.288.713T9 15Zm9.85.95l-3.4-3.4l.85-.85q.3-.3.3-.7t-.3-.7l-2.1-2.1q-.2-.2-.45-.087t-.25.387V10h-.6L8.05 5.15L10.6 2.6q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-2.55 2.55Z"/>`),
		g.Group(children),
	)
}