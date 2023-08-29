package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadlightsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M136 64v128a8 8 0 0 1-8 8H88a72 72 0 0 1-72-72.55C16.3 87.75 49.2 56 88.9 56H128a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M160 80a8 8 0 0 1 8-8h72a8 8 0 0 1 0 16h-72a8 8 0 0 1-8-8Zm80 88h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0-64h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm0 32h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm-96-72v128a16 16 0 0 1-16 16H88a80 80 0 0 1-80-80.61C8.33 83.62 44.62 48 88.9 48H128a16 16 0 0 1 16 16Zm-16 0H88.9C53.38 64 24.26 92.49 24 127.51A64 64 0 0 0 88 192h40Z"/></g>`),
		g.Group(children),
	)
}