package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneByOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.496 40.331S38.607 32.791 43.5 24C38.607 15.208 27.496 7.668 27.496 7.668v32.664M7.634 19.986H4.5v-8.828s7.313-3.49 12.993-3.49l-.01 32.664H7.634Z"/>`),
		g.Group(children),
	)
}