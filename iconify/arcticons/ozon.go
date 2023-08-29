package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ozon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.172 8.298l5.104 5.201l7.177-1.26l-3.369 6.462l3.416 6.436l-7.186-1.208l-5.066 5.238l-1.072-7.207l-6.547-3.2l6.523-3.247l1.02-7.215z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.464 13.92c-7.027.432-18.082 3.579-25.004 15.771c6.17-3.977 10.34-3.712 10.34-3.712c-7.542 4.114-9.517 7.072-11.3 12.763c2.88-3.429 8.708-6.987 14.33-8.5c-4.8 3.289-3.84 9.46-3.84 9.46c3.909-5.006 7.24-10.806 15.474-14.536"/>`),
		g.Group(children),
	)
}