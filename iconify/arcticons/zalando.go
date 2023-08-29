package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zalando(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.368 5.765c9.867-2.32 30.508 11.19 32.038 17.726c1.482 6.332-22.66 21.338-31.713 18.701c-6.262-1.824-7.361-34.774-.325-36.428Z"/>`),
		g.Group(children),
	)
}