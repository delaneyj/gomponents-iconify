package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trabeepocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M43.5 23.52V11.834a4.216 4.216 0 0 0-4.216-4.216H8.716A4.216 4.216 0 0 0 4.5 11.834V23.52a8.219 8.219 0 0 0 3.425 6.676l12.65 9.084a5.871 5.871 0 0 0 6.85 0l12.65-9.084A8.219 8.219 0 0 0 43.5 23.52Z"/><path fill="none" stroke="currentColor" d="M29.997 26.098H27.82a2.178 2.178 0 0 0-1.708.827l-1.457 1.84a.834.834 0 0 1-1.308 0l-1.456-1.84a2.178 2.178 0 0 0-1.708-.827h-2.179a.609.609 0 0 1-.527-.913l5.998-10.4a.608.608 0 0 1 1.053 0l5.998 10.4a.609.609 0 0 1-.528.913Z"/>`),
		g.Group(children),
	)
}