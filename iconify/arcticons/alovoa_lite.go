package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlovoaLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 35.515l-9.636-4.94l-1.516-2.15l-13.09-10.486l-3.425-6.06l-8.424-.455l-1.424 1.94L4.5 16.758l1.727 3.727l-.151-2.97l1.606 1.06l1.515 2.213l.424 4.273l3.334 4.272l4.03-1.97l4.121 4.485l9.455 4.728l-1-2.97l4.182.09l3.424 2l6.333-.18Zm-9.758-1.818l-1.393-5.273m4.818 7.273l-6.964-.184"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.611 35.605l-8.475-.18l-6.485-6.432m3.334-1.629l-1.424-2.758l.94-3.849l2.757-2.818m3.561 2.852l-13.622-.003m-2.939-7.546l2.267 1.94l-.692 1.757l-.151 1.637"/>`),
		g.Group(children),
	)
}