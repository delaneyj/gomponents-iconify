package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parenleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M182 0h-81C68 76 43 158 26 247S0 427 0 520c0 87 8 171 22 252c15 80 34 150 59 209h77c-26-61-47-134-62-215c-14-82-21-167-21-255c0-95 9-188 28-277c20-89 46-166 79-234z"/>`),
		g.Group(children),
	)
}