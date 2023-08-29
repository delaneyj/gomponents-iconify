package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextScale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 5v3h-8v18h-3V8h-8V5h19z"/><path fill="currentColor" d="M7 26V14H2v-2h12v2H9v12H7z"/>`),
		g.Group(children),
	)
}