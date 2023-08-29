package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.186 18.962l7.974 2.683l4.657-4.657l1.185 1.185l8.057-8.059a4.606 4.606 0 0 0-6.515-6.513L.186 18.961Zm10.13-7.303l6.643-6.644A2.606 2.606 0 1 1 20.645 8.7l-6.643 6.645l-3.686-3.686Zm1.087 3.915l-3.78 3.78l-3.742-1.26l5.021-5.02l2.5 2.5Z"/>`),
		g.Group(children),
	)
}