package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 2C10.477 2 6 6.477 6 12a9.985 9.985 0 0 0 4 8v9a1 1 0 0 0 1.514.858L16 27.166l4.485 2.692A1 1 0 0 0 22 29v-9a9.985 9.985 0 0 0 4-8c0-5.523-4.477-10-10-10ZM8 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm12 9.168v6.066l-3.485-2.092a1 1 0 0 0-1.03 0L12 27.234v-6.066A9.966 9.966 0 0 0 16 22a9.966 9.966 0 0 0 4-.832Z"/>`),
		g.Group(children),
	)
}