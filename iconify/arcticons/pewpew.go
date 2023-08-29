package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pewpew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.35 25.82l11.83-3.4H5.93L4.5 34.97l3.85-9.15z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.23 26.47l-9.01 2.91l.49 3.97l4.02 1.21l4.5-8.09z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.1 26.95l-5.18 10.21l-9.57 2.02l12.4.89l2.35-13.12zm-.79-8.58l5.04-2.59l-2.96 4.45m5.35-1.29l5.5-1.78l-3.88 4.29M29.74 10.6l5.5-2.67l-3.4 4.86m6.32 0l5.34-2.43l-3.39 5.42"/>`),
		g.Group(children),
	)
}