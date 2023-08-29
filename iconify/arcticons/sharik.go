package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.56 33.25A18 18 0 0 1 33.27 8.57c-5.4 16.04-5.99 18.08-24.71 24.68Zm32.62-14.63c4.22 15.2-9.77 26.76-22.57 22.56C27 27.67 30 25.92 41.18 18.62Z"/>`),
		g.Group(children),
	)
}