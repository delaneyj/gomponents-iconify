package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.825 12.025L8.7 15.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-4.6-4.6q-.15-.15-.213-.325T2.426 12q0-.2.063-.375T2.7 11.3l4.6-4.6q.3-.3.713-.3t.712.3q.3.3.3.713t-.3.712l-3.9 3.9Zm14.35-.05L15.3 8.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.6 4.6q-.3.3-.7.288t-.7-.313q-.3-.3-.3-.712t.3-.713l3.875-3.875Z"/>`),
		g.Group(children),
	)
}