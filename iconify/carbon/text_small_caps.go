package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSmallCaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 27V15h-5v-2h12v2h-5v12h-2z"/><path fill="currentColor" d="M11 27V8H2V6h20v2h-9v19h-2z"/>`),
		g.Group(children),
	)
}