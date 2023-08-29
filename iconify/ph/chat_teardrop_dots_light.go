package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropDotsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 26a98.11 98.11 0 0 0-98 98v84.33A13.68 13.68 0 0 0 47.67 222H132a98 98 0 0 0 0-196Zm0 184H47.67a1.67 1.67 0 0 1-1.67-1.67V124a86 86 0 1 1 86 86Zm10-82a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-44 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm88 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}