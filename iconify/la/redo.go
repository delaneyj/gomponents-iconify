package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.219 5.281L17.78 6.72L24.063 13H11c-3.844 0-7 3.156-7 7v7h2v-7c0-2.754 2.246-5 5-5h13.063l-6.282 6.281l1.438 1.438l8-8l.687-.719l-.687-.719z"/>`),
		g.Group(children),
	)
}