package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSeychelles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9cb38" d="M24.2 10L2.11 50.23c.322.447.693.853 1.09 1.231L50.87 9.995H24.205"/><path fill="#003f87" d="M18.2 10H10C3.373 10 0 14.925 0 21v22c0 2.773.711 5.3 2.105 7.235L24.2 10h-6z"/><path fill="#ec1c24" d="M64 21c0-6.075-3.373-11-10-11h-3.135L3.195 51.467a8.746 8.746 0 0 0 1.868 1.355L64 30.461V21z"/><path fill="#e6e7e8" d="M64 43V30.461L5.063 52.822a9.367 9.367 0 0 0 1.938.768l56.993-10.421c0-.057.006-.112.006-.169"/><path fill="#009543" d="M54 54c6.564 0 9.932-4.835 9.994-10.831L7.001 53.59c.912.258 1.903.41 2.999.41h44z"/>`),
		g.Group(children),
	)
}