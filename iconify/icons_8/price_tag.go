package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 5l-.313.28L4.28 16.813l-.686.688l.687.72l9.5 9.5l.72.687l.69-.687l11.53-11.407L27 16V5H16zm.844 2H25v8.156L14.5 25.594L6.406 17.5L16.844 7zM22 9a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}