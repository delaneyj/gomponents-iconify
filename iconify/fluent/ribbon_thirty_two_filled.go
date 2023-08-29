package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 21.813V29a1 1 0 0 0 1.514.858L16 27.166l4.485 2.692A1 1 0 0 0 22 29v-7.187a11.446 11.446 0 0 1-6 1.687c-2.199 0-4.253-.617-6-1.687ZM22 20a9.956 9.956 0 0 1-6 2a9.956 9.956 0 0 1-6.36-2.284A9.98 9.98 0 0 1 6 12C6 6.477 10.477 2 16 2s10 4.477 10 10a9.985 9.985 0 0 1-4 8Z"/>`),
		g.Group(children),
	)
}