package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inventory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.47 27.06c0 4-3 6.27-6.79 5.11a10.31 10.31 0 0 1-6.79-9.31h0c0-4 3-6.28 6.79-5.12h0a10.32 10.32 0 0 1 6.79 9.32Zm-2.38-.84a3.45 3.45 0 0 1-2.19 3.4a4.56 4.56 0 0 1-4.37-1a5.92 5.92 0 0 1-2.18-4.41m7.34-.76a.69.69 0 0 1-1 .71h0a1.43 1.43 0 0 1-.95-1.3a.7.7 0 0 1 .95-.72a1.46 1.46 0 0 1 1 1.3Zm-4.13-1.06a.7.7 0 0 1-.95.72h0a1.45 1.45 0 0 1-.95-1.3a.7.7 0 0 1 .95-.72a1.44 1.44 0 0 1 .95 1.29ZM16.8 8.43l.2 4.92M34.25 16l-.47 23.53M4.5 11.59l12.3-3.16l26.7 2.93m-39 .23l2.11 19.78l27.17 8.2l9.13-8l.59-20.25L34.25 16Z"/>`),
		g.Group(children),
	)
}