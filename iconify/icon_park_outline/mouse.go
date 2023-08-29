package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 16H12v16c0 6.627 5.373 12 12 12s12-5.373 12-12V16H24Z" clip-rule="evenodd"/><path d="M36 16c0-6.627-5.373-12-12-12v12h12ZM24 4c-6.627 0-12 5.373-12 12h12V4Z"/></g>`),
		g.Group(children),
	)
}