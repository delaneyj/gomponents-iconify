package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1009.62 550l-435 463q-25 27-62-13V24q37-40 62-13l435 463q15 16 15 38.5t-15 37.5zm-947 463q-25 27-62-13V24q37-40 62-13l435 463q15 16 15 38.5t-15 37.5z"/>`),
		g.Group(children),
	)
}