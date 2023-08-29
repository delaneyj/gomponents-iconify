package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liedgutverzeichnis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.745 22.835a5.078 5.078 0 0 1 2.352.265a2.1 2.1 0 0 1 1.034 1.779c-.098.858-1.205 1.238-1.771 1.89a6.195 6.195 0 0 0-1.448 1.795c-.558 1.61.672 2.855.06 5.113c-.572 2.106-3.486 4.159-5.92 4.29a8.984 8.984 0 0 1-8.08-4.615c-1.079-2.29-.164-5.416 1.298-7.483c.762-1.078 2.624-.91 3.455-1.937c.736-.91.46-2.337 1.04-3.352a5.705 5.705 0 0 1 1.875-2.064a5.363 5.363 0 0 1 2.383-.65a12.95 12.95 0 0 1 2.557.283M9.286 30.894l4.169 2.286"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.986 29.828l-2.093-1.254l7.734-14.21l-.88-1.782l2.576-4.741L24.992 9.8l-2.386 4.744l-3.538.85m14.089 23.255h-7.273a2.338 2.338 0 0 1-2.079-2.079V20.464"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.911 31.167v-8.105H26.923c-1.353 0-3.118-1.245-3.118-2.598s1.765-2.597 3.118-2.597H39.91m-14.026 2.597H39.91M26.923 38.649V23.062m13.092 11.534l3.118 3.118a1.47 1.47 0 0 1-2.079 2.078l-3.117-3.118"/><circle cx="36.332" cy="32.99" r="3.949" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}