package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.004 3.995l-.011.01m4.011-.01l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m-3.989-.01l-.011.01m-3.987-16.01h-8v16h8v-16Z"/>`),
		g.Group(children),
	)
}