package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.275 3.85l1.695 8.56l1.875-1.643c2.31 3.59 1.72 8.415-1.584 11.317a8.74 8.74 0 0 1-7.874 1.908l-.84 3.396c3.75.93 7.89.066 11.02-2.672c4.768-4.173 5.52-11.22 1.94-16.28l2.028-1.774l-8.26-2.813zM8.155 20.23c-2.313-3.59-1.722-8.416 1.58-11.317a8.745 8.745 0 0 1 7.876-1.91l.843-3.397A12.249 12.249 0 0 0 7.43 6.28c-4.764 4.174-5.518 11.223-1.938 16.283l-2.026 1.772l8.26 2.812l-1.693-8.56l-1.88 1.645z"/>`),
		g.Group(children),
	)
}