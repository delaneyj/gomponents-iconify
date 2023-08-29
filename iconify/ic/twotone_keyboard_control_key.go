package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneKeyboardControlKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 12l1.41 1.41L12 7.83l5.59 5.58L19 12l-7-7z"/>`),
		g.Group(children),
	)
}