package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 64a12 12 0 0 1-12 12H40a12 12 0 0 1 0-24h68a12 12 0 0 1 12 12Zm-12 28H40a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Zm0 40H40a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Zm0 40H40a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Zm40-96h68a12 12 0 0 0 0-24h-68a12 12 0 0 0 0 24Zm68 16h-68a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Zm0 40h-68a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Zm0 40h-68a12 12 0 0 0 0 24h68a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}