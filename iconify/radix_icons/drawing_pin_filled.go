package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrawingPinFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<defs><path id="radixIconsDrawingPinFilled0" d="M9.621 1.136a.5.5 0 0 1 .707 0l1.061 1.06l1.414 1.415l1.06 1.06a.5.5 0 1 1-.706.708l-.653-.653l-3.637 4.848l1.108 1.108a.5.5 0 0 1-.707.707L7.854 9.975l-1.061-1.06l-3.27 3.27a.5.5 0 1 1-.708-.708l3.27-3.27l-1.06-1.06l-1.414-1.415a.5.5 0 1 1 .707-.707l1.108 1.108l4.848-3.637l-.653-.653a.5.5 0 0 1 0-.707Z"/></defs><g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><use href="#radixIconsDrawingPinFilled0"/><use href="#radixIconsDrawingPinFilled0"/></g>`),
		g.Group(children),
	)
}