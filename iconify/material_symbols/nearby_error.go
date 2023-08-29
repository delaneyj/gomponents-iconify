package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbyError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v-8h2v8h-2Zm1 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Zm-9 0q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l8-8q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45L18 7.2v3.6l-6-6L4.8 12l7.2 7.2l6-6v3.6l-4.6 4.6q-.3.3-.663.45T12 22Zm0-5.6L7.6 12L12 7.6l4.4 4.4l-4.4 4.4Z"/>`),
		g.Group(children),
	)
}