package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 112a8 8 0 0 1-8 8H96a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm-8 24H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm72-12a100.11 100.11 0 0 1-100 100H47.67A15.69 15.69 0 0 1 32 208.33V124a100 100 0 0 1 200 0Zm-16 0a84 84 0 0 0-168 0v84h84a84.09 84.09 0 0 0 84-84Z"/>`),
		g.Group(children),
	)
}