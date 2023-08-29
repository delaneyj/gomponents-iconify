package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balcony(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12v-2h2v2H8Zm6 0v-2h2v2h-2ZM3 22v-8h1v-4q0-1.65.625-3.113t1.713-2.55q1.087-1.087 2.55-1.712T12 2q1.65 0 3.113.625t2.55 1.713q1.087 1.087 1.712 2.55T20 10v4h1v8H3Zm2-2h2v-4H5v4Zm4 0h2v-4H9v4Zm-3-6h5V4.075q-2.15.35-3.575 2.012T6 10v4Zm7 0h5v-4q0-2.25-1.425-3.913T13 4.075V14Zm0 6h2v-4h-2v4Zm4 0h2v-4h-2v4Z"/>`),
		g.Group(children),
	)
}