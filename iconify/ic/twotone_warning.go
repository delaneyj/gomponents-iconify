package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.47 19h15.06L12 5.99L4.47 19zM13 18h-2v-2h2v2zm0-4h-2v-4h2v4z" opacity=".3"/><path fill="currentColor" d="M1 21h22L12 2L1 21zm3.47-2L12 5.99L19.53 19H4.47zM11 16h2v2h-2zm0-6h2v4h-2z"/>`),
		g.Group(children),
	)
}