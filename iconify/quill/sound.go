package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12v8a1 1 0 0 0 1 1h7l6.586 6.586c1.26 1.26 3.414.367 3.414-1.414V5.828c0-1.781-2.154-2.674-3.414-1.414L15 11H8a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}