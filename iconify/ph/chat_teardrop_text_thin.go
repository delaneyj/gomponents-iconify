package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropTextThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 112a4 4 0 0 1-4 4H96a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Zm-4 28H96a4 4 0 0 0 0 8h64a4 4 0 0 0 0-8Zm68-16a96.11 96.11 0 0 1-96 96H47.67A11.68 11.68 0 0 1 36 208.33V124a96 96 0 0 1 192 0Zm-8 0a88 88 0 0 0-176 0v84.33a3.67 3.67 0 0 0 3.67 3.67H132a88.1 88.1 0 0 0 88-88Z"/>`),
		g.Group(children),
	)
}