package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alltrans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 43.5c1.1 0 2-.9 2-2v-35c0-1.1-.9-2-2-2h-27c-1.1 0-2 .9-2 2v29.2c0 4.3 3.5 7.8 7.8 7.8h21.2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35 33.5L30 18l-5.3 15.5m1.8-5.2h6.8M13 18.7h13.6m-6.8-4.2c0 7.6-.2 9.1-6.1 15.2m10.2-2.9c-1.5-1.5-3.8-4.9-4.1-7.8"/>`),
		g.Group(children),
	)
}