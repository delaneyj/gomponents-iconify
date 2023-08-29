package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNineOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15h5V5h-6v6h4v2h-3v2Zm3-6h-2V7h2v2Zm-9 9V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2Zm6-6V4v12Z"/>`),
		g.Group(children),
	)
}