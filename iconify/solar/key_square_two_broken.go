package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySquareTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linejoin="round" d="M17.26 11.44c2.618 0 4.74-2.113 4.74-4.72C22 4.113 19.878 2 17.26 2c-2.617 0-4.739 2.113-4.739 4.72c0 1.208.551 2.086.551 2.086l-5.731 5.708c-.257.256-.617.922 0 1.537l.661.658c.257.22.904.527 1.433 0l.771-.768c.772.768 1.654.329 1.984-.11c.552-.768-.11-1.537-.11-1.537l.22-.22c1.059 1.054 1.985.44 2.315 0c.551-.768 0-1.536 0-1.536c-.22-.44-.661-.44-.11-.988l.661-.659c.53.44 1.617.55 2.095.55Z"/><path stroke-linecap="round" d="M2 11.99c0 4.719 0 7.078 1.466 8.544C4.932 22 7.29 22 12.01 22s7.078 0 8.544-1.466c1.115-1.115 1.382-2.747 1.446-5.541M9.007 2c-2.794.064-4.426.33-5.541 1.446c-.977.977-1.303 2.35-1.412 4.554"/></g>`),
		g.Group(children),
	)
}