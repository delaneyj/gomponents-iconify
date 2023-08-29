package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathSymbolsTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 2a.75.75 0 0 1 .75.75V5.5h2.75a.75.75 0 0 1 0 1.5H7v2.75a.75.75 0 0 1-1.5 0V7H2.75a.75.75 0 0 1 0-1.5H5.5V2.75A.75.75 0 0 1 6.25 2Zm8 3.5a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5h-7Zm-.75 12.25a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1-.75-.75ZM17.75 16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM2.22 14.28a.75.75 0 1 1 1.06-1.06l2.97 2.97l2.97-2.97a.75.75 0 1 1 1.06 1.06l-2.97 2.97l2.97 2.97a.75.75 0 1 1-1.06 1.06l-2.97-2.97l-2.97 2.97a.75.75 0 0 1-1.06-1.06l2.97-2.97l-2.97-2.97Z"/>`),
		g.Group(children),
	)
}