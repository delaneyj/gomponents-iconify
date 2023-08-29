package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1169.166 190.752L1011.645 33.23C968.21-10.204 894.716-7.362 847.613 39.858c-47.104 47.103-50.181 120.715-6.628 164.148l157.521 157.522c43.435 43.434 116.928 40.594 164.149-6.627c47.105-47.22 50.064-120.596 6.511-164.149zM164.978 722.374l315.044 315.044l511.976-511.857l-315.044-315.044l-511.976 511.857zM0 1197.544l415.522-83.199L83.199 782.021L0 1197.544z"/>`),
		g.Group(children),
	)
}