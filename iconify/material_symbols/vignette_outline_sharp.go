package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VignetteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16q2.45 0 4.225-1.188T18 12q0-1.625-1.775-2.813T12 8Q9.55 8 7.775 9.188T6 12q0 1.625 1.775 2.813T12 16Zm0-2q-1.625 0-2.813-.6T8 12q0-.8 1.188-1.4T12 10q1.625 0 2.813.6T16 12q0 .8-1.188 1.4T12 14ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}