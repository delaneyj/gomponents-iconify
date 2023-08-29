package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDropdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 6V2H2v10h9v18h19V6ZM4 10V4h20v2H11v4Zm24 18H13V8h15Z"/><path fill="currentColor" d="M15 11h2v2h-2zm4 0h7v2h-7zm-4 6h2v2h-2zm4 0h7v2h-7zm-4 6h2v2h-2zm4 0h7v2h-7z"/>`),
		g.Group(children),
	)
}