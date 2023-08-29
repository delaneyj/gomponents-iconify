package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.292 21.707L11.586 13H7v-2h2.586l-2-2H7a3 3 0 1 0 0 6h3v2H7a5 5 0 0 1-1.255-9.841L2.292 3.707l1.415-1.414l18 18l-1.414 1.413v.001Zm-.156-5.814l-1.428-1.427A3 3 0 0 0 17 9h-3V7h3a5 5 0 0 1 3.137 8.893Z"/>`),
		g.Group(children),
	)
}