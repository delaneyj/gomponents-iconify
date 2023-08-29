package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RiscV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m152.747 318.055l123.494 146.58L512 133.046V0H152.747c180.116 12.771 179.19 297.413 0 318.055zM512 259.355V512H331.63L512 259.355zM116.785 72.148H0V512h217.905L31.47 286.45v-36.117h85.315c126.577 0 126.577-178.185 0-178.185z"/>`),
		g.Group(children),
	)
}