package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropTextLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166 112a6 6 0 0 1-6 6H96a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm-6 26H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm70-14a98.11 98.11 0 0 1-98 98H47.67A13.68 13.68 0 0 1 34 208.33V124a98 98 0 0 1 196 0Zm-12 0a86 86 0 0 0-172 0v84.33a1.67 1.67 0 0 0 1.67 1.67H132a86.1 86.1 0 0 0 86-86Z"/>`),
		g.Group(children),
	)
}