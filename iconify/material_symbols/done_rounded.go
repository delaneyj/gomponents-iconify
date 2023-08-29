package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.55 15.15l8.475-8.475q.3-.3.713-.3t.712.3q.3.3.3.713t-.3.712l-9.2 9.2q-.3.3-.7.3t-.7-.3L4.55 13q-.3-.3-.288-.713t.313-.712q.3-.3.713-.3t.712.3l3.55 3.575Z"/>`),
		g.Group(children),
	)
}