package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 172a12 12 0 1 1-12-12a12 12 0 0 1 12 12ZM92 96a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm132-48v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16ZM64 84a28 28 0 1 0 28-28a28 28 0 0 0-28 28Zm128 88a28 28 0 1 0-28 28a28 28 0 0 0 28-28Zm-2.34-105.66a8 8 0 0 0-11.32 0l-112 112a8 8 0 0 0 11.32 11.32l112-112a8 8 0 0 0 0-11.32Z"/>`),
		g.Group(children),
	)
}