package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 28a96.11 96.11 0 0 0-96 96v84.33A11.68 11.68 0 0 0 47.67 220H132a96 96 0 0 0 0-192Zm0 184H47.67a3.67 3.67 0 0 1-3.67-3.67V124a88 88 0 1 1 88 88Z"/>`),
		g.Group(children),
	)
}