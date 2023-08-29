package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.425 0-.713-.288T4 21q0-.2.238-.888t.512-1.837q.275-1.15.513-2.712T5.5 12q0-2-.238-3.563T4.75 5.726q-.275-1.15-.513-1.837T4 3q0-.425.288-.713T5 2h14q.425 0 .713.288T20 3q0 .2-.238.888t-.512 1.837q-.275 1.15-.513 2.713T18.5 12q0 2 .238 3.563t.512 2.712q.275 1.15.513 1.838T20 21q0 .425-.288.713T19 22H5Zm2.5-10q0 2.025-.288 4.038T6.35 20h11.275q-.575-1.95-.85-3.963T16.5 12q0-2.025.275-4.038T17.625 4H6.35q.575 1.95.863 3.963T7.5 12Zm4.5 0Z"/>`),
		g.Group(children),
	)
}