package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CifraClub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Zm15.098-20.972L38.621 5.5M25.963 24.766L42.5 8.07"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.041 14.083c-3.784-3.658-9.546-4.583-13.835-1.513S15.949 21.09 5.5 24m27.798-6.64c2.8 2.654 7.003 11.38.403 14.892c-5.215 2.776-6.94 1.767-8.327 10.248"/><circle cx="18.584" cy="29.159" r="4.963" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}