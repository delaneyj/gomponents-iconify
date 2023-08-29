package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l-7.075-7.075L7.25 21H3v-4.25l5.475-5.475L1.4 4.2l1.425-1.425l18.4 18.4L19.8 22.6Zm-4.225-9.925l-4.25-4.25L13.6 6.15l4.25 4.25l-2.275 2.275Zm3.725-3.75l-4.25-4.2L17.875 1.9L22.1 6.125l-2.8 2.8Z"/>`),
		g.Group(children),
	)
}