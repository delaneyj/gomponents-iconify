package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenweatherAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.43 17.33c4.46-.24 9.1 2.33 10.64 8.21a6.88 6.88 0 0 1 .05 13.71H11.89c-9.21-1.25-10.471-14.66 0-16.46a10.14 10.14 0 0 1 8.54-5.46Zm10.7-5.051v-3.53m-6.25 6.12l-2.5-2.5M37.381 27.38l2.5 2.49m.09-8.75h3.53m-6.12-6.251l2.5-2.5M27.128 16.39a6.2 6.2 0 0 1 9.285 7.987"/>`),
		g.Group(children),
	)
}