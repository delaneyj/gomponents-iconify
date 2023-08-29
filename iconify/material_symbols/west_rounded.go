package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WestRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.825 13H21q.425 0 .713-.288T22 12q0-.425-.288-.713T21 11H5.825L9.7 7.1q.275-.275.288-.688T9.7 5.7q-.275-.275-.7-.275t-.7.275l-5.6 5.6q-.15.15-.213.325T2.426 12q0 .2.063.375t.212.325l5.6 5.6q.275.275.688.275T9.7 18.3q.3-.3.3-.713t-.3-.712L5.825 13Z"/>`),
		g.Group(children),
	)
}