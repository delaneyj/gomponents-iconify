package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TramDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 80v48H48V80a24 24 0 0 1 24-24h112a24 24 0 0 1 24 24Z" opacity=".2"/><path d="M184 48h-48V24h32a8 8 0 0 0 0-16H88a8 8 0 0 0 0 16h32v24H72a32 32 0 0 0-32 32v104a32 32 0 0 0 32 32h8l-14.4 19.2a8 8 0 1 0 12.8 9.6L100 216h56l21.6 28.8a8 8 0 1 0 12.8-9.6L176 216h8a32 32 0 0 0 32-32V80a32 32 0 0 0-32-32ZM72 64h112a16 16 0 0 1 16 16v40H56V80a16 16 0 0 1 16-16Zm112 136H72a16 16 0 0 1-16-16v-48h144v48a16 16 0 0 1-16 16Zm-88-28a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm88 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/></g>`),
		g.Group(children),
	)
}