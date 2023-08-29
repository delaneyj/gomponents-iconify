package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.797 8.797c8.396-8.396 22.01-8.396 30.406 0c8.396 8.396 8.396 22.01 0 30.406s-22.01 8.396-30.406 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34 24l-16-9.238v18.476L34 24z"/>`),
		g.Group(children),
	)
}