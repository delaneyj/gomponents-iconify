package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Character(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 14v-2.108c1.984-.504 3.501-2.476 3.501-4.882c0-2.797-2.049-5.007-4.5-5.007c-2.45 0-4.5 2.21-4.5 5.007c0 2.405 1.516 4.376 3.499 4.881V14H1a1 1 0 0 1 0-2.003h1.434C1.241 10.727.501 8.961.501 7.01c0-3.872 2.91-7.01 6.5-7.01s6.5 3.138 6.5 7.01c0 1.951-.74 3.716-1.933 4.987H13A1 1 0 0 1 13 14H8Z"/>`),
		g.Group(children),
	)
}