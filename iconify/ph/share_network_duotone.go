package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareNetworkDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 200a32 32 0 1 1-32-32a32 32 0 0 1 32 32ZM176 88a32 32 0 1 0-32-32a32 32 0 0 0 32 32Z" opacity=".2"/><path d="M176 160a39.89 39.89 0 0 0-28.62 12.09l-46.1-29.63a39.8 39.8 0 0 0 0-28.92l46.1-29.63a40 40 0 1 0-8.66-13.45l-46.1 29.63a40 40 0 1 0 0 55.82l46.1 29.63A40 40 0 1 0 176 160Zm0-128a24 24 0 1 1-24 24a24 24 0 0 1 24-24ZM64 152a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm112 72a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/></g>`),
		g.Group(children),
	)
}