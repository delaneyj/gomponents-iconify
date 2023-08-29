package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeEvergreenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m227.16 189.54l-51-65.54H208a4 4 0 0 0 3.17-6.44l-80-104a4 4 0 0 0-6.34 0l-80 104A4 4 0 0 0 48 124h31.82l-51 65.54A4 4 0 0 0 32 196h92v44a4 4 0 0 0 8 0v-44h92a4 4 0 0 0 3.16-6.46ZM40.18 188l51-65.54A4 4 0 0 0 88 116H56.12L128 22.56L199.88 116H168a4 4 0 0 0-3.16 6.46l51 65.54Z"/>`),
		g.Group(children),
	)
}