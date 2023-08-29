package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneWarningAmber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21h22L12 2L1 21zm3.47-2L12 5.99L19.53 19H4.47zM11 16h2v2h-2zm0-6h2v4h-2z"/>`),
		g.Group(children),
	)
}