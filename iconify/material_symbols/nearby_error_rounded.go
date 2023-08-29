package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbyErrorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18q-.425 0-.713-.288T20 17v-6q0-.425.288-.713T21 10q.425 0 .713.288T22 11v6q0 .425-.288.713T21 18Zm0 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Zm-9 0q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l8-8q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45L18 7.2v3.6l-6-6L4.8 12l7.2 7.2l6-6v3.6l-4.6 4.6q-.3.3-.663.45T12 22Zm-.7-6.3l-3-3Q8 12.4 8 12t.3-.7l3-3q.3-.3.7-.3t.7.3l3 3q.3.3.3.7t-.3.7l-3 3q-.3.3-.7.3t-.7-.3Z"/>`),
		g.Group(children),
	)
}