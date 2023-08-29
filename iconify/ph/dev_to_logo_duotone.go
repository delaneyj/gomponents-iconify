package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevToLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 72v112a8 8 0 0 1-8 8H24a8 8 0 0 1-8-8V72a8 8 0 0 1 8-8h208a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M232 56H24A16 16 0 0 0 8 72v112a16 16 0 0 0 16 16h208a16 16 0 0 0 16-16V72a16 16 0 0 0-16-16Zm0 128H24V72h208v112Zm-104-80v16h8a8 8 0 0 1 0 16h-8v16h16a8 8 0 0 1 0 16h-24a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h24a8 8 0 0 1 0 16Zm87.7-5.83l-18 64a8 8 0 0 1-15.4 0l-18-64a8 8 0 0 1 15.4-4.34l10.3 36.62l10.3-36.62a8 8 0 1 1 15.4 4.34ZM64 88h-8a8 8 0 0 0-8 8v64a8 8 0 0 0 8 8h8a32 32 0 0 0 32-32v-16a32 32 0 0 0-32-32Zm16 48a16 16 0 0 1-16 16v-48a16 16 0 0 1 16 16Z"/></g>`),
		g.Group(children),
	)
}