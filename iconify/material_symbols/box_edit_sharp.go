package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxEditSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21v-3.175l3.775-3.775l3.175 3.175L17.175 21H14Zm-9 0q-.825 0-1.413-.588T3 19V5.8L5.3 3h13.4L21 5.8V8l-5 5V8H8v8l4-2l2 1l-2 2v4H5Zm17-4.825L18.825 13l1.425-1.425l3.175 3.175L22 16.175ZM5.4 6h13.2l-.85-1H6.25L5.4 6Z"/>`),
		g.Group(children),
	)
}