package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpellCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.54 6.002H2.457L1.95 7.39c-.17.466-.734.722-1.26.57c-.522-.15-.81-.651-.64-1.118L2.097 1.23C2.437.296 3.565-.215 4.615.087c.61.176 1.088.6 1.285 1.142L7.948 6.84c.17.467-.118.967-.643 1.119c-.525.15-1.089-.105-1.259-.571L5.54 6.002Zm-.73-2l-.812-2.226l-.811 2.226H4.81Zm2.536 7.583l4.948-4.952a.999.999 0 1 1 1.413 1.415l-5.654 5.659a.999.999 0 0 1-1.414 0l-2.827-2.83a1.001 1.001 0 0 1 1.414-1.414l2.12 2.122Z"/>`),
		g.Group(children),
	)
}