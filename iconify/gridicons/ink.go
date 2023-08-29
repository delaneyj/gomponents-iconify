package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15a7 7 0 1 0 14 0a6.963 6.963 0 0 0-1.105-3.765h.007L12 2l-5.903 9.235h.007A6.971 6.971 0 0 0 5 15z"/>`),
		g.Group(children),
	)
}