package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 224a12 12 0 0 1-12 12H64a12 12 0 0 1 0-24h128a12 12 0 0 1 12 12Zm-76-28a68.07 68.07 0 0 0 68-68V56a12 12 0 0 0-24 0v72a44 44 0 0 1-88 0V56a12 12 0 0 0-24 0v72a68.07 68.07 0 0 0 68 68Z"/>`),
		g.Group(children),
	)
}