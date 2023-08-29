package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAltFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.305 24h20.884c.28 11.59-8.888 21.214-20.479 21.494l-.405.006c-11.874 0-21.5-9.626-21.5-21.5s9.626-21.5 21.5-21.5c3.74-.015 7.418.96 10.66 2.828"/>`),
		g.Group(children),
	)
}