package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l-7.075-7.075L7.25 21H3v-4.25l5.475-5.475L1.4 4.2l1.425-1.425l18.4 18.4L19.8 22.6ZM5 19h1.4l4.9-4.9l-.7-.7l-.7-.7L5 17.6V19Zm6.3-4.9l-.7-.7l-.7-.7l1.4 1.4Zm4.275-1.425L14.15 11.25l.875-.875l-1.4-1.4l-.875.875l-1.425-1.425L13.6 6.15l4.25 4.25l-2.275 2.275Zm3.725-3.75l-4.25-4.2L17.875 1.9L22.1 6.125l-2.8 2.8Zm-5.85 1.625Z"/>`),
		g.Group(children),
	)
}