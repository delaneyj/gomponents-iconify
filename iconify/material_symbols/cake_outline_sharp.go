package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-8h2V8h6V6.55q-.45-.3-.725-.725T10 4.8q0-.375.15-.738t.45-.662L12 2l1.4 1.4q.3.3.45.662T14 4.8q0 .6-.275 1.025T13 6.55V8h6v6h2v8H3Zm4-8h10v-4H7v4Zm-2 6h14v-4H5v4Zm2-6h10H7Zm-2 6h14H5Zm14-6H5h14Z"/>`),
		g.Group(children),
	)
}