package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tesla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.28A50.82 50.82 0 0 1 24 6.095a50.82 50.82 0 0 1 19.5 4.185"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.479 14.678A10.722 10.722 0 0 1 7.556 12.8c4.22-2.427 12.03-3.26 12.03-3.26L24 41.906l4.414-32.364s7.81.832 12.03 3.259a10.722 10.722 0 0 1-2.923 1.878"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.586 9.541L24 13.431l4.414-3.89"/>`),
		g.Group(children),
	)
}