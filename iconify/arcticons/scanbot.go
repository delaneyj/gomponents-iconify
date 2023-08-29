package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanbot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.15 4.5H10.4a2 2 0 0 0-2 2v7.75m0 19.5v7.75a2 2 0 0 0 2 2h7.75M39.6 14.25V6.5a2 2 0 0 0-2-2h-7.75m0 39h7.75a2 2 0 0 0 2-2v-7.75M33.39 9.38H14.62a1.41 1.41 0 0 0-1.35 1.46h0v26.32a1.4 1.4 0 0 0 1.35 1.46h18.77a1.41 1.41 0 0 0 1.34-1.46V10.84a1.41 1.41 0 0 0-1.34-1.46Zm-17.36 6.7h12.25m-12.25 5.28h15.94m-15.94 5.28h14.65m-14.65 5.27h9.5"/>`),
		g.Group(children),
	)
}