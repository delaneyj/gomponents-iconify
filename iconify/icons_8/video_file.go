package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8V5zm12 1.438L22.563 9H20V6.437zm-7 6.78v9.562l1.5-.936l5-3L20.938 18l-1.438-.844l-5-3l-1.5-.937zm2 3.532L17.094 18L15 19.25v-2.5z"/>`),
		g.Group(children),
	)
}