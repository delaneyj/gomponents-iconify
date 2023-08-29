package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M9.167 4.5a1.167 1.167 0 1 1-2.334 0a1.167 1.167 0 0 1 2.334 0Z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0ZM1 8a7 7 0 0 1 7-7a3.5 3.5 0 1 1 0 7a3.5 3.5 0 1 0 0 7a7 7 0 0 1-7-7Zm7 4.667a1.167 1.167 0 1 1 0-2.334a1.167 1.167 0 0 1 0 2.334Z"/></g>`),
		g.Group(children),
	)
}