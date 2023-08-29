package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndDiamondOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16.175L19.175 12L15 7.825L10.825 12L15 16.175Zm-.7 2.125L9 13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h6l5.3-5.3q.3-.3.7-.3t.7.3l5.6 5.6q.3.3.3.7t-.3.7l-5.6 5.6q-.3.3-.7.3t-.7-.3ZM15 12Z"/>`),
		g.Group(children),
	)
}