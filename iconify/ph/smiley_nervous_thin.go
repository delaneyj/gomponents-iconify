package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyNervousThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M178.83 165.17a4 4 0 0 1-5.66 5.66L160 157.66l-13.17 13.17a4 4 0 0 1-5.66 0L128 157.66l-13.17 13.17a4 4 0 0 1-5.66 0L96 157.66l-13.17 13.17a4 4 0 0 1-5.66-5.66l16-16a4 4 0 0 1 5.66 0L112 162.34l13.17-13.17a4 4 0 0 1 5.66 0L144 162.34l13.17-13.17a4 4 0 0 1 5.66 0ZM228 128A100 100 0 1 1 128 28a100.11 100.11 0 0 1 100 100Zm-8 0a92 92 0 1 0-92 92a92.1 92.1 0 0 0 92-92ZM92 116a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm72 0a8 8 0 1 0-8-8a8 8 0 0 0 8 8Z"/>`),
		g.Group(children),
	)
}