package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashwallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.85 41.434l1.408-9.362h15.035c4.137 0 6.845-7.443 6.845-11.218s-3.06-4.926-6.835-4.926H13.721l1.418-9.362h11.164c8.946 0 16.197 5.332 16.197 14.288c0 8.946-7.25 20.58-16.207 20.58H9.85Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 27.519h13.926l1.066-7.038H6.556L5.5 27.52Z"/>`),
		g.Group(children),
	)
}