package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beenetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.328 29.42a3.625 3.625 0 0 1-6.775-1.796V25.27a3.625 3.625 0 0 1 7.25 0v1.178h-7.25m16.471 2.972a3.625 3.625 0 0 1-6.775-1.796V25.27a3.625 3.625 0 0 1 7.25 0v1.178h-7.25M15.481 24a3.625 3.625 0 0 1 0 7.25H9.5v-14.5h5.981a3.625 3.625 0 0 1 0 7.25Zm0 0H9.5m29 2.447h-29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}